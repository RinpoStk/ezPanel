package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinpostk/ezPanel/embed"
	"github.com/rinpostk/ezPanel/internal/middleware"
	"github.com/rinpostk/ezPanel/internal/service"
)

// SetupRouter creates and configures the Gin engine with all routes.
func SetupRouter(jwtSecret string) *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(gin.Recovery())

	// Services
	monitorSvc := service.NewMonitorCollector()
	ptyMgr := service.NewPtyManager()

	_ = service.NewWinServiceManager() // TODO: wire into service handlers

	// API v1 group
	v1 := r.Group("/api/v1")

	// Public routes (no auth required)
	auth := v1.Group("/auth")
	{
		auth.POST("/login", handleLogin)
		auth.POST("/init", handleInitAdmin)
	}

	// Protected routes (auth required)
	protected := v1.Group("")
	protected.Use(middleware.JWTAuth(jwtSecret))
	{
		// File management
		files := protected.Group("/files")
		{
			files.GET("/list", handleFileList)
			files.GET("/read", handleFileRead)
			files.POST("/write", handleFileWrite)
			files.DELETE("/delete", handleFileDelete)
		}

		// Service management
		services := protected.Group("/services")
		{
			services.GET("", handleServiceList)
			services.GET("/:name", handleServiceStatus)
			services.POST("/:name/start", handleServiceStart)
			services.POST("/:name/stop", handleServiceStop)
		}

		// Terminal WebSocket
		protected.GET("/terminal/ws", func(c *gin.Context) {
			handleTerminalWS(c, ptyMgr)
		})

		// Monitor
		monitor := protected.Group("/monitor")
		{
			monitor.GET("", func(c *gin.Context) {
				handleMonitor(c, monitorSvc)
			})
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Serve embedded frontend (SPA)
	if embed.DistFS != nil {
		fileServer := http.FileServer(http.FS(embed.DistFS))

		r.NoRoute(func(c *gin.Context) {
			// Try to serve the file directly
			path := c.Request.URL.Path
			if len(path) > 1 {
				if f, err := embed.DistFS.Open(path[1:]); err == nil {
					f.Close()
					fileServer.ServeHTTP(c.Writer, c.Request)
					return
				}
			}
			// SPA fallback: serve index.html for non-API routes
			c.Request.URL.Path = "/"
			fileServer.ServeHTTP(c.Writer, c.Request)
		})
	}

	return r
}
