package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/store/users/controller"
	"webapi/store/users/hooks/cache"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"

	"webapi/infraclientx/verifyx"
)

func writeUsersResponse(w http.ResponseWriter, users *controller.SafeUsers) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&responses.Body{
		Users: users,
	})
}

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

func Create(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Create
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !verifyx.IsInfraSessionValid(params.Environment, sessionCookie) {
		errors.BadRequest(w, &responses.Errors{
			Default: &errors.InvalidInfraSession,
		})
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
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Update
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !verifyx.IsInfraSessionValid(params.Environment, sessionCookie) {
		errors.BadRequest(w, &responses.Errors{
			Default: &errors.InvalidInfraSession,
		})
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
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.UpdateEmail
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !verifyx.IsInfraSessionValid(params.Environment, sessionCookie) {
		errors.BadRequest(w, &responses.Errors{
			Default: &errors.InvalidInfraSession,
		})
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
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.UpdatePassword
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !verifyx.IsInfraSessionValid(params.Environment, sessionCookie) {
		errors.BadRequest(w, &responses.Errors{
			Default: &errors.InvalidInfraSession,
		})
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
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Delete
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !verifyx.IsInfraSessionValid(params.Environment, sessionCookie) {
		errors.BadRequest(w, &responses.Errors{
			Default: &errors.InvalidInfraSession,
		})
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
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Undelete
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !verifyx.IsInfraSessionValid(params.Environment, sessionCookie) {
		errors.BadRequest(w, &responses.Errors{
			Default: &errors.InvalidInfraSession,
		})
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
