package server

import (
	"github.com/Rustam2202/message-processor/logger"
	"github.com/gin-gonic/gin"
)

func Run(hostPort string) {
	server := gin.Default()

	api := server.Group("/api")
	api.POST("/messages", CreateMessage)
	api.GET("/messages/processed", GetProcessedMessages)
	api.GET("/messages/count", GetMessagesCount)

	if err := server.Run(hostPort); err != nil {
		logger.Logger.Fatal().Caller().Err(err).Msg("Failed to start server")
	}
}
