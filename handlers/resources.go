package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"k8s-mcp-server/k8s"
	"k8s.io/client-go/discovery"
)

func ListResources(w http.ResponseWriter, r *http.Request) {
	// Get Kubernetes client and REST config
	restConfig := k8s.GetRESTConfig()

	// Create discovery client
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(restConfig)
	if err != nil {
		http.Error(w, "Failed to create discovery client", http.StatusInternalServerError)
		log.Printf("Error creating discovery client: %v", err)
		return
	}

	// List all resources
	apiResourceLists, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		http.Error(w, "Failed to list resources", http.StatusInternalServerError)
		log.Printf("Error listing resources: %v", err)
		return
	}

	var resources []string
	for _, apiResourceList := range apiResourceLists {
		for _, resource := range apiResourceList.APIResources {
			resources = append(resources, resource.Name)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}
