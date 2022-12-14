/*
The MessageFormatSenzing implementation returns a message in the JSON format.
*/
package messageformat

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageFormatSenzing type is for creating formatted messages in JSON.
type MessageFormatSenzing struct{}

// Fields in the formatted message.
// Order is important.
// It should be date, time, level, id, status, text, duration, location, errors, details.
type messageFormatSenzing struct {
	Date     string      `json:"date,omitempty"`     // Date of message in UTC.
	Time     string      `json:"time,omitempty"`     // Time of message in UTC.
	Level    string      `json:"level,omitempty"`    // Level:  TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC.
	Id       string      `json:"id,omitempty"`       // Message identifier.
	Text     interface{} `json:"text,omitempty"`     // Message text.
	Status   string      `json:"status,omitempty"`   // Status information.
	Duration int64       `json:"duration,omitempty"` // Duration in nanoseconds
	Location string      `json:"location,omitempty"` // Location in the code issuing message.
	Errors   interface{} `json:"errors,omitempty"`   // List of errors.
	Details  interface{} `json:"details,omitempty"`  // All instances passed into the message.
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The Message method creates a JSON formatted message.
func (messageFormat *MessageFormatSenzing) Message(date string, time string, level string, location string, id string, status string, text string, duration int64, errors interface{}, details interface{}) (string, error) {
	messageBuilder := &messageFormatSenzing{}

	if len(date) > 0 {
		messageBuilder.Date = date
	}

	if len(time) > 0 {
		messageBuilder.Time = time
	}

	if len(level) > 0 {
		messageBuilder.Level = level
	}

	if len(location) > 0 {
		messageBuilder.Location = location
	}

	if len(id) > 0 {
		messageBuilder.Id = id
	}

	if len(status) > 0 {
		messageBuilder.Status = status
	}

	if len(text) > 0 {
		if isJson(text) {
			messageBuilder.Text = jsonAsInterface(text)
		} else {
			messageBuilder.Text = text
		}
	}

	messageBuilder.Duration = duration

	if errors != nil {
		if !reflect.ValueOf(errors).IsNil() {
			messageBuilder.Errors = errors
		}
	}

	if details != nil {
		if !reflect.ValueOf(details).IsNil() {
			messageBuilder.Details = details
		}
	}

	// Convert to JSON.

	// Would love to do it this way, but HTML escaping happens.
	// Reported in https://github.com/golang/go/issues/56630
	// result, _ := json.Marshal(messageBuilder)
	// return string(result), err

	// Work-around.

	var resultBytes bytes.Buffer
	enc := json.NewEncoder(&resultBytes)
	enc.SetEscapeHTML(false)
	err := enc.Encode(messageBuilder)
	result := strings.TrimSpace(resultBytes.String())

	return result, err
}
