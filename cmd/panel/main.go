package main

import (
	"fmt"

	"github.com/rinpostk/ezPanel/internal/api"
	"github.com/rinpostk/ezPanel/internal/config"
	"github.com/rinpostk/ezPanel/internal/logger"
)

func main() {
	cfg := config.Load()

	r := api.SetupRouter(cfg.JWTSecret)

	addr := fmt.Sprintf(":%s", cfg.Port)
	logger.Info("ezPanel starting on %s", addr)
	if err := r.Run(addr); err != nil {
		logger.Fatal("failed to start server: %v", err)
	}
}
