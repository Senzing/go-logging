/*
The MessageIdDefault implementation returns a message id based on Sprintf("%v").
*/
package messageid

import "fmt"

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageIdDefault struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (messageId *MessageIdDefault) MessageId(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	result := fmt.Sprintf("%v", messageNumber)
	return result, err
}
