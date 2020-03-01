package mutations

import (
	"fmt"
	"net/http"
	"webapi/sessions"

	"webapi/hooks/constants"
	"webapi/hooks/sessions/errors"
	"webapi/interfaces/jwtx"
)

// RemoveSession - mutate session whitelist
func RemoveSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("remove session")

	sessionTokenHeader := r.Header.Get(constants.SessionTokenHeader)
	token, _ := jwtx.RetrieveTokenDetailsFromString(&sessionTokenHeader)

	result, errResponseBody := sessions.Remove(
		&sessions.RemoveParams{
			Signature: token.Signature,
		},
	)

	if errResponseBody != nil {
		errAsStr := errResponseBody.Error()
		errors.BadRequest(w, &errors.Response{
			Session: &InvalidHeadersProvided,
			Default: &errAsStr,
		})
		return
	}

	if result == true {
		w.WriteHeader(http.StatusOK)
		return
	}

	errors.BadRequest(w, &errors.Response{
		Session: &InvalidHeadersProvided,
	})
}
