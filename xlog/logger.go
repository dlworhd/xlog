package xlog

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"
)

const MODULE_NAME = "xlog"

var defaultLogger Logger = Logger{}

type Level string

const (
	DEBUG Level = "DEBUG"
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
)

var levelPriority = map[Level]int{
	DEBUG: 1,
	INFO:  2,
	WARN:  3,
	ERROR: 4,
}

type Logger struct {
	minLevel Level
	Time     string
	File     string
	Line     int
	Webhooks []WebhookClient
}

type LogMessage struct {
	Level   Level
	Time    string
	File    string
	Line    int
	Message string
}

func touchLogLevel(log_level string) Level {
	switch log_level {
	case "DEBUG":
	case "INFO":
	case "WARN":
	case "ERROR":
	default:
		log_level = "DEBUG"
	}

	return Level(log_level)
}

// Log Level = [DEBUG, INFO, WARN ,ERROR]
func Default(log_level string) {
	defaultLogger.minLevel = touchLogLevel(log_level)

}

func Info(message string) {
	LogProcess("INFO", message)
}

func Warn(message string) {
	LogProcess("WARN", message)
}

func Error(message string) {
	LogProcess("ERROR", message)
}

func Debug(message string) {
	LogProcess("DEBUG", message)
}

func LogProcess(level Level, message string) {

	logMessage := &LogMessage{
		Level:   level,
		Message: message,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}

	Init(logMessage)
	Log(logMessage)
}

func Log(logMessage *LogMessage) {

	if levelPriority[logMessage.Level] < levelPriority[defaultLogger.minLevel] {
		return
	}

	var colorLevel string
	var colorMessage string

	switch logMessage.Level {
	case "INFO":
		colorLevel = "\033[34mINF\033[0m"
		colorMessage = fmt.Sprintf("\033[34m%s\033[0m", logMessage.Message)
	case "DEBUG":
		colorLevel = "\033[32mDBG\033[0m"
		colorMessage = fmt.Sprintf("\033[32m%s\033[0m", logMessage.Message)
	case "WARN":
		colorLevel = "\033[33mWRN\033[0m"
		colorMessage = fmt.Sprintf("\033[33m%s\033[0m", logMessage.Message)
	case "ERROR":
		colorLevel = "\033[31mERR\033[0m"
		colorMessage = fmt.Sprintf("\033[31m%s\033[0m", logMessage.Message)
	default:
		colorLevel = "DEBUG"
	}

	formattedMessage := fmt.Sprintf("%s %s \033[1;30m%s:%d\033[0m %s", colorLevel, logMessage.Time, logMessage.File, logMessage.Line, colorMessage)
	fmt.Println(formattedMessage)

	WebhooksProcess(logMessage)
}

func Init(logMessage *LogMessage) {
	_, file, line := GetLineFromCalledFunction()

	files := strings.Split(file, "/")

	if len(files) >= 2 {
		file = fmt.Sprintf("%s/%s", files[len(files)-2], files[len(files)-1])
	}

	logMessage.File = file
	logMessage.Line = line
}

func GetLineFromCalledFunction() (functionName, fileName string, line int) {
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

func AddWebhook(webhook WebhookClient) error {
	switch webhook.Name {
	case "DISCORD":
	case "SLACK":
	default:
		return errors.New("NOT EXIST CLIENT")
	}
	defaultLogger.Webhooks = append(defaultLogger.Webhooks, webhook)

	return nil
}
