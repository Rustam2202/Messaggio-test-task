package handlers

import (
	"net/http"

	"github.com/Rustam2202/message-processor/kafka"
	"github.com/Rustam2202/message-processor/models"
	"github.com/gin-gonic/gin"
)

type CreateMessageRequest struct {
	Content string `json:"content"`
}

func CreateMessage(c *gin.Context) {
	var req CreateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := models.Message{Content: req.Content, Processed: false}
	if err := models.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	kafka.ProduceMessage(message)

	c.JSON(http.StatusCreated, message)
}

func GetProcessedMessages(c *gin.Context) {
	var messages []models.Message
	if err := models.DB.Where("processed = ?", true).Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func GetMessagesCount(c *gin.Context) {
	var count int64
	if err := models.DB.Model(&models.Message{}).Where("processed = ?", true).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}
