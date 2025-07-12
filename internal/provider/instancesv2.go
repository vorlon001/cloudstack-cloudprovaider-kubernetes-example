package provider

import "fmt"
import "context"
import "time"
import "encoding/json"

//import "k8s.io/klog/v2"
import cloudprovider "k8s.io/cloud-provider"
import v1 "k8s.io/api/core/v1"
import "k8s.io/klog/v2"



// InstanceExists returns true if the instance with the given provider id still exists and is running.
// If false is returned with no error, the instance will be immediately deleted by the cloud controller manager.
func (g *Cloud) InstanceExists(ctx context.Context, node *v1.Node) (bool, error) {
        providerID := node.Spec.ProviderID
	_ = providerID

	klog.V(2).Infof("DEBUG:%v - %v\n", "InstanceExists", node)


        klog.V(2).Infof("----------------------------------\n")

        data, err := json.MarshalIndent(node, "", "  ")
        if err != nil {
                fmt.Println("error:", err)
                return false, err
        }
        fmt.Println(string(data))

        klog.V(2).Infof("----------------------------------\n")

        return true, nil
}



func (g *Cloud) nodeAddressesFromInstance(instance string) ([]v1.NodeAddress, error) {

	nodeAddresses := []v1.NodeAddress{}

	var matrix = map[string]string{
		"node130": "192.168.200.130",
		"node131": "192.168.200.131",
		"node132": "192.168.200.132",
		"node133": "192.168.200.133",
		"node140": "192.168.200.140",
		"node141": "192.168.200.141",
		"node142": "192.168.200.142",
		"node143": "192.168.200.143",
		}

	nodeAddresses = append(nodeAddresses, v1.NodeAddress{Type: v1.NodeInternalIP, Address: matrix[instance] })
	// nodeAddresses = append(nodeAddresses, v1.NodeAddress{Type: v1.NodeExternalIP, Address: "1.1.1.1"})

	return nodeAddresses, nil

}

// InstanceMetadata returns metadata of the specified instance.
func (g *Cloud) InstanceMetadata(ctx context.Context, node *v1.Node) (*cloudprovider.InstanceMetadata, error) {
	klog.V(2).Infof("DEBUG:%v - %#v\n", "InstanceMetadata", node)
	klog.V(2).Infof("----------------------------------\n")

        data, err := json.MarshalIndent(node, "", "  ")
        if err != nil {
                fmt.Println("error:", err)
		return nil, err
        }
        fmt.Println(string(data))

	klog.V(2).Infof("----------------------------------\n")
	klog.V(2).Infof("DEBUG:%#v\n", node.Name)

        timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Hour)
	_ = timeoutCtx
        defer cancel()


	var addresses []v1.NodeAddress
	addresses, err = g.nodeAddressesFromInstance(node.Name)
        if err != nil {
                fmt.Println("error:", err)
                return nil, err
        }

	// https://github.com/kubernetes/cloud-provider/blob/master/cloud.go#L300
        return &cloudprovider.InstanceMetadata{
			ProviderID: "iblog-cluster",
			InstanceType: "server-iblog-cluster",
			NodeAddresses: addresses,
			Zone: "cluster56",
			Region: "nodes56",
		}, nil
}


// InstanceShutdown returns true if the instance is in safe state to detach volumes
func (g *Cloud) InstanceShutdown(ctx context.Context, node *v1.Node) (bool, error) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "InstanceShutdown", node)

        klog.V(2).Infof("----------------------------------\n")

        data, err := json.MarshalIndent(node, "", "  ")
        if err != nil {
                fmt.Println("error:", err)
                return false, err
        }
        fmt.Println(string(data))

        klog.V(2).Infof("----------------------------------\n")

	return true, nil
//        return false, cloudprovider.NotImplemented
}

