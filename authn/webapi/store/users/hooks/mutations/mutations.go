package mutations

import (
	"encoding/json"
	"net/http"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"
	"webapi/store/users/controller"
)

func Create(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, errors.FailedToCreateRole)
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

	errors.CustomErrorResponse(w, errors.FailedToCreateRole)
}

func Update(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, errors.FailedToUpdateRole)
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

	errors.CustomErrorResponse(w, errors.FailedToUpdateRole)
}

func UpdateAccess(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, errors.FailedToUpdateAccessRole)
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

	errors.CustomErrorResponse(w, errors.FailedToUpdateAccessRole)
}

func Delete(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, errors.FailedToDeleteRole)
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

	errors.CustomErrorResponse(w, errors.FailedToDeleteRole)
}

func Undelete(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, errors.FailedToUndeleteRole)
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

	errors.CustomErrorResponse(w, errors.FailedToUndeleteRole)
}
