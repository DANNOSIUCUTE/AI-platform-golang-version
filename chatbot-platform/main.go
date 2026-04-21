package main

import (
	"log"

	"chatbot-platform/controllers"
	"chatbot-platform/database"
	"chatbot-platform/middlewares" // Thêm đường dẫn tới middleware

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	database.ConnectDB()

	r := gin.Default()

	r.Static("/public", "./public")
	r.Static("/views", "./views")
	r.Static("/images", "./public/images")

	r.GET("/", func(c *gin.Context) { c.File("./views/login.html") })
	r.GET("/login", func(c *gin.Context) { c.File("./views/login.html") })
	r.GET("/dashboard", func(c *gin.Context) { c.File("./views/dashboard.html") })

	api := r.Group("/api")
	{
		// Cổng mở (không check Token)
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Vùng Cấm (Phải có Token hợp lệ mới truy cập được)
		protected := api.Group("")
		protected.Use(middlewares.JWTAuth())
		{
			protected.GET("/users", controllers.GetUsers)
			
			// Chatbot Service cần độ chính danh thông qua Token
			protected.POST("/chatbots/:id/chat", controllers.ProcessChat)
			protected.GET("/chatbots/:id/history", controllers.GetHistory)
		}
	}

	log.Println("Golang Server đang khởi chạy localhost:3000")
	r.Run(":3000")
}
