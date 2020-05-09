// brian taylor vann
// taylorvann dot com

package requests

import (
	"webapi/mailbox/mailx"
)

type EmailParams = mailx.EmailParams

type Body struct {
	Action string
	Params interface{}
}