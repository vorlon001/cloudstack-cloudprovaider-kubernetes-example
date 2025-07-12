package provider

import "context"
import "time"

import "k8s.io/klog/v2"
import "k8s.io/apimachinery/pkg/types"
import cloudprovider "k8s.io/cloud-provider"
import v1 "k8s.io/api/core/v1"



// Gets the named instance, returning cloudprovider.InstanceNotFound if the instance is not found
func (g *Cloud) getInstanceByName(name string) (*iblogInstance, error) {
        klog.V(2).Infof("DEBUG:%v - %v\n", "getInstanceByName", name)
        iblogCloud := iblogInstance{}
        return &iblogCloud, nil
//        return nil, cloudprovider.InstanceNotFound
}



// NodeAddresses is an implementation of Instances.NodeAddresses.
func (g *Cloud) NodeAddresses(ctx context.Context, nodeName types.NodeName) ([]v1.NodeAddress, error) {
        timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Hour)
	_ = timeoutCtx
        defer cancel()

	klog.V(2).Infof("DEBUG:%v - %v\n", "NodeAddresses", nodeName)

	nodeAddresses := []v1.NodeAddress{}
        //instanceName := string(nodeName)
        return nodeAddresses, nil
}


// NodeAddressesByProviderID will not be called from the node that is requesting this ID.
// i.e. metadata service and other local methods cannot be used here
func (g *Cloud) NodeAddressesByProviderID(ctx context.Context, providerID string) ([]v1.NodeAddress, error) {
        timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Hour)
        _ = timeoutCtx
        defer cancel()


	klog.V(2).Infof("DEBUG:%v - %v\n", "NodeAddressesByProviderID", providerID)

        nodeAddresses := []v1.NodeAddress{}
        //instanceName := string(nodeName)
        return nodeAddresses, nil

}


// InstanceID returns the cloud provider ID of the node with the specified NodeName.
func (g *Cloud) InstanceID(ctx context.Context, nodeName types.NodeName) (string, error) {
        instanceName := mapNodeNameToInstanceName(nodeName)

        klog.V(2).Infof("DEBUG:%v - %v\n", "InstanceID", nodeName)

        return instanceName, nil
}


// InstanceType returns the type of the specified node with the specified NodeName.
func (g *Cloud) InstanceType(ctx context.Context, nodeName types.NodeName) (string, error) {
        klog.V(2).Infof("DEBUG:%v - %v\n", "InstanceType", nodeName)
//        instanceName := mapNodeNameToInstanceName(nodeName)
/*
        if g.useMetadataServer {
                // Use metadata, if possible, to fetch ID. See issue #12000
                if g.isCurrentInstance(instanceName) {
                        mType, err := getCurrentMachineTypeViaMetadata()
                        if err == nil {
                                return mType, nil
                        }
                }
        }
        instance, err := g.getInstanceByName(instanceName)
        if err != nil {
                return "", err
        }
*/
//        return instance.Type, nil
        return "", cloudprovider.NotImplemented
}




// InstanceTypeByProviderID returns the cloudprovider instance type of the node
// with the specified unique providerID This method will not be called from the
// node that is requesting this ID. i.e. metadata service and other local
// methods cannot be used here
func (g *Cloud) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
        klog.V(2).Infof("DEBUG:%v - %v\n", "InstanceTypeByProviderID", providerID)
//        instanceName := mapNodeNameToInstanceName(nodeName)

/*
        if g.useMetadataServer {
                // Use metadata, if possible, to fetch ID. See issue #12000
                if g.isCurrentInstance(instanceName) {
                        mType, err := getCurrentMachineTypeViaMetadata()
                        if err == nil {
                                return mType, nil
                        }
                }
        }
        instance, err := g.getInstanceByName(instanceName)
        if err != nil {
                return "", err
        }
        return instance.Type, nil
*/
        return "", cloudprovider.NotImplemented

}



// CurrentNodeName returns the name of the node we are currently running on
// On most clouds (e.g. GCE) this is the hostname, so we provide the hostname
func (g *Cloud) CurrentNodeName(ctx context.Context, hostname string) (types.NodeName, error) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "CurrentNodeName", hostname)
        return types.NodeName(hostname), nil
}




// AddSSHKeyToAllInstances adds an SSH public key as a legal identity for all instances
// expected format for the key is standard ssh-keygen format: <protocol> <blob>
func (g *Cloud) AddSSHKeyToAllInstances(ctx context.Context, user string, keyData []byte) error {
	klog.V(2).Infof("DEBUG:%v - %v\n", "AddSSHKeyToAllInstances", user)
	return nil
}


// InstanceShutdownByProviderID returns true if the instance is in safe state to detach volumes
func (g *Cloud) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "InstanceShutdownByProviderID", providerID)
//        return false, cloudprovider.NotImplemented
	return true, nil
}



// InstanceExistsByProviderID returns true if the instance with the given provider id still exists and is running.
// If false is returned with no error, the instance will be immediately deleted by the cloud controller manager.
func (g *Cloud) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "InstanceExistsByProviderID", providerID)
        return true, nil
}

