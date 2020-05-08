// brian taylor vann
// taylorvann dot com

package mailx

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"webapi/mailx/constants"
)

type EmailParams struct{
	Body 		 				 string
	RecipientAddress string
	ReplyAddress		 string
	ReplyName				 string
	Subject	 				 string
}

func createFromStatement(name string, email string) string {
	return fmt.Sprintf(constants.FromStatement, name, email)
}

func createReplyToStatement(email string) string {
	return fmt.Sprintf(constants.ReplyToStatement, email)
}

func checkEmailParams(p *EmailParams) (*EmailParams, error) {
	if p == nil {
		return nil, errors.New("nil parameters given")
	}
	if p.ReplyAddress == "" {
		return nil, errors.New("sender is empty string")
	}
	if p.RecipientAddress == "" {
		return nil, errors.New("recipient is empty string")
	}
	if p.Body == "" {
		return nil, errors.New("body is empty string")
	}
	if p.Subject == "" {
		return nil, errors.New("subject is empty string")
	}
	if p.ReplyName == "" {
		return nil, errors.New("user is empty string")
	}

	return p, nil
}

func SetupSendOnlySettings() (string, error) {
	cmd := exec.Command(
		constants.Postconf,
		constants.SetupPostconfAsSendOnlyStatements...,
	)
	
	output, err := cmd.Output()
	return string(output), err
}

func StopPostfixService() (string, error) {
	cmd := exec.Command(
		constants.Postfix,
		constants.Stop,
	)
	output, err := cmd.Output()
	return string(output), err
}

func StartPostfixService() (string, error) {
	cmd := exec.Command(
		constants.Postfix,
		constants.Start,
	)

	output, err := cmd.Output()
	return string(output), err
}

func Setup() {
	StopPostfixService()
	SetupSendOnlySettings()
	StartPostfixService()
}

// mail -s {subject} -r {name<from_address>} -S replyto={reply_addres} {recipient_address}
func SendEmail(p *EmailParams) (string, error) {
	p, errParams := checkEmailParams(p)
	if errParams != nil {
		return "", errParams
	}
	
	fromStatement := createFromStatement(p.ReplyName, p.ReplyAddress)
	replyToStatement := createReplyToStatement(p.ReplyAddress)

	cmd := exec.Command(
		constants.Mailx,
		constants.SubjectOption,
		p.Subject,
		constants.FromOption,
		fromStatement,
		constants.ReplyToOption,
		replyToStatement,
		p.RecipientAddress,
	)

	bodyAsByteBuffer := bytes.NewBuffer([]byte(p.Body))
	cmd.Stdin = bodyAsByteBuffer

	output, errCmd := cmd.Output()
	return string(output), errCmd
}
