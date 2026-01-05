package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func Send(message string) error {
	webhook := os.Getenv("DISCORD_WEBHOOK_URL")

	body := map[string]string{"content": message}
	jsonBody, _ := json.Marshal(body)

	_, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonBody))
	return err
}
