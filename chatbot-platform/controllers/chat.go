package controllers

import (
	"net/http"
	"os"

	"chatbot-platform/database"
	"chatbot-platform/models"
	"chatbot-platform/services"

	"github.com/gin-gonic/gin"
)

type ChatInput struct {
	Message string `json:"message" binding:"required"`
}

func ProcessChat(c *gin.Context) {
	chatbotID := c.Param("id")
	var input ChatInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Sai định dạng văn bản"})
		return
	}

	apiKey := os.Getenv("CLAUDE_API_KEY")

	// [SERVICE CALL] Gom phần logic tạo Request sang file claude_service.go
	botResponse, err := services.GetClaudeResponse(apiKey, input.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Lỗi API AI", "error": err.Error()})
		return
	}

	// Lưu DB qua GORM
	history := models.ChatHistory{
		ChatbotID:   chatbotID,
		UserMessage: input.Message,
		BotResponse: botResponse,
	}
	database.DB.Create(&history)

	c.JSON(http.StatusOK, gin.H{"success": true, "message": botResponse})
}

func GetHistory(c *gin.Context) {
	chatbotID := c.Param("id")
	var history []models.ChatHistory

	database.DB.Where("chatbot_id = ?", chatbotID).Order("timestamp DESC").Limit(50).Find(&history)

	c.JSON(http.StatusOK, gin.H{"success": true, "history": history})
}
