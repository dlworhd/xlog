package main

import (
	"time"

	"github.com/dlworhd/xlog/xlog"
)

func main() {
	xlog.Default("ERROR")

	for {
		xlog.Info("INFO TEST")
		xlog.Error("ERROR TEST")
		xlog.Debug("DEBUG TEST")
		xlog.Warn("WARN TEST")

		time.Sleep(1 * time.Second)
	}

}
