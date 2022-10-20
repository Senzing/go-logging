package main

import (
	"log"
	"time"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagelogger"
)

// Values updated via "go install -ldflags" parameters.

// var programName string = "unknown"
// var buildVersion string = "0.0.0"
// var buildIteration string = "0"

const MessageIdFormat = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func complexProcess() string {
	time.Sleep(1000 * time.Second)
	return "slept"
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {

	// Configure the "log" standard library.

	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds | log.LUTC)

	// Configure the logger. If not configured, no functions will print.

	messagelogger.SetLevel(messagelogger.LevelInfo)

	// ------------------------------------------------------------------------
	// The following demonstrates the high-level messagelogger calls for
	// LogMessage, LogMessageUsingMap, and LogMessageFromError.
	// ------------------------------------------------------------------------

	// Log a message.

	// messagelogger.LogMessage(MessageIdFormat, 999, "Test INFO message 1", programName, buildVersion, buildIteration)

	// Log a message using a map.

	// detailsMap := map[string]interface{}{
	// 	"FirstVariable":  "First value",
	// 	"SecondVariable": "Second value",
	// }
	// messagelogger.LogMessageUsingMap(MessageIdFormat, 1000, "Test WARN message 2", detailsMap)

	// Log an error based on a prior error.

	// anError := errors.New("this is a new error")
	// messagelogger.LogMessageFromError(MessageIdFormat, 2000, "Test ERROR message 3", anError, "Variable1", "Variable2")

	// Won't print because of logging level.

	// messagelogger.LogMessageFromErrorUsingMap(MessageIdFormat, 3000, "Test DEBUG message 4", anError, detailsMap)

	// Change logging level and try again. Then restore logging level

	// messagelogger.SetLevel(messagelogger.LevelDebug)
	// messagelogger.LogMessageFromErrorUsingMap(MessageIdFormat, 3000, "Test DEBUG message 5", anError, detailsMap)
	// messagelogger.SetLevel(messagelogger.LevelInfo)

	// ------------------------------------------------------------------------
	// The following demonstrates the low-level logger calls for
	// Trace, Debug, Info, Warn, and Error.
	// ------------------------------------------------------------------------

	log.Println("Test Trace")
	logger.Trace("trace prints")
	logger.Tracef("trace A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Debug")
	logger.Debug("debug prints")
	logger.Debugf("debug A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Info")
	logger.Info("info prints")
	logger.Infof("info A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Warn")
	logger.Warn("warn prints")
	logger.Warnf("warn A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Error")
	logger.Error("error prints")
	logger.Errorf("error A: %s B: %s C: %d", "aaa", "bbb", 35)

	// Avoid long running logging when appropriate.

	if logger.IsDebug() {
		logger.Debugf("%s", complexProcess())
	}

	// Note:  the first Fatal or Panic issued will exit the program.

	log.Println("Test Fatal")
	//	logger.Fatal("fatal prints")
	//	logger.Fatalf("fatal A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Panic")
	//		logger.Fatal("fatal prints")
	//	logger.Fatalf("fatal A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test varadic")
	_, err := time.LoadLocation("bob")
	logger.Info("Should be error: ", err)

	log.Println("End")

}
