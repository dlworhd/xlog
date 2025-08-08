package xlog

import (
	"fmt"
	"time"
)

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
}

func Default(min Level) {
	defaultLogger.minLevel = min
}

func Info(message string) {
	defaultLogger.Log("INFO", message)
}

func Warn(message string) {
	defaultLogger.Log("WARN", message)
}

func Error(message string) {
	defaultLogger.Log("ERROR", message)
}

func Debug(message string) {
	defaultLogger.Log("DEBUG", message)
}

func (l *Logger) Log(level Level, message string) {
	l.init()
	l.log(level, message)
}

func (l *Logger) log(level Level, message string) {
	if levelPriority[level] < levelPriority[l.minLevel] {
		return
	}
	formattedMessage := fmt.Sprintf("[%s - %-5s] {\"Message\": \"%s\"}", l.Time, level, message)
	fmt.Println(formattedMessage)
}

func (l *Logger) init() {
	l.Time = time.Now().Format("2006-01-02 15:04:05")
}
