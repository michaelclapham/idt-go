package main

import (
	"os"
	"path"
)

func main() {
	// Use arguments to decide command and which file to apply to
	if len(os.Args) < 3 {
		println("Usage: idt-go [COMMAND] [FILENAME]")
		println("Example: idt-go extract test.idml")
		os.Exit(-1)
	}

	filename := os.Args[2]

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		println("Looking in input folder for ", filename)
		filename = path.Join("./input", filename)
	}

	if os.Args[1] == "extract" {
		Extract(filename)
	} else if os.Args[1] == "translate" {
		Translate(filename)
	} else {
		println("Unknown command", os.Args[1])
	}
}
