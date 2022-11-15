/*
The messageerrors package produces a date string.
*/
package messageerrors

import (
	"encoding/json"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageErrorsInterface type defines methods for aggregating errors from details.
type MessageErrorsInterface interface {
	MessageErrors(messageNumber int, details ...interface{}) (interface{}, error)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func isJson(unknownString string) bool {
	unknownStringUnescaped, err := strconv.Unquote(unknownString)
	if err != nil {
		unknownStringUnescaped = unknownString
	}
	var jsonString json.RawMessage
	return json.Unmarshal([]byte(unknownStringUnescaped), &jsonString) == nil
}

func jsonAsInterface(unknownString string) interface{} {
	unknownStringUnescaped, err := strconv.Unquote(unknownString)
	if err != nil {
		unknownStringUnescaped = unknownString
	}
	var jsonString json.RawMessage
	json.Unmarshal([]byte(unknownStringUnescaped), &jsonString)
	return jsonString
}
