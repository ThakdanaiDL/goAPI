package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Send(message string) error {
	// webhook := os.Getenv("DISCORD_WEBHOOK_URL")
	webhook := "https://discordapp.com/api/webhooks/1452582113042763809/EBlk05Ydp8WRxjS-X8j0PQ1G_6At-voEiDxoyU92eki2Z1hQdYfUBdrZvW5wDEOB-DAv"

	body := map[string]string{"content": message}
	jsonBody, _ := json.Marshal(body)

	_, err := http.Post(webhook, "application/json", bytes.NewBuffer(jsonBody))
	return err
}
