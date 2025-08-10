package logxyz

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var defaultLogger Logger = Logger{
	output: os.Stdout,
}

type Level int

const (
	DEBUG Level = iota + 1
	INFO
	WARN
	ERROR
)
const MODULE_NAME = "logxyz"

var FormattedANSI = map[string]string{
	"BLUE":   "\033[34m%s\033[0m",
	"GREEN":  "\033[32m%s\033[0m",
	"YELLOW": "\033[33m%s\033[0m",
	"RED":    "\033[31m%s\033[0m",
}

var LevelNames = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

type Logger struct {
	minLevel Level
	output   io.Writer
	webhooks []Webhook
}

type Webhook interface {
	SendMessageToWebhook(logMessage LogMessage)
}

type LogMessage struct {
	Level   Level
	Time    string
	File    string
	Line    int
	Message string
}

func SetOutput(w io.Writer) {
	defaultLogger.output = w
}

func AddWebhooks(wh Webhook) []Webhook {
	defaultLogger.webhooks = append(defaultLogger.webhooks, wh)
	return defaultLogger.webhooks
}

func touchLogLevel(log_level string) Level {
	switch log_level {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	default:
		return DEBUG
	}
}

// [DEBUG=1, INFO=2, WARN=3 ,ERROR=4]
func Default(log_level string) {
	defaultLogger.minLevel = touchLogLevel(log_level)
}

func Debug(message string) {
	LogProcess(1, message)
}

func Info(message string) {
	LogProcess(2, message)
}

func Warn(message string) {
	LogProcess(3, message)
}

func Error(message string) {
	LogProcess(4, message)
}

func LogProcess(level Level, message string) {
	logMessage := &LogMessage{}
	logMessage.Init(level, message)
	logMessage.FilePreProcess(2)
	logMessage.Print()
}

func (l *LogMessage) WebhookProcess() {
	for _, webhook := range defaultLogger.webhooks {
		webhook.SendMessageToWebhook(*l)
	}
}

func (l *LogMessage) Init(level Level, message string) {
	l.Level = level
	l.Message = message
	l.Time = time.Now().Format("2006-01-02 15:04:05")
}

func (l *LogMessage) FilePreProcess(depth int) {
	if depth < 1 {
		depth = 1
	}

	_, file, line := GetLineFromCalledFunction()

	files := strings.Split(file, "/")
	if len(files) >= depth {
		file = strings.Join(files[(len(files)-depth):], "/")
	}

	l.File = file
	l.Line = line
}

func (l *LogMessage) Print() {
	if l.Level < defaultLogger.minLevel {
		return
	}

	var colorLevel string
	var colorMessage string

	switch LevelNames[l.Level] {
	case "INFO":
		colorLevel = fmt.Sprintf(FormattedANSI["BLUE"], "INF")
		colorMessage = fmt.Sprintf(FormattedANSI["BLUE"], l.Message)
	case "DEBUG":
		colorLevel = fmt.Sprintf(FormattedANSI["GREEN"], "DBG")
		colorMessage = fmt.Sprintf(FormattedANSI["GREEN"], l.Message)
	case "WARN":
		colorLevel = fmt.Sprintf(FormattedANSI["YELLOW"], "WRN")
		colorMessage = fmt.Sprintf(FormattedANSI["YELLOW"], l.Message)
	case "ERROR":
		colorLevel = fmt.Sprintf(FormattedANSI["RED"], "ERR")
		colorMessage = fmt.Sprintf(FormattedANSI["RED"], l.Message)
	}

	formattedMessage := fmt.Sprintf("%s %s \033[1;30m%s:%d\033[0m %s", colorLevel, l.Time, l.File, l.Line, colorMessage)
	defaultLogger.output.Write([]byte(formattedMessage + "\n"))

	l.WebhookProcess()
}

func GetLineFromCalledFunction() (functionName, fileName string, line int) {
	// program counters
	pcs := make([]uintptr, 10)
	n := runtime.Callers(2, pcs)

	pcs = pcs[:n]
	frames := runtime.CallersFrames(pcs)

	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.Function, MODULE_NAME) {
			return frame.Function, frame.File, frame.Line
		}

		if !more {
			break
		}
	}

	return "unknown", "unknown", 0
}
