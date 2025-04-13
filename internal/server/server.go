package server

import (
	"k8s-mcp-server/internal/cluster"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var clusters []cluster.Cluster

func init() {
	var err error
	clusters, err = cluster.LoadClusters("config/clusters.yaml")
	if err != nil {
		log.Fatalf("Failed to load clusters: %v", err)
	}
}

// Start initializes and starts the HTTP server
func Start() {
	r := gin.Default()

	// Define routes
	r.GET("/clusters", func(c *gin.Context) {
		c.JSON(http.StatusOK, clusters)
	})

	r.GET("/clusters/:name/health", func(c *gin.Context) {
		clusterName := c.Param("name")
		health, err := cluster.GetClusterHealth(clusterName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cluster": clusterName, "health": health})
	})

	r.GET("/clusters/:name/nodes", func(c *gin.Context) {
		clusterName := c.Param("name")
		nodes, err := cluster.GetClusterNodes(clusterName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cluster": clusterName, "nodes": nodes})
	})

	r.GET("/clusters/:name/nodes/:node", func(c *gin.Context) {
		clusterName := c.Param("name")
		nodeName := c.Param("node")
		nodeDetails, err := cluster.GetNodeDetails(clusterName, nodeName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cluster": clusterName, "node": nodeDetails})
	})

	r.GET("/clusters/:name/pods", func(c *gin.Context) {
		clusterName := c.Param("name")
		pods, err := cluster.GetClusterPods(clusterName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cluster": clusterName, "pods": pods})
	})

	r.GET("/clusters/:name/pods/:namespace/:pod", func(c *gin.Context) {
		clusterName := c.Param("name")
		namespace := c.Param("namespace")
		podName := c.Param("pod")
		podDetails, err := cluster.GetPodDetails(clusterName, namespace, podName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cluster": clusterName, "pod": podDetails})
	})

	r.GET("/clusters/:name/deployments", func(c *gin.Context) {
		clusterName := c.Param("name")
		deployments, err := cluster.GetClusterDeployments(clusterName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cluster": clusterName, "deployments": deployments})
	})

	r.GET("/clusters/:name/services", func(c *gin.Context) {
		clusterName := c.Param("name")
		services, err := cluster.GetClusterServices(clusterName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cluster": clusterName, "services": services})
	})

	r.GET("/clusters/:name/events", func(c *gin.Context) {
		clusterName := c.Param("name")
		events, err := cluster.GetClusterEvents(clusterName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cluster": clusterName, "events": events})
	})

	// Start the server
	r.Run() // Default port is 8080
}
