package main

import (
	"os"
)

func main() {
	// Use arguments to decide command and which file to apply to
	if len(os.Args) < 3 {
		println("Usage: idt-go [COMMAND] [FILENAME]")
		println("Example: idt-go extract test.idml")
		os.Exit(-1)
	}

	if os.Args[1] == "extract" {
		Extract(os.Args[2])
	} else if os.Args[1] == "translate" {
		Translate(os.Args[2])
	} else {
		println("Unknown command", os.Args[1])
	}
}
