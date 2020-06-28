package verify

import (
	"weblog/store/users/hooks/errors"
	
	"github.com/taylor-vann/weblog/toolbox/golang/verifyx"
)

func isRequestBodyValid(
	w http.ResponseWriter,
	requestBody *requests.Body,
) bool {
	if requestBody != nil && requestBody.Params != nil {
		return true
	}
	errors.BadRequest(w, &responses.Errors{
		RequestBody: &errors.BadRequestFail,
	})
	return false
}

func isInfraSessionValid(
	w http.ResponseWriter,
	environment string,
	sessionToken string,
) bool {
	isValid, errValidate := verifyx.ValidateInfraSession(
		environment,
		sessionToken,
	)
	if isValid {
		return true
	}
	if errValidate != nil {
		errors.DefaultResponse(w, errValidate)
		return false
	}
	
	errors.CustomResponse(w, errors.InvalidInfraSession)
	return false
}