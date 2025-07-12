package provider

import "context"
import "k8s.io/klog/v2"


// ListClusters will return a list of cluster names for the associated project
func (g *Cloud) ListClusters(ctx context.Context) ([]string, error) {
        allClusters := []string{}

	klog.V(2).Infof("DEBUG:%v - %v\n", "ListClusters", "-")

/*
        for _, zone := range g.managedZones {
                clusters, err := g.listClustersInZone(zone)
                if err != nil {
                        return nil, err
                }
                // TODO: Scoping?  Do we need to qualify the cluster name?
                allClusters = append(allClusters, clusters...)
        }
*/

        return allClusters, nil
}


// Master returned the dns address of the master
func (g *Cloud) Master(ctx context.Context, clusterName string) (string, error) {
	klog.V(2).Infof("DEBUG:%v - %v\n", "Master", clusterName)
        return "k8s-" + clusterName + "-master.internal", nil
}



