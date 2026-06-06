package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleServiceList lists all services.
func handleServiceList(c *gin.Context) {
	// TODO: implement service listing
	c.JSON(http.StatusOK, gin.H{"message": "TODO: service list"})
}

// handleServiceStatus returns the status of a specific service.
func handleServiceStatus(c *gin.Context) {
	name := c.Param("name")
	// TODO: implement service status
	c.JSON(http.StatusOK, gin.H{"message": "TODO: service status", "service": name})
}

// handleServiceStart starts a service.
func handleServiceStart(c *gin.Context) {
	name := c.Param("name")
	// TODO: implement service start
	c.JSON(http.StatusOK, gin.H{"message": "TODO: service start", "service": name})
}

// handleServiceStop stops a service.
func handleServiceStop(c *gin.Context) {
	name := c.Param("name")
	// TODO: implement service stop
	c.JSON(http.StatusOK, gin.H{"message": "TODO: service stop", "service": name})
}
