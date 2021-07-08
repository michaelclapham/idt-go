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

func GetSpreadFiles(tempFolder string) []string {
	xmlFile, err := os.Open(path.Join(tempFolder, "designmap.xml"))
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	xmlBytes, _ := ioutil.ReadAll(xmlFile)

	var designMapXML DesignMapXML
	err = xml.Unmarshal(xmlBytes, &designMapXML)

	if err != nil {
		fmt.Println(err)
	}

	var spreadFileList []string

	for _, v := range designMapXML.Spread {
		spreadFileList = append(spreadFileList, v.Src)
	}
	return spreadFileList
}

func GetStoryIdsPerSpread(tempFolder string) [][]string {
	spreadFiles := GetSpreadFiles(tempFolder)

	storyIdsBySpread := make([][]string, len(spreadFiles))

	for i, spreadFile := range spreadFiles {
		xmlFile, err := os.Open(path.Join(tempFolder, spreadFile))
		if err != nil {
			fmt.Println("Error opening file", err)
			storyIdsBySpread[i] = make([]string, 0)
			break
		}

		xmlBytes, _ := ioutil.ReadAll(xmlFile)

		var spreadXML SpreadXML
		err = xml.Unmarshal(xmlBytes, &spreadXML)
		if err != nil {
			fmt.Println("Error parsing spread file XML for ", spreadFile, err)
		}

		storyIdsBySpread[i] = make([]string, 0)
		for _, textFrame := range spreadXML.Spread.TextFrame {
			if len(textFrame.ParentStory) > 0 {
				storyIdsBySpread[i] = append(storyIdsBySpread[i], textFrame.ParentStory)
			} else {
				fmt.Println("Empty story id in spread file", spreadFile)
			}

		}

		xmlFile.Close()
	}

	return storyIdsBySpread
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
