package main

import (
	"path"
	"regexp"
)

func FilenameWithoutExtension(filename string) string {
	re := regexp.MustCompile(`(.*)\.[^\.]*$`)
	matches := re.FindAllStringSubmatch(path.Base(filename), -1)
	return matches[0][1]
}
