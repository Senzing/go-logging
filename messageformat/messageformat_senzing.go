/*
The MessageFormatJson implementation returns a message in the JSON format.
*/
package messageformat

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageFormatJson type is for creating formatted messages in JSON.
type MessageFormatSenzing struct{}

// Fields in the formatted message.
// Order is important.
// It should be id, status, text, errors, details.
type messageFormatSenzing struct {
	Date     string        `json:"date,omitempty"`
	Time     string        `json:"time,omitempty"`
	Level    string        `json:"level,omitempty"`
	Location string        `json:"location,omitempty"`
	Id       string        `json:"id,omitempty"`
	Status   string        `json:"status,omitempty"`
	Text     interface{}   `json:"text,omitempty"`
	Duration int64         `json:"duration,omitempty"`
	Errors   []interface{} `json:"errors,omitempty"`
	Details  interface{}   `json:"details,omitempty"`
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The Message method creates a JSON formatted message.
func (messageFormat *MessageFormatSenzing) Message(date string, time string, level string, location string, id string, status string, text string, duration int64, details ...interface{}) (string, error) {
	var err error = nil
	messageBuilder := &messageFormatSenzing{}

	// Set output Id, Status, and Text fields.

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

	// Work with details.

	if len(details) > 0 {
		var errorsList []interface{}
		detailMap := make(map[string]interface{})

		// Process different types of details.

		for index, value := range details {
			switch typedValue := value.(type) {
			case nil:
				detailMap[strconv.Itoa(index+1)] = "<nil>"

			case error:
				errorMessage := typedValue.Error()
				var priorError interface{}
				if isJson(errorMessage) {
					priorError = jsonAsInterface(errorMessage)
				} else {
					priorError = &messageFormatSenzing{
						Text: errorMessage,
					}
				}
				errorsList = append(errorsList, priorError)

			case map[string]string:
				for mapIndex, mapValue := range typedValue {
					mapValueAsString := stringify(mapValue)
					if isJson(mapValueAsString) {
						detailMap[mapIndex] = jsonAsInterface(mapValueAsString)
					} else {
						detailMap[mapIndex] = mapValueAsString
					}
				}

			default:
				valueAsString := stringify(typedValue)
				if isJson(valueAsString) {
					detailMap[strconv.Itoa(index+1)] = jsonAsInterface(valueAsString)
				} else {
					detailMap[strconv.Itoa(index+1)] = valueAsString
				}
			}
		}

		// Set output Errors and Details fields.

		messageBuilder.Errors = errorsList
		messageBuilder.Details = detailMap
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
	err = enc.Encode(messageBuilder)
	result := strings.TrimSpace(resultBytes.String())

	return result, err
}
