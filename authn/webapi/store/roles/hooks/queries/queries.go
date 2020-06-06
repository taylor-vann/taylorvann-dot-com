package queries

import (
	"encoding/json"
	"net/http"

	"webapi/store/clientx/fetch"
	fetchRequests "webapi/store/clientx/fetch/requests"

	"log"

	"webapi/store/roles/controller"
	"webapi/store/roles/hooks/cache"
	"webapi/store/roles/hooks/errors"
	"webapi/store/roles/hooks/requests"
	"webapi/store/roles/hooks/responses"
)

const InfraOverlordAdmin = "INFRA_OVERLORD_ADMIN"

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

func ValidateInfra(w http.ResponseWriter, sessionCookie *http.Cookie, requestBody *requests.Body)  {
	log.Println("ROLES QUERY - MADE IT TO VALIDATE INFRA ROLE")
	log.Println("ROLES QUERY - cookie")
	log.Println(sessionCookie)	
	// drop if guest session is not valid
	// need role 
	if requestBody == nil || requestBody.Params == nil {
		log.Println("Queries ValidateInfra -  bad body")
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToReadRole,
			Body: &errors.BadRequestFail,
		})
		return
	}

	// digest body interface{}
	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.ValidateInfra
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		log.Println("Queries ValidateInfra -  error unmarshaling body")
		errors.DefaultResponse(w, errParamsMarshal)

		return
	}

	log.Println("requesting role")

	resp, errResp := fetch.ValidateGuestUser(
		fetchRequests.ValidateGuestUser(params),
		sessionCookie,
	)
	if errResp != nil {
		log.Println("Queries ValidateInfra -  error validating user")
		errors.DefaultResponse(w, errResp)
		return
	}
	if resp == nil {
		log.Println("Queries ValidateInfra -  nil users returned")
		errors.BadRequest(w, nil)
		return
	}
	log.Println("Queries ValidateInfra - validated user")

	roleParams := requests.Read{
		Environment: params.Environment,
		UserID: resp.ID,
		Organization: InfraOverlordAdmin,
	}
	// check cache
	roles, errReadRolesCache := cache.GetReadEntry(&roleParams)
	if errReadRolesCache != nil {
		log.Println("error reading cache")
		errors.DefaultResponse(w, errReadRolesCache)
		return
	}
	if roles != nil {
		log.Println("we found the role!")
		writeRolesResponse(w, roles)
		return
	}
	
	// check store
	rolesStore, errRolesStore := controller.Read(&roleParams)
	if errRolesStore != nil {
		log.Println("we failed to find role")
		log.Println(errRolesStore)

		errors.DefaultResponse(w, errRolesStore)
		return
	}
	if rolesStore != nil {
		log.Println("we found the role!")
		// write to cache
		writeRolesResponse(w, &rolesStore)
		return
	}
	log.Println("unable to find roles")

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

