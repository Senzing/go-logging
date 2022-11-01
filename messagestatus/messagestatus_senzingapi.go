/*
The MessageStatusSenzingApi implementation calculates a status value based on message id and Senzing return code.
*/
package messagestatus

import (
	"fmt"
	"sort"
	"strings"

	"github.com/senzing/go-logging/logger"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

/*
The MessageStatusSenzingApi type is for constructing status values by first
looking at the Senzing error code.
If it doesn't exist, use the messageNumber to calculate a status.
*/
type MessageStatusSenzingApi struct {
	IdRanges   map[int]string
	IdStatuses map[int]string
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

/*
Types of Senzing errors.
These are the strings that may be returned from MessageStatus()
*/
const (
	Debug              = logger.LevelDebugName
	Error              = logger.LevelErrorName
	ErrorBadUserInput  = "bad_user_input"
	ErrorRetryable     = "retryable"
	ErrorUnrecoverable = "unrecoverable"
	Fatal              = logger.LevelFatalName
	Info               = logger.LevelInfoName
	Panic              = logger.LevelPanicName
	Trace              = logger.LevelTraceName
	Warn               = logger.LevelWarnName
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// A map of Senzing errors to the corresponding error level.
var senzingApiErrorsMap = map[string]string{
	"0002E":  Info,
	"0019E":  ErrorUnrecoverable,
	"0037E":  ErrorBadUserInput,  // Unknown resolved entity value
	"0063E":  ErrorUnrecoverable, // G2ConfigMgr is not initialized
	"30121E": ErrorBadUserInput,  // JSON parsing Failure
}

// The order of severity/verbosity from most severe to most verbose.
var messagePrecedence = []string{
	Panic,
	Fatal,
	ErrorUnrecoverable,
	ErrorBadUserInput,
	ErrorRetryable,
	Error,
	Warn,
	Info,
	Debug,
	Trace,
}

func (messageStatus *MessageStatusSenzingApi) messageStatusBySenzingError(messageNumber int, details ...interface{}) string {

	// Create a list of Senzing errors by looking at details in reverse order.

	var senzingErrors []string
	if len(details) > 0 {
		for index := len(details) - 1; index >= 0; index-- {
			detail := details[index]
			switch typedDetail := detail.(type) {
			case error:
				errorMessage := typedDetail.Error()
				messageSplits := strings.Split(errorMessage, "|")
				for key, value := range senzingApiErrorsMap {
					if messageSplits[0] == key {
						senzingErrors = append(senzingErrors, value)
					}
				}
			}
		}
	}

	// In the list of Senzing errors, determine the highest priority error.

	if len(senzingErrors) > 0 {
		for _, messagePrecedenceLevel := range messagePrecedence {
			for _, senzingError := range senzingErrors {
				if senzingError == messagePrecedenceLevel {
					return senzingError
				}
			}
		}
	}
	return ""
}

func (messageStatus *MessageStatusSenzingApi) messageStatusByIdRanges(messageNumber int) string {
	// Create a list of sorted keys.

	messageLevelKeys := make([]int, 0, len(messageStatus.IdRanges))
	for key := range messageStatus.IdRanges {
		messageLevelKeys = append(messageLevelKeys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(messageLevelKeys)))

	// Using the sorted message number, find the level.

	for _, messageLevelKey := range messageLevelKeys {
		if messageNumber >= messageLevelKey {
			return messageStatus.IdRanges[messageLevelKey]
		}
	}
	return ""
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageStatus method returns a status based on a message number indexed into senzingApiErrorsMap.
func (messageStatus *MessageStatusSenzingApi) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil

	// --- Status based on Senzing error passed in via details ----------------

	result := messageStatus.messageStatusBySenzingError(messageNumber, details...)
	if len(result) > 0 {
		return result, err
	}

	// --- Status based on messageNumber ----------------------------------------

	if messageStatus.IdStatuses != nil {
		result, ok := messageStatus.IdStatuses[messageNumber]
		if ok {
			return result, err
		}
	}

	// --- Status based on messageNumber range ----------------------------------

	if messageStatus.IdRanges != nil {
		result = messageStatus.messageStatusByIdRanges(messageNumber)
		if len(result) > 0 {
			return result, err
		}
	}

	// --- At this point, failed to find status -------------------------------

	return result, fmt.Errorf("could not determine status for message number %d", messageNumber)
}
