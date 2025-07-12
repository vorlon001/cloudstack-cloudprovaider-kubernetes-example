package provider


//import "errors"
import "fmt"
import "reflect"
import "crypto/rand"
import "encoding/hex"

import cloudprovider "k8s.io/cloud-provider"
import "k8s.io/apimachinery/pkg/fields"
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import "k8s.io/client-go/tools/record"
import v1core "k8s.io/client-go/kubernetes/typed/core/v1"
import "k8s.io/client-go/kubernetes/scheme"
import "k8s.io/client-go/tools/cache"
import "k8s.io/klog/v2"
import v1 "k8s.io/api/core/v1"
import "k8s.io/apimachinery/pkg/runtime"
import "k8s.io/apimachinery/pkg/watch"


func (g *Cloud) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {

	klog.V(2).Infof("DEBUG:%v - %#v\n", "Initialize", clientBuilder)

	g.clientBuilder = clientBuilder
	g.client = clientBuilder.ClientOrDie("cloud-provider")

	g.eventBroadcaster = record.NewBroadcaster()
	g.eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: g.client.CoreV1().Events("")})
	g.eventRecorder = g.eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: "g-cloudprovider"})

	go g.watchClusterID(stop)
}




// Continually watches for changes to the cluster id config map
func (g *Cloud) watchClusterID(stop <-chan struct{}) {

	klog.V(2).Infof("DEBUG:%v - %#v\n", "watchClusterID", "-")
	klog.V(2).Infof("DEBUG:%v, %v/%v\n", "watchClusterID", UIDNamespace, UIDConfigMapName)


	g.ClusterID = ClusterID{
		cfgMapKey: fmt.Sprintf("%v/%v", UIDNamespace, UIDConfigMapName),
		client:    g.client,
	}


	mapEventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			m, ok := obj.(*v1.ConfigMap)
			if !ok || m == nil {
				klog.Errorf("Expected v1.ConfigMap, item=%+v, typeIsOk=%v", obj, ok)
				return
			}
			if m.Namespace != UIDNamespace ||
				m.Name != UIDConfigMapName {
				return
			}

			klog.V(4).Infof("Observed new configmap for clusteriD: %v, %v; setting local values", m.Name, m.Data)
			g.ClusterID.update(m)
			klog.V(2).Infof("DEBUG:%v|mapEventHandler|AddFunc - %#v\n", "watchClusterID", m)
		},
		UpdateFunc: func(old, cur interface{}) {
			m, ok := cur.(*v1.ConfigMap)
			if !ok || m == nil {
				klog.Errorf("Expected v1.ConfigMap, item=%+v, typeIsOk=%v", cur, ok)
				return
			}

			if m.Namespace != UIDNamespace ||
				m.Name != UIDConfigMapName {
				return
			}

			if reflect.DeepEqual(old, cur) {
				return
			}

			klog.V(4).Infof("Observed updated configmap for clusteriD %v, %v; setting local values", m.Name, m.Data)
			klog.V(2).Infof("DEBUG:%v|mapEventHandler|UpdateFunc - %#v\n", "watchClusterID", m)

			g.ClusterID.update(m)
		},
	}

	listerWatcher := cache.NewListWatchFromClient(g.ClusterID.client.CoreV1().RESTClient(), "configmaps", UIDNamespace, fields.Everything())
	var controller cache.Controller
	g.ClusterID.store, controller = cache.NewInformer(newSingleObjectListerWatcher(listerWatcher, UIDConfigMapName), &v1.ConfigMap{}, updateFuncFrequency, mapEventHandler)

	controller.Run(stop)
}


func (ci *ClusterID) update(m *v1.ConfigMap) {
	ci.idLock.Lock()
	defer ci.idLock.Unlock()
	if clusterID, exists := m.Data[UIDCluster]; exists {
		ci.clusterID = &clusterID
	}
	if provID, exists := m.Data[UIDProvider]; exists {
		ci.providerID = &provID
	}
}

func makeUID() (string, error) {
	b := make([]byte, UIDLengthBytes)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func newSingleObjectListerWatcher(lw cache.ListerWatcher, objectName string) *singleObjListerWatcher {
	return &singleObjListerWatcher{lw: lw, objectName: objectName}
}

type singleObjListerWatcher struct {
	lw         cache.ListerWatcher
	objectName string
}

func (sow *singleObjListerWatcher) List(options metav1.ListOptions) (runtime.Object, error) {
	options.FieldSelector = "metadata.name=" + sow.objectName
	return sow.lw.List(options)
}

func (sow *singleObjListerWatcher) Watch(options metav1.ListOptions) (watch.Interface, error) {
	options.FieldSelector = "metadata.name=" + sow.objectName
	return sow.lw.Watch(options)
}
