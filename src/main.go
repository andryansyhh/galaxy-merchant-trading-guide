package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to GalaxyMerchant!")

	// Initialize a new paragraph
	paragraph := NewParagraph()

	// Read the input from console, validate, and process
	output := paragraph.Read()

	for _, item := range output {
		fmt.Println(item)
	}
}