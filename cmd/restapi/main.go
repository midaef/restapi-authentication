package main

import (
	"packages/internal/app/server"

	"go.uber.org/zap"
)

func main() {

	config := server.NewConfig("resources/configs/config.json")
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Starting server",
		zap.String("host", config.ServerHost),
		zap.String("port", config.Port),
		zap.String("log_level", config.LogLevel))
	err := server.NewApiServer(config)
	if err != nil {
		logger.Error("ListenAndServe", zap.Error(err))
	}
}
