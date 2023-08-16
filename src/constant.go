package main

import "fmt"

// ErrorCodes represents the error codes.
type ErrorCodes int

const (
	NO_INPUT ErrorCodes = iota
	INVALID
	INVALID_ROMAN_CHARACTER
	INVALID_ROMAN_STRING
	INCORRECT_LINE_TYPE
	NO_IDEA
)

// Roman type for Roman numerals and their corresponding values
type Roman int

const (
	I Roman = 1
	V Roman = 5
	X Roman = 10
	L Roman = 50
	C Roman = 100
	D Roman = 500
	M Roman = 1000
)

// ErrorMessage maps error codes to error messages.
type ErrorMessage struct{}

func NewErrorMessage() *ErrorMessage {
	return &ErrorMessage{}
}

// PrintMessage prints the message for the particular error code.
func (e *ErrorMessage) PrintMessage(error ErrorCodes) {
	message := getMessage(error)

	if message != "" {
		fmt.Println(message)
	}
}

func getMessage(error ErrorCodes) string {
	var message string

	switch error {
	case NO_INPUT:
		message = "No input was specified! Program exited"
	case INVALID:
		message = "Input format is wrong! Input discarded"
	case INVALID_ROMAN_CHARACTER:
		message = "Illegal character specified in Roman number! Input discarded"
	case INVALID_ROMAN_STRING:
		message = "Wrong Roman number, violated Roman number format"
	case INCORRECT_LINE_TYPE:
		message = "Exception caused during processing due to incorrect line type supplied"
	case NO_IDEA:
		message = "I have no idea what you are talking about"
	default:
	}

	return message
}