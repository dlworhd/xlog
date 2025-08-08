package xlog

import (
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
	formattedMessage := fmt.Sprintf("[%s - %-5s] {\"Message\": \"%s\"} \"%s:%d\"", l.Time, level, message, l.File, l.Line)
	fmt.Println(formattedMessage)
}

func (l *Logger) init() {
	_, file, line := getInfo()

	files := strings.Split(file, "/")

	if len(files) >= 2 {
		file = fmt.Sprintf("%s/%s", files[len(files)-2], files[len(files)-1])
	}

	l.File = file
	l.Line = line
	l.Time = time.Now().Format("2006-01-02 15:04:05")
}

func getInfo() (funcname, filename string, line int) {
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
