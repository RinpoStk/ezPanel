package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleFileList lists files in a directory.
func handleFileList(c *gin.Context) {
	// TODO: implement directory listing
	c.JSON(http.StatusOK, gin.H{"message": "TODO: file list", "path": c.Query("path")})
}

// handleFileRead reads file content.
func handleFileRead(c *gin.Context) {
	// TODO: implement file reading
	c.JSON(http.StatusOK, gin.H{"message": "TODO: file read", "path": c.Query("path")})
}

// handleFileWrite writes content to a file.
func handleFileWrite(c *gin.Context) {
	// TODO: implement file writing
	c.JSON(http.StatusOK, gin.H{"message": "TODO: file write"})
}

// handleFileDelete deletes a file.
func handleFileDelete(c *gin.Context) {
	// TODO: implement file deletion
	c.JSON(http.StatusOK, gin.H{"message": "TODO: file delete", "path": c.Query("path")})
}
