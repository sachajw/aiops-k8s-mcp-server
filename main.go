package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"k8s-mcp-server/pkg/k8s" 
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Kubernetes Tools",
		"0.0.1",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	// Add a list_pods tool
	listPodsTool := mcp.NewTool("list_pods",
		mcp.WithDescription("List all pods in the cluster"),
	)

	// Add the list_pods handler
	s.AddTool(listPodsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Get Kubernetes client from the k8s package
		clientset := k8s.GetClient()
		if clientset == nil {
			log.Printf("Failed to get Kubernetes client from k8s package")
			return nil, errors.New("Failed to get Kubernetes client")
		}
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Printf("Failed to list pods: %v", err)
			return nil, errors.New("Failed to list pods")
		}

		podNames := ""
		for _, pod := range pods.Items {
			podNames += fmt.Sprintf("- %s/%s\n", pod.Namespace, pod.Name)
		}

		return mcp.NewToolResultText(podNames), nil
	})

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	fmt.Println("server started on port 8080")
	}
}
