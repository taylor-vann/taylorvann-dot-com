package mutations

import (
	"encoding/json"
	"net/http"
	"webapi/store/roles/hooks/errors"
	"webapi/store/roles/hooks/requests"
	"webapi/store/roles/hooks/responses"
	"webapi/store/roles/controller"
)

func Create(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToCreateRole,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Create)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToCreateRole,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	roles, errCreateSession := controller.Create(&params)
	if errCreateSession != nil {
		errors.DefaultResponse(w, errCreateSession)
		return
	}

	if roles != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Roles: &roles,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToCreateRole,
	})
}

func Update(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToUpdateRole,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Update)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToUpdateRole,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	roles, errUpdateRoles := controller.Update(&params)
	if errUpdateRoles != nil {
		errors.DefaultResponse(w, errUpdateRoles)
		return
	}

	if roles != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Roles: &roles,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToUpdateRole,
	})
}

func UpdateAccess(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToUpdateAccessRole,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.UpdateAccess)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToUpdateAccessRole,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	roles, errUpdateAccessRoles := controller.UpdateAccess(&params)
	if errUpdateAccessRoles != nil {
		errors.DefaultResponse(w, errUpdateAccessRoles)
		return
	}

	if roles != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Roles: &roles,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToUpdateAccessRole,
	})
}

func Delete(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToDeleteRole,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Delete)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToDeleteRole,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	roles, errDeleteRole := controller.Delete(&params)
	if errDeleteRole != nil {
		errors.DefaultResponse(w, errDeleteRole)
		return
	}

	if roles != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Roles: &roles,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToDeleteRole,
	})
}

func Undelete(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToUndeleteRole,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Undelete)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToUndeleteRole,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	roles, errUndeleteRole := controller.Undelete(&params)
	if errUndeleteRole != nil {
		errors.DefaultResponse(w, errUndeleteRole)
		return
	}

	if roles != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Roles: &roles,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Roles: &errors.FailedToUndeleteRole,
	})
}
