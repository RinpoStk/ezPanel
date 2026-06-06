package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleLogin handles user login.
func handleLogin(c *gin.Context) {
	// TODO: validate credentials, issue JWT
	c.JSON(http.StatusOK, gin.H{"message": "TODO: login"})
}

// handleInitAdmin handles initial admin account setup.
func handleInitAdmin(c *gin.Context) {
	// TODO: create admin user if none exists
	c.JSON(http.StatusOK, gin.H{"message": "TODO: init admin"})
}
