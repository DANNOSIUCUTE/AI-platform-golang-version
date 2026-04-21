package models

import "time"

type ChatHistory struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ChatbotID   string    `gorm:"size:50;not null" json:"chatbotId"`
	UserMessage string    `gorm:"type:text;not null" json:"userMessage"`
	BotResponse string    `gorm:"type:text;not null" json:"botResponse"`
	Timestamp   time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
