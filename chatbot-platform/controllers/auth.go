package controllers

import (
	"chatbot-platform/database"
	"chatbot-platform/models"
	"chatbot-platform/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dữ liệu không hợp lệ"})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tên đăng nhập đã tồn tại!"})
		return
	}

	newUser := models.User{
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}

	database.DB.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{"message": "Đăng ký thành công", "result": newUser})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dữ liệu không hợp lệ"})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Sai tên đăng nhập hoặc mật khẩu"})
		return
	}

	// [SERVICE CALL] Sinh Token JWT và trả về cho Client
	tokenString, err := services.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Không thể thiết lập Token bảo vệ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Đăng nhập thành công",
		"token":   tokenString, // Trả token về
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Select("id", "username", "email", "created_at").Find(&users)
	c.JSON(http.StatusOK, users)
}
