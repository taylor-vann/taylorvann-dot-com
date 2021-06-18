package details

import (
	"testing"
)

const (
	exampleDetailsPath = "/usr/local/config/config.example.json"
)

func TestReadFile(t *testing.T) {
	currentFile, errCurrentFile := readFile(exampleDetailsPath)
	if currentFile == nil {
		t.Fail()
		t.Logf("There should be an example init file available.")
	}

	if errCurrentFile != nil {
		t.Fail()
		t.Logf(errCurrentFile.Error())
	}
}

func TestParseDetails(t *testing.T) {
	exampleFile, errExampleFile := readFile(exampleDetailsPath)
	exampleDetails, errExampleDetails := parseDetails(exampleFile, errExampleFile)

	if exampleDetails == nil {
		t.Fail()
		t.Logf("There should be details that can be parsed")
	}

	if errExampleDetails != nil {
		t.Fail()
		t.Logf(errExampleDetails.Error())
	}
}

func TestReadDetailsFromFile(t *testing.T) {
	exampleDetails, errExampleDetails := ReadDetailsFromFile(exampleDetailsPath)
	if errExampleDetails != nil {
		t.Fail()
		t.Logf(errExampleDetails.Error())
	}

	if exampleDetails == nil {
		t.Fail()
		t.Logf("Application details should be created")
	}
}