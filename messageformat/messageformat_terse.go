/*
Package message ...
*/
package messageformat

import (
	// "errors"
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageFormatTerse struct {
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// func (messageFormat *MessageFormatTerse) Error(id string, status string, text string, details ...interface{}) error {
// 	message, err := messageFormat.Message(id, status, text, details...)
// 	if err != nil {
// 		return err
// 	}
// 	return errors.New(message)
// }

func (messageFormat *MessageFormatTerse) Message(id string, status string, text string, details ...interface{}) (string, error) {
	var err error = nil

	result := ""
	if len(id) > 0 {
		result = result + fmt.Sprintf("%s: ", id)
	}
	if len(status) > 0 {
		result = result + fmt.Sprintf("(%s) ", status)
	}
	if len(text) > 0 {
		result = result + fmt.Sprintf("%s ", text)
	}
	if len(details) > 0 {
		detailMap := make(map[int]interface{})
		for index, value := range details {
			detailMap[index+1] = value
		}
		result = result + fmt.Sprintf("%v ", detailMap)
	}

	return result, err
}
