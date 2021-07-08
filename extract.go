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
	tempFolder := path.Join("./temp", withoutExt)
	translateJSONFolder := path.Join("./translate_json", withoutExt)
	_ = os.MkdirAll(translateJSONFolder, os.ModePerm)

	start := time.Now()

	_, err := Unzip(filename, tempFolder)
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Unzipping", filename, "took", elapsed)

	storyDir := path.Join(tempFolder, "Stories")

	start = time.Now()
	storyIdsPerSpread := GetStoryIdsPerSpread(tempFolder)

	for pageIndex, storyIds := range storyIdsPerSpread {
		var entriesForPage []TranslationEntry
		for _, storyId := range storyIds {
			storyFile := path.Join(storyDir, "Story_"+storyId+".xml")
			entries := GetTranslationEntriesForStory(storyFile)
			entriesForPage = append(entriesForPage, entries...)
		}

		jsonBytes, _ := json.MarshalIndent(entriesForPage, "", "   ")
		pageFileName := fmt.Sprintf("%d.json", pageIndex+1)
		ioutil.WriteFile(path.Join(translateJSONFolder, pageFileName), jsonBytes, 0644)
	}
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("Parsing XML stories for ", filename, "took", elapsed)
}
