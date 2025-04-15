package handlers

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "strings"
	"k8s-mcp-server/k8s"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPods(w http.ResponseWriter, r *http.Request) {
    // Parse query parameters
    namespaces := r.URL.Query().Get("namespaces") // Comma-separated namespaces

    // Get Kubernetes client
    clientset := k8s.GetClient()

    // List pods
    var podList []string
    nsList := strings.Split(namespaces, ",")
    for _, ns := range nsList {
        pods, err := clientset.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
        if err != nil {
            http.Error(w, "Failed to list pods", http.StatusInternalServerError)
            log.Printf("Error listing pods in namespace %s: %v", ns, err)
            return
        }

        for _, pod := range pods.Items {
            podList = append(podList, ns+"/"+pod.Name)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(podList)
}