package main

import (
	"os"

	"github.com/dlworhd/xlog/xlog"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env.local")
	log_level := os.Getenv("LOG_LEVEL")

	xlog.Default(log_level)

	xlog.Debug("DEBUG TEST")
	xlog.Info("INFO TEST")
	xlog.Warn("WARN TEST")
	xlog.Error("ERROR TEST")
}
