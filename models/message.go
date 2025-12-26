package model

import (
	"gorm.io/gorm"
)

type MessageLog struct {
	gorm.Model
	Content string `json:"content"`
	Status  string `json:"status"`
}

type UserData struct {
	gorm.Model
	Name    string `json:"name"`
	Winrate string `json:"winrate"`
	Rank    string `json:"rank"`
}
