package mutations

import (
	"encoding/json"
	"net/http"

	"webapi/infraclientx/verifyx"
	"webapi/store/roles/controller"
	"webapi/store/roles/hooks/cache"
	"webapi/store/roles/hooks/errors"
	"webapi/store/roles/hooks/requests"
	"webapi/store/roles/hooks/responses"
)

func writeRolesResponse(w http.ResponseWriter, roles *controller.Roles) {
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

	roles, errCreateSession := controller.Create(&params)
	if errCreateSession != nil {
		errors.DefaultResponse(w, errCreateSession)
		return
	}

	if roles != nil {
		cache.UpdateReadEntry(params.Environment, &roles)
		writeRolesResponse(w, &roles)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToCreateRole,
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

	roles, errUpdateRoles := controller.Update(&params)
	if errUpdateRoles != nil {
		errors.DefaultResponse(w, errUpdateRoles)
		return
	}

	if roles != nil {
		cache.UpdateReadEntry(params.Environment, &roles)
		writeRolesResponse(w, &roles)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToUpdateRole,
	})
}

func UpdateAccess(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if !isRequestBodyValid(w, requestBody) {
		return
	}

	var params requests.UpdateAccess
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

	roles, errUpdateAccessRoles := controller.UpdateAccess(&params)
	if errUpdateAccessRoles != nil {
		errors.DefaultResponse(w, errUpdateAccessRoles)
		return
	}

	if roles != nil {
		cache.UpdateReadEntry(params.Environment, &roles)
		writeRolesResponse(w, &roles)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToUpdateAccessRole,
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

	roles, errDeleteRole := controller.Delete(&params)
	if errDeleteRole != nil {
		errors.DefaultResponse(w, errDeleteRole)
		return
	}

	if roles != nil {
		cache.UpdateReadEntry(params.Environment, &roles)
		writeRolesResponse(w, &roles)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToDeleteRole,
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

	var params requests.Read
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

	roles, errUndeleteRole := controller.Undelete(&params)
	if errUndeleteRole != nil {
		errors.DefaultResponse(w, errUndeleteRole)
		return
	}

	if roles != nil {
		cache.UpdateReadEntry(params.Environment, &roles)
		writeRolesResponse(w, &roles)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToUndeleteRole,
	})
}
