package queries

import (
	"encoding/json"
	"net/http"

	"github.com/taylor-vann/weblog/toolbox/golang/clientx"
	"github.com/taylor-vann/weblog/toolbox/golang/clientx/fetch"
	fetchRequests "github.com/taylor-vann/weblog/toolbox/golang/clientx/fetch/requests"

	"webapi/store/roles/controller"
	"webapi/store/roles/hooks/cache"
	"webapi/store/roles/hooks/errors"
	"webapi/store/roles/hooks/requests"
	"webapi/store/roles/hooks/responses"

	"github.com/taylor-vann/weblog/toolbox/golang/jwtx"
)

const InfraOverlordAdmin = "INFRA_OVERLORD_ADMIN"

func writeRolesResponse(w http.ResponseWriter, roles *controller.Roles) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&responses.Body{
		Roles: roles,
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
)  {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params requests.Read
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

	roles, errReadUserCache := cache.GetReadEntry(&params)
	if errReadUserCache != nil {
		errors.DefaultResponse(w, errReadUserCache)
		return
	}
	if roles != nil {
		writeRolesResponse(w, roles)
		return
	}

	rolesStore, errRolesStore := controller.Read(&params)
	if errRolesStore != nil {
		errors.DefaultResponse(w, errRolesStore)
		return
	}
	if rolesStore != nil {
		cache.UpdateReadEntry(params.Environment, &rolesStore)
		writeRolesResponse(w, &rolesStore)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToReadRole,
	})
}

func ValidateInfra(w http.ResponseWriter, sessionCookie *http.Cookie, requestBody *requests.Body)  {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}
	
	var params requests.ValidateInfra
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	resp, errResp := fetch.ValidateGuestUser(
		fetchRequests.ValidateGuestUser(params),
		sessionCookie,
	)
	if errResp != nil {
		errors.DefaultResponse(w, errResp)
		return
	}
	if resp == nil {
		errors.BadRequest(w, nil)
		return
	}
	
	roleParams := requests.Read{
		Environment: params.Environment,
		UserID: resp.ID,
		Organization: InfraOverlordAdmin,
	}

	// check cache
	roles, errReadRolesCache := cache.GetReadEntry(&roleParams)
	if errReadRolesCache != nil {
		errors.DefaultResponse(w, errReadRolesCache)
		return
	}
	if roles != nil {
		writeRolesResponse(w, roles)
		return
	}
	
	// check store
	rolesStore, errRolesStore := controller.Read(&roleParams)
	if errRolesStore != nil {
		errors.DefaultResponse(w, errRolesStore)
		return
	}
	if rolesStore != nil {
		cache.UpdateReadEntry(params.Environment, &rolesStore)
		writeRolesResponse(w, &rolesStore)
		return
	}
	
	// default error
	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToReadRole,
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

	roles, errIndexRoles := controller.Index(&params)
	if errIndexRoles != nil {
		errors.DefaultResponse(w, errIndexRoles)
		return
	}
	if roles != nil {
		writeRolesResponse(w, &roles)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToIndexRoles,
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

	roles, errRoles := controller.Search(&params)
	if errRoles != nil {
		errors.DefaultResponse(w, errRoles)
		return
	}
	if roles != nil {
		writeRolesResponse(w, &roles)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToSearchRoles,
	})
}

