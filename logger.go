package ftp

import (
	"fmt"
	"log"
)

// Logger wraps several logging methods
type Logger interface {
	Print(sessionID string, message interface{})
	Printf(sessionID string, format string, v ...interface{})
	PrintCommand(sessionID string, command string, params string)
	PrintResponse(sessionID string, code int, message string)
}

type stdLogger struct{}

func (logger *stdLogger) Print(sessionID string, message interface{}) {
	log.Printf("%s  %s", sessionID, message)
}

func (logger *stdLogger) Printf(sessionID string, format string, v ...interface{}) {
	logger.Print(sessionID, fmt.Sprintf(format, v...))
}

func (logger *stdLogger) PrintCommand(sessionID string, command string, params string) {
	if command == "PASS" {
		log.Printf("%s > PASS ****", sessionID)
	} else {
		log.Printf("%s > %s %s", sessionID, command, params)
	}
}

func (logger *stdLogger) PrintResponse(sessionID string, code int, message string) {
	log.Printf("%s < %d %s", sessionID, code, message)
}
