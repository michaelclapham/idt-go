package main

import "os"

func main() {
	// Use arguments to decide command and which file to apply to
	if len(os.Args) < 3 {
		println("Usage: idt-rust [COMMAND] [FILENAME]")
		println("Example: idt-rust extract test.idml")
		os.Exit(-1)
	}

	if os.Args[1] == "extract" {
		extract(os.Args[2])
	} else if os.Args[1] == "translate" {
		translate(os.Args[2])
	} else {
		println("Unknown command", os.Args[1])
	}
}

func extract(filename string) {
	println("Extracting strings from ", filename)
}

func translate(filename string) {
	println("Translating file ", filename)
}
