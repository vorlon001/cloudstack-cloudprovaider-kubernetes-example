package provider

import cloudprovider "k8s.io/cloud-provider"



var _ cloudprovider.Interface = (*Cloud)(nil)
var _ cloudprovider.Instances = (*Cloud)(nil)
var _ cloudprovider.LoadBalancer = (*Cloud)(nil)
//var _ cloudprovider.Routes = (*Cloud)(nil)
//var _ cloudprovider.Zones = (*Cloud)(nil)
//var _ cloudprovider.PVLabeler = (*Cloud)(nil)
var _ cloudprovider.Clusters = (*Cloud)(nil)



// https://github.com/kubernetes/cloud-provider/blob/master/cloud.go
/*


// Interface is an abstract, pluggable interface for cloud providers.
type Interface interface {
	// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
	// to perform housekeeping or run custom controllers specific to the cloud provider.
	// Any tasks started here should be cleaned up when the stop channel closes.
	Initialize(clientBuilder ControllerClientBuilder, stop <-chan struct{})
	// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
	LoadBalancer() (LoadBalancer, bool)
	// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
	Instances() (Instances, bool)
	// InstancesV2 is an implementation for instances and should only be implemented by external cloud providers.
	// Implementing InstancesV2 is behaviorally identical to Instances but is optimized to significantly reduce
	// API calls to the cloud provider when registering and syncing nodes. Implementation of this interface will
	// disable calls to the Zones interface. Also returns true if the interface is supported, false otherwise.
	InstancesV2() (InstancesV2, bool)
	// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
	// DEPRECATED: Zones is deprecated in favor of retrieving zone/region information from InstancesV2.
	// This interface will not be called if InstancesV2 is enabled.
	Zones() (Zones, bool)
	// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
	Clusters() (Clusters, bool)
	// Routes returns a routes interface along with whether the interface is supported.
	Routes() (Routes, bool)
	// ProviderName returns the cloud provider ID.
	ProviderName() string
	// HasClusterID returns true if a ClusterID is required and set
	HasClusterID() bool
}


// InstancesV2 is an abstract, pluggable interface for cloud provider instances.
// Unlike the Instances interface, it is designed for external cloud providers and should only be used by them.
// Implementation of this interface will disable calls to the Zones interface.
type InstancesV2 interface {
	// InstanceExists returns true if the instance for the given node exists according to the cloud provider.
	// Use the node.name or node.spec.providerID field to find the node in the cloud provider.
	InstanceExists(ctx context.Context, node *v1.Node) (bool, error)
	// InstanceShutdown returns true if the instance is shutdown according to the cloud provider.
	// Use the node.name or node.spec.providerID field to find the node in the cloud provider.
	InstanceShutdown(ctx context.Context, node *v1.Node) (bool, error)
	// InstanceMetadata returns the instance's metadata. The values returned in InstanceMetadata are
	// translated into specific fields and labels in the Node object on registration.
	// Implementations should always check node.spec.providerID first when trying to discover the instance
	// for a given node. In cases where node.spec.providerID is empty, implementations can use other
	// properties of the node like its name, labels and annotations.
	InstanceMetadata(ctx context.Context, node *v1.Node) (*InstanceMetadata, error)
}

*/
