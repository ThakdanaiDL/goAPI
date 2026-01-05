package model

import "gorm.io/gorm"

type MessageLog struct {
	gorm.Model
	Content string `json:"content"`
	Status  string `json:"status"`
}

type UserData struct {
	gorm.Model
	Name           string `json:"name"`
	Rank           string `json:"rank"`
	TotalGames     int    `json:"totalGame"  gorm:"default:0"`
	CurrentCourtID *int   `json:"currentCourtId"`
	LastPartnerID  *uint  `json:"lastPartnerId"`
}

type UserSummary struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Rank string `json:"rank"`
}
