package main

import (
	"regexp"
)

// Type represents the different types of conversation lines.
type Type int

const (
	// ASSIGNED represents an assignment line.
	ASSIGNED Type = iota
	// CREDITS represents a credits line.
	CREDITS
	// QUESTION_HOW_MUCH represents a question asking how much.
	QUESTION_HOW_MUCH
	// QUESTION_HOW_MANY represents a question asking how many credits.
	QUESTION_HOW_MANY
	// NOMATCH represents a line that doesn't match any of the defined patterns.
	NOMATCH
)

// LineFilter represents a filter for a specific line type.
type LineFilter struct {
	Type    Type
	Pattern string
}

// ConversationLine represents the main struct for conversation lines.
type ConversationLine struct {
	lineFilters []LineFilter
}

// NewConversationLine initializes a ConversationLine struct.
func NewConversationLine() *ConversationLine {
	return &ConversationLine{
		lineFilters: []LineFilter{
			{Type: ASSIGNED, Pattern: "^([A-Za-z]+) is ([I|V|X|L|C|D|M])$"},
			{Type: CREDITS, Pattern: "^([A-Za-z]+)([A-Za-z\\s]*) is ([0-9]+) ([c|C]redits)$"},
			{Type: QUESTION_HOW_MUCH, Pattern: "^how much is (([A-Za-z\\s])+)\\?$"},
			{Type: QUESTION_HOW_MANY, Pattern: "^how many [c|C]redits is (([A-Za-z\\s])+)\\?$"},
		},
	}
}

// GetLineType determines the type of a given line.
func (cl *ConversationLine) GetLineType(line string) Type {
	line = regexp.MustCompile(`\s+`).ReplaceAllString(line, " ")
	line = regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(line, "")
	result := NOMATCH

	for _, filter := range cl.lineFilters {
		matched, _ := regexp.MatchString(filter.Pattern, line)
		if matched {
			result = filter.Type
			break
		}
	}

	return result
}

// func main() {
// 	// Example usage
// 	conversationLine := NewConversationLine()

// 	line := "glob is V"
// 	lineType := conversationLine.GetLineType(line)
// 	switch lineType {
// 	case ASSIGNED:
// 		println("Assigned line:", line)
// 	case CREDITS:
// 		println("Credits line:", line)
// 	case QUESTION_HOW_MUCH:
// 		println("How much question:", line)
// 	case QUESTION_HOW_MANY:
// 		println("How many question:", line)
// 	case NOMATCH:
// 		println("No match:", line)
// 	}
// }
