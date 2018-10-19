package logger

// name is "logger" instead of "log" so that users of this package can do:
// var log = logger.New("componentName")
// log.Info("hello")

import (
	"log"
)

// https://dave.cheney.net/2015/11/05/lets-talk-about-logging
// Food for thought: https://labs.ig.com/logging-level-wrong-abstraction

type Logger struct {
	middle string
}

// [DEBUG] componentName: your message
func (l *Logger) Debug(msg string) {
	log.Println("[DEBUG" + l.middle + msg)
}

// [INFO] componentName: your message
func (l *Logger) Info(msg string) {
	log.Println("[INFO" + l.middle + msg)
}

// [ERROR] componentName: your message
func (l *Logger) Error(msg string) {
	log.Println("[ERROR" + l.middle + msg)
}

func New(componentName string) *Logger {
	return &Logger{
		middle: "] " + componentName + ": ",
	}
}
