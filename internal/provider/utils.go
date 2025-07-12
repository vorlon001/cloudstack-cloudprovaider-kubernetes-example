package provider

import "fmt"

//import "k8s.io/klog/v2"
import "k8s.io/apimachinery/pkg/types"
import compute "google.golang.org/api/compute/v1"
import v1 "k8s.io/api/core/v1"



// mapNodeNameToInstanceName maps a k8s NodeName to a GCE Instance Name
// This is a simple string cast.
func mapNodeNameToInstanceName(nodeName types.NodeName) string {
        return string(nodeName)
}

func nodeAddressesFromInstance(instance *compute.Instance) ([]v1.NodeAddress, error) {
        if len(instance.NetworkInterfaces) < 1 {
                return nil, fmt.Errorf("could not find network interfaces for instanceID %q", instance.Id)
        }
        nodeAddresses := []v1.NodeAddress{}

        return nodeAddresses, nil
}


