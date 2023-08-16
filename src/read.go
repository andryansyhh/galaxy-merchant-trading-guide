package main

import (
	"fmt"
)

// Paragraph struct
type Paragraph struct {
	constantAssignments map[string]string
	computedLiterals    map[string]string
	output              []string
}

// NewParagraph creates a new instance of Paragraph
func NewParagraph() *Paragraph {
	return &Paragraph{
		constantAssignments: make(map[string]string),
		computedLiterals:    make(map[string]string),
		output:              []string{},
	}
}

// Read method reads the paragraph from input and processes it
func (p *Paragraph) Read() []string {
	var line string
	count := 0

	for {
		_, err := fmt.Scanln(&line)
		if err != nil || line == "" {
			break
		}

		count++
	}

	switch count {
	case 0:
		p.output = append(p.output, "No input provided")
	default:
	}

	return p.output
}
