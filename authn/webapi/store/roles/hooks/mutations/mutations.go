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
	if requestBody == nil {
		errors.BadRequest(w, responses.Errors{
			Roles: errors.FailedToCreateRole,
			Body: errors.BadRequestFail
		})
		return
	}

	roles, errCreateSession := controller.Create(&controller.CreateParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.UserID,
		Organization: requestBody.Params.Organization,
		ReadAccess: requestBody.Params.ReadAccess,
		WriteAccess: requestBody.Params.WriteAccess,
	})

	if errCreateSession != nil {
		errors.DefaultErrorResponse(w, errCreateSession)
		return
	}

	if roles != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Roles: roles,
		})
		return
	}

	errors.BadRequest(w, responses.Errors{
		Roles: errors.FailedToCreateRole,
	})
}

func Update(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, responses.Errors{
			Roles: errors.FailedToUpdateRole,
			Body: errors.BadRequestFail
		})
		return
	}

	roles, errUpdateRoles := controller.Update(controller.UpdateParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.UserID,
		Organization: requestBody.Params.Organization,
		ReadAccess: requestBody.Params.ReadAccess,
		WriteAccess: requestBody.Params.WriteAccess,
		IsDeleted: requestBody.Params.IsDeleted,
	})

	if errUpdateRoles != nil {
		errors.DefaultErrorResponse(w, errUpdateRoles)
		return
	}

	if roles != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Roles: roles,
		})
		return
	}

	errors.BadRequest(w, responses.Errors{
		Roles: errors.FailedToUpdateRole,
	})
}

func UpdateAccess(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, responses.Errors{
			Roles: errors.FailedToUpdateAccessRole,
			Body: errors.BadRequestFail
		})
		return
	}

	roles, errUpdateAccessRoles := controller.UpdateAccess(
		&controller.UpdateAccessParams{
			Environment: requestBody.Params.Environment,
			UserID: requestBody.Params.UserID,
			Organization: requestBody.Params.Organization,
			ReadAccess: requestBody.Params.ReadAccess,
			WriteAccess: requestBody.Params.WriteAccess,
		},
	)

	if errUpdateAccessRoles != nil {
		errors.DefaultErrorResponse(w, errUpdateAccessRoles)
		return
	}

	if roles != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Roles: roles,
		})
		return
	}

	errors.BadRequest(w, responses.Errors{
		Roles: errors.FailedToUpdateAccessRole,
	})
}

func Delete(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, responses.Errors{
			Roles: errors.FailedToDeleteRole,
			Body: errors.BadRequestFail
		})
		return
	}

	roles, errDeleteRole := controller.Delete(&controller.DeleteParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.UserID,
		Organization: requestBody.Params.Organization,
	})

	if errDeleteRole != nil {
		errors.DefaultErrorResponse(w, errDeleteRole)
		return
	}

	if roles != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Roles: roles,
		})
		return
	}

	errors.BadRequest(w, responses.Errors{
		Roles: errors.FailedToDeleteRole,
	})
}

func Undelete(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, responses.Errors{
			Roles: errors.FailedToUndeleteRole,
			Body: errors.BadRequestFail
		})
		return
	}

	roles, errUndeleteRole := controller.Undelete(&controller.UndeleteParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.UserID,
		Organization: requestBody.Params.Organization,
	})

	if errUndeleteRole != nil {
		errors.DefaultErrorResponse(w, errUndeleteRole)
		return
	}

	if roles != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Roles: roles,
		})
		return
	}

	errors.BadRequest(w, responses.Errors{
		Roles: errors.FailedToUndeleteRole,
	})
}
