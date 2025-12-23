package model

import (
	"gorm.io/gorm"
)

type MessageLog struct {
	gorm.Model
	Content string `json:"content"`
	Status  string `json:"status"`
}
