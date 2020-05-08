// brian taylor vann
// taylorvann dot com

package hooks

import (
	"encoding/json"
	"net/http"

	"webapi/mailbox/hooks/errors"
	"webapi/mailbox/hooks/requests"
	"webapi/mailbox/hooks/responses"
	"webapi/mailbox/hooks/sends"
)

const CreateSendonlyEmail = "CREATE_SEND-ONLY_EMAIL"

func NoReply(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.NilBodyGiven,
		})
		return
	}

	var body requests.Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		errAsStr := err.Error()
		errors.BadRequest(w, &responses.Errors{
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
		return
	}

	switch body.Action {
	case CreateSendonlyEmail:
		sends.NoReply(w, &body)
	default:
		errors.BadRequest(w, &responses.Errors{
			Mail: &errors.UnrecognizedAction,
		})
	}
}