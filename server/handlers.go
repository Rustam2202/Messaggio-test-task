package server

import (
	"net/http"

	"github.com/Rustam2202/message-processor/kafka"
	"github.com/Rustam2202/message-processor/logger"
	"github.com/Rustam2202/message-processor/models"
	"github.com/gin-gonic/gin"
)

type CreateMessageRequest struct {
	Content string `json:"content"`
}

func CreateMessage(c *gin.Context) {
	var req CreateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error().Caller().Err(err).Msg("Failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := models.Message{Content: req.Content, Processed: false}
	if err := models.DB.Create(&message).Error; err != nil {
		logger.Logger.Error().Caller().Err(err).Msg("Failed to create message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	kafka.ProduceMessage(message)

	logger.Logger.Debug().Msgf("Message created: %+v", message)
	c.JSON(http.StatusCreated, message)
}

func GetProcessedMessages(c *gin.Context) {
	var messages []models.Message
	if err := models.DB.Where("processed = ?", true).Find(&messages).Error; err != nil {
		logger.Logger.Error().Caller().Err(err).Msg("Failed to get processed messages")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Debug().Msgf("Processed messages retrieved: %+v", messages)
	c.JSON(http.StatusOK, messages)
}

func GetMessagesCount(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.Message{}).Where("processed = ?", true).Count(&count).Error; err != nil {
		logger.Logger.Error().Caller().Err(err).Msg("Failed to get messages count")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Logger.Debug().Msgf("Messages count retrieved: %d", count)
	c.JSON(http.StatusOK, gin.H{"count": count})
}
