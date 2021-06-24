package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

type TranslationEntry struct {
	SourceText string `json:"sourceText"`
	Text       string `json:"text"`
	StoryId    string `json:"storyId"`
	Note       string `json:"note,omitempty"`
	Type       string `json:"type,omitempty"`
}

func FilenameWithoutExtension(filename string) string {
	re := regexp.MustCompile(`(.*)\.[^\.]*$`)
	matches := re.FindAllStringSubmatch(path.Base(filename), -1)
	return matches[0][1]
}

func GetTranslationEntriesForStory(filename string) []TranslationEntry {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	xmlBytes, _ := ioutil.ReadAll(xmlFile)

	var storyXML StoryXML
	err = xml.Unmarshal(xmlBytes, &storyXML)

	if err != nil {
		fmt.Println("XML Parsing error")
		fmt.Println(err)
	}

	storyId := storyXML.Story.Self
	// fmt.Println("Story id", storyId)

	var translationEntries []TranslationEntry

	for _, psr := range storyXML.Story.ParagraphStyleRange {
		for _, csr := range psr.CharacterStyleRange {
			translationEntries = append(translationEntries, TranslationEntry{
				SourceText: csr.Content,
				Text:       csr.Content,
				StoryId:    storyId,
			})
		}
	}
	return translationEntries
}
