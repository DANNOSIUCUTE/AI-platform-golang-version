package middlewares

import (
	"net/http"
	"strings"

	"chatbot-platform/services"

	"github.com/gin-gonic/gin"
)

// Middleware này được gắn trước bất kì Request nào cần bảo mật
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Yêu cầu Client gửi header: "Authorization: Bearer <token>"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Vui lòng đăng nhập để sử dụng tính năng này (Thiếu Token)"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Định dạng Token không đứng: Yêu cầu bắt đầu bằng 'Bearer'"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Validate Token với Secret Key
		claims, err := services.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ hoặc đã hết hạn"})
			c.Abort()
			return
		}

		// Setup biến username sang context của nhánh request
		c.Set("username", claims.Username)

		// Tất cả Ok, đi vào Controller Logic (vd: chat)
		c.Next()
	}
}
