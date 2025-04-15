package handlers

import (
    "encoding/json"
    "log"
    "net/http"
	"k8s-mcp-server/k8s"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListNamespaces(w http.ResponseWriter, r *http.Request) {
    // Get Kubernetes client
    clientset := k8s.GetClient()

    // List namespaces
    namespaces, err := clientset.CoreV1().Namespaces().List(r.Context(), metav1.ListOptions{})
    if err != nil {
        http.Error(w, "Failed to list namespaces", http.StatusInternalServerError)
        log.Printf("Error listing namespaces: %v", err)
        return
    }

    // Prepare response
    var namespaceNames []string
    for _, ns := range namespaces.Items {
        namespaceNames = append(namespaceNames, ns.Name)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(namespaceNames)
}