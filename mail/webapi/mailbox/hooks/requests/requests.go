// brian taylor vann
// taylorvann dot com

package requests

import (
	"webapi/mailbox/mailx"
)

type NoReply struct {
	Action string
	Params *mailx.EmailParams
}

type Body struct {
	Action string
	Params interface{}
}