package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rinpostk/ezPanel/internal/service"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: restrict origins in production
	},
}

// handleTerminalWS handles WebSocket terminal connections.
func handleTerminalWS(c *gin.Context, ptyMgr *service.PtyManager) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "websocket upgrade failed"})
		return
	}
	defer conn.Close()

	// TODO: spawn PTY and wire up read/write
	_ = ptyMgr

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		// Echo back for now
		_ = messageType
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}
