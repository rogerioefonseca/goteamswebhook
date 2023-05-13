package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type Payload struct {
	Text string `json:"text"`
}

func main() {

	if os.Getenv("TEAMS_WEBHOOK_URL") == "" {
		panic("TEAMS_WEBHOOK_URL environment variable is not set")
	}

	if len(os.Args) < 2 {
		panic("Please provide the message text as a command-line argument")
	}

	webhookURL := os.Getenv("TEAMS_WEBHOOK_URL")
	messageText := os.Args[1]

	payload := Payload{
		Text: "**[IMPORTANT]** " + messageText,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("Failed to send message:" + resp.Status)
	}
	println("Message Sent")
}
