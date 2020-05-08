// brian taylor vann
// taylorvann dot com

package sends

import (
	"encoding/json"
	"net/http"
	
	"webapi/mailbox/hooks/errors"
	"webapi/mailbox/hooks/requests"
	"webapi/mailbox/hooks/responses"
	"webapi/mailbox/mailx"
)

var SuccessfullyExecutedMailxCommand = "successfully executed mailx command"

func NoReply(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.NilBodyGiven,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var action requests.NoReply
	errActionMarshal := json.Unmarshal(bytes, &action)
	if errActionMarshal != nil {
		errAsStr := errActionMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
		return
	}

	_, errSendMail := mailx.SendEmail(action.Params)
	if errSendMail == nil {
		json.NewEncoder(w).Encode(&responses.Body{
			Mail: &SuccessfullyExecutedMailxCommand,
		})
		return
	}

	errSendMailAsStr := errActionMarshal.Error()
	errors.BadRequest(w, &responses.Errors{
		Mail: &errors.UnrecognizedParams,
		Default: &errSendMailAsStr,
	})
}
