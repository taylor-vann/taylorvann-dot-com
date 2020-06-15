package queries

import (
	"encoding/json"
	"net/http"

	"github.com/taylor-vann/weblog/toolbox/golang/clientx"
	"github.com/taylor-vann/weblog/toolbox/golang/clientx/fetch"
	fetchRequests "github.com/taylor-vann/weblog/toolbox/golang/clientx/fetch/requests"

	"webapi/store/users/controller"
	"webapi/store/users/hooks/cache"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

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

func Read(
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

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Read
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
	if dropRequestNotValidBody(w, requestBody) {
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

	resp, errResp := fetch.ValidateGuestSession(
		fetchRequests.ValidateSession{
			Environment: params.Environment,
			Token: sessionCookie.Value,
		},
		sessionCookie,
	)
	if errResp != nil {
		errors.DefaultResponse(w, errResp)
		return
	}
	if resp == nil {
		errors.BadRequest(w, &responses.Errors{
			Default: &errors.FailedToValidateGuestSession,
		})
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
	if dropRequestNotValidBody(w, requestBody) {
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
	if dropRequestNotValidBody(w, requestBody) {
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

