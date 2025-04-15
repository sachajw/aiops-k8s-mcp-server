package k8s

import (
	"log"
	"os"
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
		kubeconfigPath := os.Getenv("KUBECONFIG")
		if kubeconfigPath == "" {
			kubeconfigPath = clientcmd.RecommendedHomeFile // Use default if not specified
		}

		var err error
		restConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			log.Fatalf("Failed to load kubeconfig from %s: %v", kubeconfigPath, err)
		}

		// Create Kubernetes client
		clientset, err = kubernetes.NewForConfig(restConfig)
		if err != nil {
			log.Fatalf("Failed to create Kubernetes client: %v", err)
		}
	})

	return clientset
}
