package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/store/users/controller"
	"webapi/store/users/hooks/cache"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"

	"webapi/store/clientx"
	clientxRequests "webapi/store/clientx/fetch/requests"


	"github.com/taylor-vann/tvgtb/jwtx"
)

const SessionCookieHeader = "briantaylorvann.com_session"

func writeUsersResponse(w http.ResponseWriter, users *controller.SafeUsers) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&responses.Body{
		Users: users,
	})
}

func checkInfraSession(sessionCookie *http.Cookie) (bool, error) {
	details, errDetails := jwtx.RetrieveTokenDetailsFromString(sessionCookie.Value)
	if errDetails != nil {
		return false, errDetails
	}
	if details == nil {
		return false, nil
	}
	if details.Payload.Sub == "infra" && details.Payload.Aud == "public" {
		return true, nil
	}

	return false, nil
}

// just pass the environment
func validateInfraSession(sessionCookie *http.Cookie, environment string) (bool, error) {
	sessionIsValid, errSessionIsValid := checkInfraSession(sessionCookie)
	if !sessionIsValid || errSessionIsValid != nil {
		return sessionIsValid, errSessionIsValid
	}

	// remote check, TODO
	session, errSession := clientx.ValidateSession(clientxRequests.ValidateSession{
		Environment: environment,
		Token: sessionCookie.Value,
	})
	if errSession != nil {
		return false, errSession
	}
	if session != "" {
		return true, nil
	}

	return false, nil
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

func Create(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Create
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Update
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.UpdateEmail
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.UpdatePassword
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Delete
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Undelete
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
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
