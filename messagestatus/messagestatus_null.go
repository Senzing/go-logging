/*
The MessageStatusNull implementation returns an empty string for a status value.
*/
package messagestatus

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageStatusNull struct{}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Get the "status" value given the message id and it's details.
func (messageStatus *MessageStatusNull) MessageStatus(messageNumber int, details ...interface{}) (string, error) {
	return "", nil
}
