package queries

import (
	"encoding/json"
	"net/http"

	"webapi/store/users/controller"
	"webapi/store/users/hooks/cache"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"

	"webapi/utils/infraclientx/verifyx"
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

func Read(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Read
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

	var params requests.Validate
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !verifyx.IsGuestSessionValid(params.Environment, sessionCookie) {
		errors.BadRequest(w, &responses.Errors{
			Default: &errors.InvalidGuestSession,
		})
		return
	}

	usersStore, errUserStore := controller.Validate(&params)
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

func ValidateUser(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.Validate
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

	usersStore, errUserStore := controller.Validate(&params)
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

	var params requests.Index
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

	var params requests.Search
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
