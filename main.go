package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"k8s-mcp-server/pkg/k8s" // Import the k8s package
)

func main() {
	// MCP server
	s := server.NewMCPServer(
		"Kubernetes Tools",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
		server.WithRecovery(),
	)

	// Get Kubernetes client from the k8s package
	clientset := k8s.GetClient()
	if clientset == nil {
		log.Fatalf("Failed to get Kubernetes client from k8s package")
		return
	}

	// Add a list_pods tool
	listPodsTool := mcp.NewTool("list_pods",
		mcp.WithDescription("List all pods in the cluster"),
	)

	// Add the list_pods handler
	s.AddTool(listPodsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

	// Add a list_nodes tool
	listNodesTool := mcp.NewTool("list_nodes",
		mcp.WithDescription("List all nodes in the cluster"),
	)

	// Add the list_nodes handler
	s.AddTool(listNodesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Printf("Failed to list nodes: %v", err)
			return nil, errors.New("Failed to list nodes")
		}

		nodeNames := ""
		for _, node := range nodes.Items {
			nodeNames += fmt.Sprintf("- %s\n", node.Name)
		}

		return mcp.NewToolResultText(nodeNames), nil
	})

	// Add a list_namespaces tool
	listNamespacesTool := mcp.NewTool("list_namespaces",
		mcp.WithDescription("List all namespaces in the cluster"),
	)

	// Add the list_namespaces handler
	s.AddTool(listNamespacesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Printf("Failed to list namespaces: %v", err)
			return nil, errors.New("Failed to list namespaces")
		}

		namespaceNames := ""
		for _, namespace := range namespaces.Items {
			namespaceNames += fmt.Sprintf("- %s\n", namespace.Name)
		}

		return mcp.NewToolResultText(namespaceNames), nil
	})

	// Add a list_configmaps tool
	listConfigMapsTool := mcp.NewTool("list_configmaps",
		mcp.WithDescription("List all configmaps in the cluster"),
	)

	// Add the list_configmaps handler
	s.AddTool(listConfigMapsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		configmaps, err := clientset.CoreV1().ConfigMaps("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Printf("Failed to list configmaps: %v", err)
			return nil, errors.New("Failed to list configmaps")
		}

		configmapNames := ""
		for _, configmap := range configmaps.Items {
			configmapNames += fmt.Sprintf("- %s/%s\n", configmap.Namespace, configmap.Name)
		}

		return mcp.NewToolResultText(configmapNames), nil
	})

	// Add a list_services tool
	listServicesTool := mcp.NewTool("list_services",
		mcp.WithDescription("List all services in the cluster"),
	)

	s.AddTool(listServicesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Printf("Failed to list services: %v", err)
			return nil, errors.New("Failed to list services")
		}

		serviceNames := ""
		for _, service := range services.Items {
			serviceNames += fmt.Sprintf("- %s/%s\n", service.Namespace, service.Name)
		}

		return mcp.NewToolResultText(serviceNames), nil
	})

	listIngressesTool := mcp.NewTool("list_ingresses",
		mcp.WithDescription("List all ingresses in the cluster"),
	)

	s.AddTool(listIngressesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		ingresses, err := clientset.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Printf("Failed to list ingresses: %v", err)
			return nil, errors.New("Failed to list ingresses")
		}

		ingressNames := ""
		for _, ingress := range ingresses.Items {
			ingressNames += fmt.Sprintf("- %s/%s\n", ingress.Namespace, ingress.Name)
		}

		return mcp.NewToolResultText(ingressNames), nil
	})

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
