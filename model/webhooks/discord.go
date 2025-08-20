package webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	logger "github.com/dlworhd/logger/model"
)

type DiscordNotifier struct {
	WebhookUrl string
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

type Sender interface {
	Send(logMessage *logger.LogMessage) error
}

// Log Level, Time, Message가 담긴 객체를 message대신 넣기
func (n *DiscordNotifier) SendMessageToWebhook(logMessage logger.LogMessage) {

	var color int

	switch logger.LevelNames[logMessage.Level] {
	case "DEBUG":
		color = 0x31CD31
	case "INFO":
		color = 0x3131CD
	case "WARN":
		color = 0xCDCD31
	case "ERROR":
		color = 0xCD3131
	}

	message := DiscordMessage{
		Embeds: []DiscordEmbed{
			{
				Fields: []DiscordEmbedField{
					{Name: "Level", Value: logger.LevelNames[logMessage.Level], Inline: true},
					{Name: "Date", Value: logMessage.Time, Inline: true},
					{Name: "Message", Value: fmt.Sprintln(logMessage.Messages...), Inline: false},
				},
				Color: color,
			},
		},
	}

	encodedMessage, _ := json.Marshal(message)

	resp, err := http.Post(n.WebhookUrl, "application/json", bytes.NewBuffer(encodedMessage))

	if err != nil {
		// fmt.Println("Error 발생!!!!", err.Error())
	}
	defer resp.Body.Close()
}
