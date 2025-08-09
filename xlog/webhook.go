package xlog

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

type WebhookClient struct {
	Name WebhookClientType
	Url  string
}

type WebhookClientType string

const (
	DISCORD WebhookClientType = "DISCORD"
	SLACK   WebhookClientType = "SLACK"
)

type MessageSender interface {
	Send(*LogMessage)
}

type DiscordEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type DiscordEmbed struct {
	Fields []DiscordEmbedField `json:"fields"`
	Color  int                 `json:"color"`
}

type DiscordMessage struct {
	Embeds []DiscordEmbed `json:"embeds"`
}

// Log Level, Time, Message가 담긴 객체를 message대신 넣기
func SendMessageToWebhook(webhook WebhookClient, logMessage *LogMessage, wg *sync.WaitGroup) {
	defer wg.Done()

	var message []byte
	switch webhook.Name {
	case "DISCORD":
		msg := DiscordMessage{
			Embeds: []DiscordEmbed{
				{
					Fields: []DiscordEmbedField{
						{Name: "Level", Value: string(logMessage.Level), Inline: true},
						{Name: "Date", Value: logMessage.Time, Inline: true},
						{Name: "Message", Value: logMessage.Message, Inline: false},
					},
					Color: 0xff0000,
				},
			},
		}
		message, _ = json.Marshal(msg)
	default:
		return
	}

	resp, err := http.Post(webhook.Url, "application/json", bytes.NewBuffer(message))

	if err != nil {
		Info(err.Error())
	}
	defer resp.Body.Close()
}

func WebhooksProcess(logMessage *LogMessage) {
	var wg sync.WaitGroup
	wg.Add(len(defaultLogger.Webhooks))

	for _, webhook := range defaultLogger.Webhooks {
		go SendMessageToWebhook(webhook, logMessage, &wg)
	}

	wg.Wait()
}
