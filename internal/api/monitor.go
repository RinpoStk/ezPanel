package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinpostk/ezPanel/internal/service"
)

// handleMonitor returns current system monitoring data.
func handleMonitor(c *gin.Context, collector *service.MonitorCollector) {
	data, err := collector.Collect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to collect metrics"})
		return
	}
	c.JSON(http.StatusOK, data)
}
