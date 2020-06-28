package queries

import (
	"encoding/json"
	"net/http"

	"webapi/store/users/controller"
	"webapi/store/users/hooks/cache"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"
	
	"github.com/taylor-vann/weblog/toolbox/golang/verifyx"
)

func writeUsersResponse(w http.ResponseWriter, users *controller.SafeUsers) {
	w.Header().Set("Content-Type", "application/json")
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

func isGuestSessionValid(
	w http.ResponseWriter,
	environment string,
	sessionToken string,
) bool {
	isValid, errValidate := verifyx.ValidateGuestSession(
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

func Read(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Read
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !isInfraSessionValid(w, params.Environment, sessionCookie.Value) {
		return
	}

	users, errReadUserCache := cache.GetReadEntry(&params)
	if errReadUserCache != nil {
		errors.DefaultResponse(w, errReadUserCache)
		return
	}
	if users != nil {
		writeUsersResponse(w, users)
		return
	}

	usersStore, errUserStore := controller.Read(&params)
	if errUserStore != nil {
		errors.DefaultResponse(w, errUserStore)
		return
	}
	if users != nil {
		cache.UpdateReadEntry(params.Environment, &usersStore)
		writeUsersResponse(w, &usersStore)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToReadUser,
	})
}

func ValidateGuest(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.Validate
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !isGuestSessionValid(w, params.Environment, sessionCookie.Value) {
		return
	}

	usersStore, errUserStore := controller.Validate(&requests.Validate{
		Environment: params.Environment,
		Email: params.Email,
		Password: params.Password,
	})
	if errUserStore != nil {
		errors.DefaultResponse(w, errUserStore)
		return
	}
	if usersStore != nil {
		cache.UpdateReadEntry(params.Environment, &usersStore)
		writeUsersResponse(w, &usersStore)
		return
	}
	
	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToReadUser,
	})
}

func Index(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.Index
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !isInfraSessionValid(w, params.Environment, sessionCookie.Value) {
		return
	}

	users, errIndexUsers := controller.Index(&params)
	if errIndexUsers != nil {
		errors.DefaultResponse(w, errIndexUsers)
		return
	}
	if users != nil {
		writeUsersResponse(w, &users)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToIndexUsers,
	})
}

func Search(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.Search
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !isInfraSessionValid(w, params.Environment, sessionCookie.Value) {
		return
	}

	users, errSearchUsers := controller.Search(&params)
	if errSearchUsers != nil {
		errors.DefaultResponse(w, errSearchUsers)
		return
	}
	if users != nil {
		writeUsersResponse(w, &users)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToSearchUsers,
	})
}

