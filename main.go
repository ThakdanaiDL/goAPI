package main

import (
	"log"
	"net/http"
	"os"

	"bytes"
	"encoding/json"

	"github.com/labstack/echo/v4"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MessageLog struct {
	gorm.Model
	Content string `json:"content"`
	Status  string `json:"status"`
}

var db *gorm.DB

func main() {

	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	// 3. สั่งให้สร้าง Table อัตโนมัติ (Auto Migrate)
	db.AutoMigrate(&MessageLog{})

	e := echo.New()

	// 4. API สำหรับดูข้อมูลทั้งหมด (Read)
	e.GET("/history", func(c echo.Context) error {
		var logs []MessageLog
		db.Find(&logs)
		return c.JSON(http.StatusOK, logs)
	})

	// // Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Health Check Route
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "UP",
			"message": "Echo API is running on Cloud Run!",
		})
	})

	e.GET("/send", func(c echo.Context) error {
		msg := c.QueryParam("msg")
		if msg == "" {
			msg = "สวัสดีจาก Go API บน Railway! 🚀"
		}

		err := sendDiscordNotify(msg)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ส่งไม่สำเร็จ"})
		}

		db.Create(&MessageLog{Content: msg, Status: "Sent"})

		return c.JSON(http.StatusOK, map[string]string{"status": "ส่งเข้า Discord แล้ว!"})
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to My Go API on Render!")
	})

	// สำคัญ: Cloud Run จะกำหนด Port ผ่าน Environment Variable ชื่อ PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//ddd
	e.Logger.Fatal(e.Start(":" + port))
}

func sendDiscordNotify(message string) error {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL") //แนะนำให้ใช้ เพื่อความปลอดภัย

	payload := map[string]string{
		"content": message,
	}
	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
