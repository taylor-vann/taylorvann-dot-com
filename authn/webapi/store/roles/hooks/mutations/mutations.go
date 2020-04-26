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
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToCreateUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	roles, errCreateSession := controller.Create(&controller.CreateParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.Create.UserID,
		Organization: requestBody.Params.Create.Organization,
		ReadAccess: requestBody.Params.Create.ReadAccess,
		WriteAccess: requestBody.Params.Create.WriteAccess,
	})

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
		Roles: &errors.FailedToCreateUser,
	})
}

func Update(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToUpdateUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	roles, errUpdateRoles := controller.Update(&controller.UpdateParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.Update.UserID,
		Organization: requestBody.Params.Update.Organization,
		ReadAccess: requestBody.Params.Update.ReadAccess,
		WriteAccess: requestBody.Params.Update.WriteAccess,
		IsDeleted: requestBody.Params.Update.IsDeleted,
	})

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
		Roles: &errors.FailedToUpdateUser,
	})
}

func UpdateAccess(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToUpdateAccessUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	roles, errUpdateAccessRoles := controller.UpdateAccess(
		&controller.UpdateAccessParams{
			Environment: requestBody.Params.Environment,
			UserID: requestBody.Params.UpdateAccess.UserID,
			Organization: requestBody.Params.UpdateAccess.Organization,
			ReadAccess: requestBody.Params.UpdateAccess.ReadAccess,
			WriteAccess: requestBody.Params.UpdateAccess.WriteAccess,
		},
	)

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
		Roles: &errors.FailedToUpdateAccessUser,
	})
}

func Delete(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToDeleteUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	roles, errDeleteRole := controller.Delete(&controller.DeleteParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.Delete.UserID,
		Organization: requestBody.Params.Delete.Organization,
	})

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
		Roles: &errors.FailedToDeleteUser,
	})
}

func Undelete(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToUndeleteUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	roles, errUndeleteRole := controller.Undelete(&controller.UndeleteParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.Undelete.UserID,
		Organization: requestBody.Params.Undelete.Organization,
	})

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
		Roles: &errors.FailedToUndeleteUser,
	})
}
