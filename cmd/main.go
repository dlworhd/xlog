package main

import (
	"os"

	"github.com/dlworhd/xlog/xlog"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env.local")
	log_level := os.Getenv("LOG_LEVEL")
	discord_webhook_url := os.Getenv("DISCORD_WEBHOOK_URL")

	xlog.Default(log_level)
	xlog.AddWebhook(xlog.WebhookClient{Name: "DISCORD", Url: discord_webhook_url})

	xlog.Info("LOG TEST")
}
