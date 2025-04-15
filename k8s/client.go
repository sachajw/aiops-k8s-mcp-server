package k8s

import (
	"log"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientset  *kubernetes.Clientset
	restConfig *rest.Config
	once       sync.Once
)

// GetClient returns a singleton Kubernetes clientset
func GetClient() *kubernetes.Clientset {
	once.Do(func() {
		// Load kubeconfig
		var err error
		restConfig, err = clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
		if err != nil {
			log.Fatalf("Failed to load kubeconfig: %v", err)
		}

		// Create Kubernetes client
		clientset, err = kubernetes.NewForConfig(restConfig)
		if err != nil {
			log.Fatalf("Failed to create Kubernetes client: %v", err)
		}
	})

	return clientset
}

// GetRESTConfig returns the REST config used to create the clientset
func GetRESTConfig() *rest.Config {
	if restConfig == nil {
		log.Fatalf("REST config is not initialized. Call GetClient() first.")
	}
	return restConfig
}
