package queries

import (
	"encoding/json"
	"net/http"

	"webapi/store/roles/controller"
	"webapi/store/roles/hooks/cache"
	"webapi/store/roles/hooks/errors"
	"webapi/store/roles/hooks/requests"
	"webapi/store/roles/hooks/responses"
)

func writeRolesResponse(w http.ResponseWriter, roles *controller.Roles) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&responses.Body{
		Roles: roles,
	})
}

func Read(w http.ResponseWriter, requestBody *requests.Body)  {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToReadRole,
			Body: &errors.BadRequestFail,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Read
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToReadRole,
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
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
		writeRolesResponse(w, &rolesStore)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToReadRole,
	})
}

func ValidateInfra(w http.ResponseWriter, requestBody *requests.Body)  {
	// drop if session is not valid
	
	// infrax validate guest user
	
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToReadRole,
			Body: &errors.BadRequestFail,
		})
		return
	}

	// digest body interface{}
	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Read
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToReadRole,
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
		return
	}

	// check cache
	roles, errReadUserCache := cache.GetReadEntry(&params)
	if errReadUserCache != nil {
		errors.DefaultResponse(w, errReadUserCache)
		return
	}
	if roles != nil {
		writeRolesResponse(w, roles)
		return
	}

	// check store
	rolesStore, errRolesStore := controller.Read(&params)
	if errRolesStore != nil {
		errors.DefaultResponse(w, errRolesStore)
		return
	}
	if rolesStore != nil {
		writeRolesResponse(w, &rolesStore)
		return
	}

	// default error
	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToReadRole,
	})
}


func Index(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToIndexRoles,
			Body: &errors.BadRequestFail,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Index
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToIndexRoles,
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
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

func Search(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToSearchRoles,
			Body: &errors.BadRequestFail,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Search
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToIndexRoles,
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
		return
	}

	roles, errRoles := controller.Search(&params)
	if errRoles != nil {
		errAsStr := errRoles.Error()
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToReadRole,
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
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

