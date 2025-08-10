package xlog

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"
)

type SlackNotifier struct {
	WebhookUrl string
}

type DiscordNotifier struct {
	WebhookUrl string
}

//

type WebhookClient struct {
	Name WebhookClientType
	Url  string
}

type WebhookClientType string

const (
	DISCORD WebhookClientType = "DISCORD"
	SLACK   WebhookClientType = "SLACK"
)

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
		var color int

		switch logMessage.Level {
		case "DEBUG":
			color = 0x31CD31
		case "INFO":
			color = 0x3131CD
		case "WARN":
			color = 0xCDCD31
		case "ERROR":
			color = 0xCD3131
		}

		msg := DiscordMessage{
			Embeds: []DiscordEmbed{
				{
					Fields: []DiscordEmbedField{
						{Name: "Level", Value: string(logMessage.Level), Inline: true},
						{Name: "Date", Value: logMessage.Time, Inline: true},
						{Name: "Message", Value: logMessage.Message, Inline: false},
					},
					Color: color,
				},
			},
		}
		message, _ = json.Marshal(msg)
	case "SLACK":
		encodedMessage, _ := json.Marshal(map[string]string{"text": logMessage.Message})
		message = []byte(encodedMessage)

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
