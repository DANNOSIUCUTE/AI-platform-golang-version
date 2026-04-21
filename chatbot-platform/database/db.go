package database

import (
	"fmt"
	"log"
	"os"

	"chatbot-platform/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Môi trường Docker sẽ thay thế các thông số này, trên localhost thì chạy mặc định
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", ""), // Trên local thường rỗng, trên docker sẽ là rootpassword
		getEnv("DB_HOST", "127.0.0.1"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "chatbot_platform"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể kết nối đến Database: \n", err)
	}

	log.Println("Kết nối Database MySQL (GORM) thành công!")

	// GORM Tự động nhận diện cấu trúc Go Struct và thiết lập/cập nhật bảng MySQL
	err = db.AutoMigrate(&models.User{}, &models.ChatHistory{})
	if err != nil {
		log.Fatal("Lỗi lúc Auto-Migrate: \n", err)
	}

	DB = db
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
