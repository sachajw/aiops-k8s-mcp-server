package routes

import (
    "github.com/gorilla/mux"
    "k8s-mcp-server/handlers"
)

func InitializeRouter() *mux.Router {
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/namespaces", handlers.ListNamespaces).Methods("GET")
    r.HandleFunc("/pods", handlers.ListPods).Methods("GET")
    r.HandleFunc("/resources", handlers.ListResources).Methods("GET")

    return r
}