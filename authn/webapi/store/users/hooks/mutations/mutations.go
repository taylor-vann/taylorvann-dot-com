package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/store/users/controller"
	"webapi/store/users/hooks/cache"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"

	"toolbox/clientx"
	fetchRequests "toolbox/clientx/fetch/requests"

	"toolbox/jwtx"
)

const SessionCookieHeader = "briantaylorvann.com_session"

func writeUsersResponse(w http.ResponseWriter, users *controller.SafeUsers) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&responses.Body{
		Users: users,
	})
}

func dropRequestNotValidBody(w http.ResponseWriter, requestBody *requests.Body) bool {
	if requestBody != nil && requestBody.Params != nil {
		return false
	}
	errors.BadRequest(w, &responses.Errors{
		RequestBody: &errors.BadRequestFail,
	})
	return true
}

func checkInfraSession(sessionToken string) (bool, error) {
	isValid := jwtx.ValidateGenericToken(&jwtx.ValidateGenericTokenParams{
		Token: sessionToken,
		Issuer: "briantaylorvann.com",
	})
	if !isValid {
		return false, nil
	}

	details, errDetails := jwtx.RetrieveTokenDetailsFromString(sessionToken)
	if errDetails != nil {
		return false, errDetails
	}

	if details.Payload.Sub == "infra" {
		return true, nil
	}

	return false, nil
}

func validateSessionRemotely(environment string, sessionToken string) (bool, error) {
	sessionStr, errSessionStr := clientx.ValidateSession(
		fetchRequests.ValidateSession{
			Environment: environment,
			Token: sessionToken,
		},
	)
	if errSessionStr != nil {
		return false, errSessionStr
	}
	if sessionStr != "" {
		return true, nil
	}

	return false, nil
}

func validateInfraSession(environment string, sessionToken string) (bool, error) {
	infraSessionExists, errInfraSessionExists := validateSessionRemotely(
		environment,
		sessionToken,
	)
	if errInfraSessionExists != nil {
		return false, errInfraSessionExists
	}
	if !infraSessionExists {
		return false, nil
	}

	return checkInfraSession(sessionToken)
}

func Create(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.Create
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	sessionIsValid, errSessionIsValid := validateInfraSession(
		params.Environment,
		sessionCookie.Value,
	)
	if errSessionIsValid != nil{
		errors.DefaultResponse(w, errSessionIsValid)
		return
	}
	if !sessionIsValid {
		errors.CustomResponse(w, errors.InvalidInfraSession)
		return
	}

	users, errCreateUser := controller.Create(&params)
	if errCreateUser != nil {
		errors.DefaultResponse(w, errCreateUser)
		return
	}

	if users != nil {
		cache.UpdateReadEntry(params.Environment, &users)
		writeUsersResponse(w, &users)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToCreateUser,
	})
}

func Update(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.Update
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	sessionIsValid, errSessionIsValid := validateInfraSession(
		params.Environment,
		sessionCookie.Value,
	)
	if errSessionIsValid != nil{
		errors.DefaultResponse(w, errSessionIsValid)
		return
	}
	if !sessionIsValid {
		errors.CustomResponse(w, errors.InvalidInfraSession)
		return
	}

	users, errUpdateUsers := controller.Update(&params)
	if errUpdateUsers != nil {
		errors.DefaultResponse(w, errUpdateUsers)
		return
	}

	if users != nil {
		cache.UpdateReadEntry(params.Environment, &users)
		writeUsersResponse(w, &users)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToUpdateUser,
	})
}

func UpdateEmail(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.UpdateEmail
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	sessionIsValid, errSessionIsValid := validateInfraSession(
		params.Environment,
		sessionCookie.Value,
	)
	if errSessionIsValid != nil{
		errors.DefaultResponse(w, errSessionIsValid)
		return
	}
	if !sessionIsValid {
		errors.CustomResponse(w, errors.InvalidInfraSession)
		return
	}

	users, errUpdateEmailUser := controller.UpdateEmail(&params)
	if errUpdateEmailUser != nil {
		errors.DefaultResponse(w, errUpdateEmailUser)
		return
	}

	if users != nil {
		cache.UpdateReadEntry(params.Environment, &users)
		writeUsersResponse(w, &users)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToUpdateEmailUser,
	})
}

func UpdatePassword(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.UpdatePassword
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	sessionIsValid, errSessionIsValid := validateInfraSession(
		params.Environment,
		sessionCookie.Value,
	)
	if errSessionIsValid != nil{
		errors.DefaultResponse(w, errSessionIsValid)
		return
	}
	if !sessionIsValid {
		errors.CustomResponse(w, errors.InvalidInfraSession)
		return
	}

	users, errUpdatePasswordUser := controller.UpdatePassword(&params)
	if errUpdatePasswordUser != nil {
		errors.DefaultResponse(w, errUpdatePasswordUser)
		return
	}

	if users != nil {
		cache.UpdateReadEntry(params.Environment, &users)
		writeUsersResponse(w, &users)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToUpdatePasswordUser,
	})
}

func Delete(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.Delete
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	sessionIsValid, errSessionIsValid := validateInfraSession(
		params.Environment,
		sessionCookie.Value,
	)
	if errSessionIsValid != nil{
		errors.DefaultResponse(w, errSessionIsValid)
		return
	}
	if !sessionIsValid {
		errors.CustomResponse(w, errors.InvalidInfraSession)
		return
	}

	users, errDeleteUsers := controller.Delete(&params)
	if errDeleteUsers != nil {
		errors.DefaultResponse(w, errDeleteUsers)
		return
	}

	if users != nil {
		cache.UpdateReadEntry(params.Environment, &users)
		writeUsersResponse(w, &users)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToDeleteUser,
	})
}

func Undelete(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.Undelete
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	sessionIsValid, errSessionIsValid := validateInfraSession(
		params.Environment,
		sessionCookie.Value,
	)
	if errSessionIsValid != nil{
		errors.DefaultResponse(w, errSessionIsValid)
		return
	}
	if !sessionIsValid {
		errors.CustomResponse(w, errors.InvalidInfraSession)
		return
	}

	users, errUndeleteUsers := controller.Undelete(&params)
	if errUndeleteUsers != nil {
		errors.DefaultResponse(w, errUndeleteUsers)
		return
	}

	if users != nil {
		cache.UpdateReadEntry(params.Environment, &users)
		writeUsersResponse(w, &users)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToUndeleteUser,
	})
}
