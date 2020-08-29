package queries

import (
	"encoding/json"
	"net/http"

	"webapi/store/roles/controller"
	"webapi/store/roles/hooks/cache"
	"webapi/store/roles/hooks/errors"
	"webapi/store/roles/hooks/requests"
	"webapi/store/roles/hooks/responses"

	"webapi/infraclientx/fetchx"
	fetchRequests "webapi/infraclientx/fetchx/requests"
	"webapi/infraclientx/verifyx"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"

	InfraOverlordAdmin = "INFRA_OVERLORD_ADMIN"
)

func writeRolesResponse(w http.ResponseWriter, roles *controller.Roles) {
	w.Header().Set(ContentType, ApplicationJson)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&responses.Body{
		Roles: roles,
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

	var params requests.Read
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !verifyx.IsInfraSessionValid(w, params.Environment, sessionCookie) {
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

func ValidateInfra(w http.ResponseWriter, sessionCookie *http.Cookie, requestBody *requests.Body) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}
	if sessionCookie == nil {
		errors.CustomResponse(w, errors.NilInfraCredentials)
		return
	}

	var params fetchRequests.ValidateInfraRole
	bytes, _ := json.Marshal(requestBody.Params)
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	if !verifyx.IsGuestSessionValid(w, params.Environment, sessionCookie) {
		return
	}

	resp, errResp := fetchx.ValidateGuestUser(
		&params,
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
		Environment:  params.Environment,
		UserID:       resp.ID,
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

	if !verifyx.IsInfraSessionValid(w, params.Environment, sessionCookie) {
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

	if !verifyx.IsInfraSessionValid(w, params.Environment, sessionCookie) {
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
