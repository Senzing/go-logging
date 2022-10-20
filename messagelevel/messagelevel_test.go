package messagelevel

import (
	"fmt"
	"testing"
)

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func printResult(test *testing.T, title string, result interface{}) {
	if 1 == 1 {
		test.Logf("%s: %v", title, fmt.Sprintf("%v", result))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, testObject MessageLevelInterface, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelSenzingApi - names begin with "Test"
// ----------------------------------------------------------------------------

// -- BuildError --------------------------------------------------------------

func TestSenzingApiBuildError(test *testing.T) {
	testObject := &MessageLevelSenzingApi{}
	actual, err := testObject.CalculateMessageLevel(1, "This is message text")
	testError(test, testObject, err)
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Test interface functions for MessageLevelInfo - names begin with "Test"
// ----------------------------------------------------------------------------

// -- BuildError --------------------------------------------------------------

func TestInfoBuildError(test *testing.T) {
	testObject := &MessageLevelInfo{}
	actual, err := testObject.CalculateMessageLevel(1, "This is message text")
	testError(test, testObject, err)
	printActual(test, actual)
}
