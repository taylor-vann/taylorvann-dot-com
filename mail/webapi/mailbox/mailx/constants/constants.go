// brian taylor vann
// taylorvann dot com

package constants

import (
	"os"
)

const (
	Postconf = "postconf"
	OverrideOption = "-o"
	EditOption = "-e"
	InetInterfaces = "loopback-only"

	Postfix = "postfix"
	Stop = "stop"
	Start = "start"

	May = "may"
	EmptyString = ""
	MyDomain="$mydomain"

	Mailx = "mailx"
	SubjectOption = "-s"
	FromOption = "-r"
	FromStatement = "%s<%s>"
	ReplyToOption = "-S"
	ReplyToStatement = "replyto=%s"
)

var (
	Environment = os.Getenv("STAGE")
	Hostname = os.Getenv("POSTFIX_HOSTNAME")
	Host = os.Getenv("POSTFIX_HOST")
)

var postconfSettings = map[string]string{
	"myhostname": Hostname,
	"myorigin": MyDomain,
	"relayhost": MyDomain,
	"smtp_tls_security_level": May,
	"inet_interfaces": InetInterfaces,
	"mydestination": EmptyString,
}

var SetupPostconfAsSendOnlyStatements = buildPostconfSendOnlyCommand()

func buildPostconfSendOnlyCommand() []string {
	var postconfEdits []string

	if postconfSettings["myhostname"] == "" {
		return postconfEdits
	}
	if postconfSettings["myorigin"] == "" {
		return postconfEdits
	}

	for key, value := range postconfSettings {
		overrideKeyValuePair := key + "=" + value
		postconfEdits = append(
			postconfEdits, 
			OverrideOption,
			overrideKeyValuePair,
		)
	}

	return postconfEdits
}