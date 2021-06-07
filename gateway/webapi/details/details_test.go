package details

import (
	"fmt"
	"testing"
)

const (
	exampleDetailsPath = "/usr/local/config/details.init.example.json"
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

func TestReadCorrectDetailsFromFile(t *testing.T) {
	expectedCerts := "/usr/local/certs/fullchain.pem"
	expectedPrivKey := "/usr/local/certs/privkey.pem"
	superawesome := "https://superawesome.com"
	expectedAddress := "https://127.0.0.1:5000"

	exampleFile, errExampleFile := readFile(exampleDetailsPath)
	exampleDetails, _ := parseDetails(exampleFile, errExampleFile)
	superawesomeAddress, keyFound := exampleDetails.Routes[superawesome]

	if exampleDetails.CertPaths.Cert != expectedCerts {
		t.Fail()
		t.Logf(fmt.Sprint("Example detail cert path should be:", expectedCerts))
	}

	if exampleDetails.CertPaths.PrivateKey != expectedPrivKey {
		t.Fail()
		t.Logf(fmt.Sprint("Example detail cert path should be:", expectedPrivKey))
	}

	if keyFound != true {
		t.Fail()
		t.Logf(fmt.Sprint("Key should exist:", superawesome))
	}

	if superawesomeAddress != expectedAddress {
		t.Fail()
		t.Logf(fmt.Sprint("Expected address:", superawesomeAddress))
	}
}
