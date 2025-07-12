package provider

import "fmt"
import "context"
import "encoding/json"

import "k8s.io/klog/v2"
import cloudprovider "k8s.io/cloud-provider"
import v1 "k8s.io/api/core/v1"





// GetLoadBalancer returns whether the specified load balancer exists, and if so, what its status is.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (g *Cloud) GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (status *v1.LoadBalancerStatus, exists bool, err error) {
	klog.V(2).Infof("Get LoadBalancer cluster: %s service: %s", clusterName, service.Name)

        return nil, false, cloudprovider.NotImplemented
}


// GetLoadBalancerName returns the name of the load balancer.
func (g *Cloud) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
	klog.V(2).Infof("Get LoadBalancerNmae cluster: %s service: %s", clusterName, service.Name)

        return "Not Yet Set"
}


// EnsureLoadBalancer creates a new load balancer 'name', or updates the existing one. Returns the status of the balancer
func (g *Cloud) EnsureLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
	klog.V(2).Infof("Ensure LoadBalancer cluster: %s service: %s", clusterName, service.Name)
	klog.V(2).Infof("DEBUG:%v - %#v\nDEBUG SERVICE:  %#v\nDEBUG NODES:%#v.\n", "EnsureLoadBalancer", clusterName, service, nodes)

        klog.V(2).Infof("----------------------------------\n")

        data, err := json.MarshalIndent(service, "", "  ")
        if err != nil {
                fmt.Println("error:", err)
                return nil, err
        }
        fmt.Println(string(data))

        klog.V(2).Infof("----------------------------------\n")



/*        klog.V(2).Infof("----------------------------------\n")

        data, err := json.MarshalIndent(nodes, "", "  ")
        if err != nil {
                fmt.Println("error:", err)
                return nil, err
        }
        fmt.Println(string(data))

        klog.V(2).Infof("----------------------------------\n")
*/

        return nil, cloudprovider.NotImplemented
}

// UpdateLoadBalancer updates hosts under the specified load balancer.
func (g *Cloud) UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error {
	klog.V(2).Infof("Update LoadBalancer cluster: %s service: %s", clusterName, service.Name)
	klog.V(2).Infof("DEBUG:%v - %#v,  %#v,  %#v.\n", "UpdateLoadBalancer", clusterName, service, nodes)

        klog.V(2).Infof("----------------------------------\n")

        data, err := json.MarshalIndent(nodes, "", "  ")
        if err != nil {
                fmt.Println("error:", err)
                return err
        }
        fmt.Println(string(data))

        klog.V(2).Infof("----------------------------------\n")


        return cloudprovider.NotImplemented
}


