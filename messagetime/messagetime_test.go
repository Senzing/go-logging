package messagetime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name             string
	messageNumber    int
	messageTimestamp time.Time
	details          []interface{}
	expectedDefault  string
	expectedSenzing  string
}{
	{
		name:             "Test case: #1",
		messageNumber:    1001,
		messageTimestamp: time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC),
		expectedSenzing:  "01:01:01.000000001",
	},
	{
		name:             "Test case: #2",
		messageNumber:    1002,
		messageTimestamp: time.Date(2999, time.December, 31, 0, 0, 0, 0, time.UTC),
		expectedSenzing:  "00:00:00.000000000",
	},

	{
		name:             "Test case: #3",
		messageNumber:    1003,
		messageTimestamp: time.Date(2999, time.December, 31, 0, 0, 0, 999999999, time.UTC),
		expectedSenzing:  "00:00:00.999999999",
	},
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject MessageTimeInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageTimeNull - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageTimeNull(test *testing.T) {
	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			testObject := &MessageTimeNull{}
			actual, err := testObject.MessageTime(testCase.messageNumber, testCase.messageTimestamp, testCase.details...)
			testError(test, testObject, err)
			assert.Equal(test, "", actual, testCase.name)
		})
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageTimeSenzing - names begin with "Test"
// ----------------------------------------------------------------------------

func TestMessageTimeSenzing(test *testing.T) {
	for _, testCase := range testCases {
		if len(testCase.expectedSenzing) > 0 {
			test.Run(testCase.name, func(test *testing.T) {
				testObject := &MessageTimeSenzing{}
				actual, err := testObject.MessageTime(testCase.messageNumber, testCase.messageTimestamp, testCase.details...)
				testError(test, testObject, err)
				assert.Equal(test, testCase.expectedSenzing, actual, testCase.name)
			})
		}
	}
}