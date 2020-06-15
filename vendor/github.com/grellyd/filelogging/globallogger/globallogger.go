package globallogger

import (
	"github.com/grellyd/filelogging/level"
	"github.com/grellyd/filelogging/logger"
	"github.com/grellyd/filelogging/state"
	"fmt"
)

/*

	GlobalLogger is a wrapper around a filelogging/logger.
	It enables a package to import the globallogger and not need to maintain a reference to it locally nor pass it to objects it calls into.
	
*/

// private global variable for the singleton pattern
var globalLogger *logger.Logger

// NewGlobalLogger creates a new global single instance of a logger.
func NewGlobalLogger(loggerName string, state state.State) (err error) {
	if globalLogger != nil {
		return fmt.Errorf("logger already exists")
	}
	globalLogger, err = logger.NewLogger(loggerName, state)
	if err != nil {
		return fmt.Errorf("unable to create a globallogger: %s", err)
	}
	return nil
}

// Debug Level log
func Debug(data string) {
	if globalLogger == nil {
		fmt.Println("LOGGING ERROR: Global logger uninitialised!")
		fmt.Println(data)
		return
	}
	globalLogger.Log(level.DEBUG, data)
}

// Info Level log
func Info(data string) {
	if globalLogger == nil {
		fmt.Println("LOGGING ERROR: Global logger uninitialised!")
		fmt.Println(data)
		return
	}
	globalLogger.Log(level.INFO, data)
}

// Warning Level log
func Warning(data string) {
	if globalLogger == nil {
		fmt.Println("LOGGING ERROR: Global logger uninitialised!")
		fmt.Println(data)
		return
	}
	globalLogger.Log(level.WARNING, data)
}

// Error Level log
func Error(data string) {
	if globalLogger == nil {
		fmt.Println("LOGGING ERROR: Global logger uninitialised!")
		fmt.Println(data)
		return
	}
	globalLogger.Log(level.ERROR, data)
}

// Fatal Level log
func Fatal(data string) {
	if globalLogger == nil {
		fmt.Println("LOGGING ERROR: Global logger uninitialised!")
		fmt.Println(data)
		return
	}
	globalLogger.Log(level.FATAL, data)
}
