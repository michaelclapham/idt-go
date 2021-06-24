package main

import "testing"

func TestFilenameWithoutExtension(t *testing.T) {
	testFilename := "joe.idml"
	actualResult := FilenameWithoutExtension(testFilename)
	expectedResult := "joe"
	if actualResult != expectedResult {
		t.Errorf("Expected %s received %s", expectedResult, actualResult)
	}
}

func TestFilenameWithoutExtension_MultipleDots(t *testing.T) {
	testFilename := "tricky.filename.idml"
	actualResult := FilenameWithoutExtension(testFilename)
	expectedResult := "tricky.filename"
	if actualResult != expectedResult {
		t.Errorf("Expected %s received %s", expectedResult, actualResult)
	}
}

func TestFilenameWithoutExtension_WithPrecedingPath(t *testing.T) {
	testFilename := "/path/to/filename.idml"
	actualResult := FilenameWithoutExtension(testFilename)
	expectedResult := "filename"
	if actualResult != expectedResult {
		t.Errorf("Expected %s received %s", expectedResult, actualResult)
	}
}
