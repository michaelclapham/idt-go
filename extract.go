package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

func Extract(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		println("Looking in input folder for ", filename)
		filename = path.Join("./input", filename)
	}

	println("Extracting strings from ", filename)

	start := time.Now()
	outputFolder := path.Join("output", FilenameWithoutExtension(filename))
	_, err := Unzip(filename, outputFolder)
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Unzipping", filename, "took", elapsed)

	storyDir := path.Join(outputFolder, "Stories")
	storyFiles, err := os.ReadDir(storyDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, storyFile := range storyFiles {
		entries := GetTranslationEntriesForStory(path.Join(storyDir, storyFile.Name()))
		for _, entry := range entries {
			fmt.Println(entry)
		}
	}
}
