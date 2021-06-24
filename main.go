package main

import (
	"log"
	"os"
	"path"
	"regexp"
)

func main() {
	// Use arguments to decide command and which file to apply to
	if len(os.Args) < 3 {
		println("Usage: idt-go [COMMAND] [FILENAME]")
		println("Example: idt-go extract test.idml")
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

func FilenameWithoutExtension(filename string) string {
	re := regexp.MustCompile(`(.*)\.[^\.]*$`)
	matches := re.FindAllStringSubmatch(path.Base(filename), -1)
	return matches[0][1]
}

func extract(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		println("Looking in input folder for ", filename)
		filename = path.Join("./input", filename)
	}

	println("Extracting strings from ", filename)
	outputFolder := path.Join("output", FilenameWithoutExtension(filename))
	_, err := Unzip(filename, outputFolder)
	if err != nil {
		log.Fatal(err)
	}
}

func translate(filename string) {
	println("Translating file ", filename)
}
