package main

import (
	"github.com/dlworhd/xlog/xlog"
)

func main() {
	xlog.Default("DBG")

	xlog.Info("INFO TEST")
	xlog.Debug("DEBUG TEST")
	xlog.Warn("WARN TEST")
	xlog.Error("ERROR TEST")

}
