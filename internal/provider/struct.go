package provider

import "sync"
import "k8s.io/client-go/tools/cache"
import "k8s.io/client-go/tools/record"
import cloudprovider "k8s.io/cloud-provider"
import clientset "k8s.io/client-go/kubernetes"
import compute "google.golang.org/api/compute/v1"

type iblogInstance struct {
        Zone  string
        Name  string
        ID    uint64
        Disks []*compute.AttachedDisk
        Type  string
}

type ConfigFile struct {
}

// Cloud is an implementation of Interface, LoadBalancer and Instances for Google Compute Engine.
type Cloud struct {

	ClusterID	ClusterID

	client           clientset.Interface
	clientBuilder    cloudprovider.ControllerClientBuilder

	eventBroadcaster record.EventBroadcaster
	eventRecorder    record.EventRecorder
	projectID        string
	region           string
}



// ClusterID is the struct for maintaining information about this cluster's ID
type ClusterID struct {
        idLock     sync.RWMutex
        client     clientset.Interface
        cfgMapKey  string
        store      cache.Store
        providerID *string
        clusterID  *string
}
