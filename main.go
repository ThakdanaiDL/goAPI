package main

import (
	"net/http"
	"os"

	"bytes"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

		err := sendDiscordNotify(msg) // เปลี่ยนมาเรียกใช้ Discord
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ส่งไม่สำเร็จ"})
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "ส่งเข้า Discord แล้ว!"})
	})

	// สำคัญ: Cloud Run จะกำหนด Port ผ่าน Environment Variable ชื่อ PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//ddd
	e.Logger.Fatal(e.Start(":" + port))
}

// อย่าลืม import "bytes" และ "encoding/json" นะครับ

func sendDiscordNotify(message string) error {
	// แนะนำให้ใช้ os.Getenv("DISCORD_WEBHOOK_URL") เพื่อความปลอดภัย
	webhookURL := "https://discordapp.com/api/webhooks/1452582113042763809/EBlk05Ydp8WRxjS-X8j0PQ1G_6At-voEiDxoyU92eki2Z1hQdYfUBdrZvW5wDEOB-DAv"

	// โครงสร้าง JSON ที่ Discord ต้องการ
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
