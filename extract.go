package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

func Extract(filename string) {
	println("Extracting strings from ", filename)

	withoutExt := FilenameWithoutExtension(filename)
	tempFolder := path.Join("temp", withoutExt)
	translateJSONFolder := path.Join("translate_json", withoutExt)

	start := time.Now()

	_, err := Unzip(filename, tempFolder)
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Unzipping", filename, "took", elapsed)

	storyDir := path.Join(tempFolder, "Stories")
	storyFiles, err := os.ReadDir(storyDir)
	if err != nil {
		log.Fatal(err)
	}

	start = time.Now()
	var allEntries []TranslationEntry
	for _, storyFile := range storyFiles {
		entries := GetTranslationEntriesForStory(path.Join(storyDir, storyFile.Name()))
		allEntries = append(allEntries, entries...)
	}
	jsonBytes, _ := json.MarshalIndent(allEntries, "", "   ")
	ioutil.WriteFile(path.Join(translateJSONFolder, "entries.json"), jsonBytes, 0644)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("Parsing XML stories for ", filename, "took", elapsed)
}
