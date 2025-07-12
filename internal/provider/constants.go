package provider

import "time"
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
        // ProviderName is the official const representation of the Google Cloud Provider
        ProviderName = "iblogCloud"
)


const (
        // UIDConfigMapName is the Key used to persist UIDs to configmaps.
        UIDConfigMapName = "ingress-uid"

        // UIDNamespace is the namespace which contains the above config map
        UIDNamespace = metav1.NamespaceSystem

        // UIDCluster is the data keys for looking up the clusters UID
        UIDCluster = "uid"

        // UIDProvider is the data keys for looking up the providers UID
        UIDProvider = "provider-uid"

        // UIDLengthBytes is the length of a UID
        UIDLengthBytes = 8

        // Frequency of the updateFunc event handler being called
        // This does not actually query the apiserver for current state - the local cache value is used.
        updateFuncFrequency = 10 * time.Minute
)

