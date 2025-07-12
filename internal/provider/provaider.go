package provider

import "io"
import "context"

import "k8s.io/klog/v2"
import cloudprovider "k8s.io/cloud-provider"
import v1 "k8s.io/api/core/v1"



// HasClusterID returns true if the cluster has a clusterID
func (g *Cloud) HasClusterID() bool {
        return true
}



// EnsureLoadBalancerDeleted deletes the specified load balancer if it
// exists, returning nil if the load balancer specified either didn't exist or
// was successfully deleted.
func (g *Cloud) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error {
	klog.V(2).Infof("Ensure LoadBalancer deleted cluster: %s service: %s", clusterName, service.Name)
	klog.V(2).Infof("DEBUG:%v - %#v|%#v\n", "GetLoadBalancer", clusterName, service.Name)
	return nil
}

// LoadBalancer returns an implementation of LoadBalancer for Google Compute Engine.
func (g *Cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "LoadBalancer", "-")
        return nil, false
}

// Instances returns an implementation of Instances for Google Compute Engine.
func (g *Cloud) Instances() (cloudprovider.Instances, bool) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "Instances", "-")
        return nil, false
}


// InstancesV2 returns an implementation of InstancesV2 for Google Compute Engine.
// Implement ONLY for external cloud provider
func (g *Cloud) InstancesV2() (cloudprovider.InstancesV2, bool) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "InstancesV2", "-")
        return g, true
}

// Zones returns an implementation of Zones for Google Compute Engine.
func (g *Cloud) Zones() (cloudprovider.Zones, bool) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "Zones", "-")
        return nil, false
}

// Clusters returns an implementation of Clusters for Google Compute Engine.
func (g *Cloud) Clusters() (cloudprovider.Clusters, bool) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "Clusters", "-")
        return g, true
}

// Routes returns an implementation of Routes for Google Compute Engine.
func (g *Cloud) Routes() (cloudprovider.Routes, bool) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "Routes", "-")
        return nil, false
}

// ProviderName returns the cloud provider ID.
func (g *Cloud) ProviderName() string {
	klog.V(2).Infof("DEBUG:%v - %v\n", "ProviderName", ProviderName)
        return ProviderName
}



func init() {
        cloudprovider.RegisterCloudProvider(
                ProviderName,
                func(config io.Reader) (cloudprovider.Interface, error) {
                        return newIBlogCloud(config)
                })
}


// newGCECloud creates a new instance of Cloud.
func newIBlogCloud(config io.Reader) (gceCloud *Cloud, err error) {
	cloud := Cloud{}
        return &cloud, nil
}


/*func readConfig(reader io.Reader) (*ConfigFile, error) {
        cfg := &ConfigFile{}
        return cfg, nil
}*/






