package mailx

import (
	"testing"
)

var testEmail = EmailParams{
	RecipientAddress:	"brian.t.vann@gmail.com",
	Subject: "taylorvann.com unit test!",
	Body: "Hey brian, it's brian!\n\nThis is a unit test. You can ignore it :)\n\nBest\nBrian",
	ReplyAddress: "unit_tests@taylorvann.com",
	ReplyName: "Unit Tests",
}

func TestStopPostfixService(t *testing.T) {
	output, err := StopPostfixService()
	if err != nil {
		t.Error(err.Error())
	}
	if output != "" {
		t.Error("we expect no output from mail")
	}
}

func TestSetupSendOnlySettings(t *testing.T) {
	output, err := SetupSendOnlySettings()
	if err != nil {
		t.Error(err.Error())
	}
	if output == "" {
		t.Error("we expect a ton of output from postconf")
	}
}

func TestStartPostfixService(t *testing.T) {
	output, err := StartPostfixService()

	if err != nil {
		t.Error(err.Error())
	}
	if output != "" {
		t.Error("we expect no output from mail")
	}
}

func TestSendEmail(t *testing.T) {
	output, err := SendEmail(&testEmail)

	if err != nil {
		t.Error(err.Error())
	}
	if output != "" {
		t.Error("we expect no output from mail")
	}
}