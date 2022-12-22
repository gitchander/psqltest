package main

import (
	"fmt"
	"log"
)

type LogLevel int

const (
	LogLevelInfo    = iota
	LogLevelWarning = iota
	LogLevelError   = iota
)

var levelsNames = map[LogLevel]string{
	LogLevelInfo:    "INF",
	LogLevelWarning: "WAR",
	LogLevelError:   "ERR",
}

func formatLogLevel(level LogLevel) string {
	name, ok := levelsNames[level]
	if ok {
		return name
	}
	return fmt.Sprintf("%s(%d)", "LogLevel", int(level))
}

func logMessage(level LogLevel, message string) {
	log.Printf("%s %s", formatLogLevel(level), message)
}

func logInfo(message string) {
	logMessage(LogLevelInfo, message)
}

func logWarning(message string) {
	logMessage(LogLevelWarning, message)
}

func logError(message string) {
	logMessage(LogLevelError, message)
}

func errLog(err error) {
	if err != nil {
		logError(err.Error())
	}
}

func handleError(err error) bool {
	if err == nil {
		return false
	}
	logError(err.Error())
	return true
}
