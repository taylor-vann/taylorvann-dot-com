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
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToCreateUser,
			Body: &errors.BadRequestFail,
		})
	}

	users, errCreateUser := controller.Create(&controller.CreateParams{
		Environment: requestBody.Params.Environment,
		Email: requestBody.Params.Create.Email,
		Password: requestBody.Params.Create.Password,
	})

	if errCreateUser != nil {
		errors.DefaultResponse(w, errCreateUser)
		return
	}

	if users != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Users: &users,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToCreateUser,
	})
}

func Update(w http.ResponseWriter, requestBody *requests.Body) {
	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToUpdateUser,
		Body: &errors.BadRequestFail,
	})

	users, errUpdateUsers := controller.Update(&controller.UpdateParams{
		Environment: requestBody.Params.Environment,
		CurrentEmail: requestBody.Params.Update.CurrentEmail,
		UpdatedEmail: requestBody.Params.Update.UpdatedEmail,
		Password: requestBody.Params.Update.Password,
		IsDeleted: requestBody.Params.Update.IsDeleted,
	})

	if errUpdateUsers != nil {
		errors.DefaultResponse(w, errUpdateUsers)
		return
	}

	if users != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Users: &users,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToUpdateUser,
	})
}

func UpdateEmail(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUpdateEmailUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	users, errUpdateEmailUser := controller.UpdateEmail(
		&controller.UpdateEmailParams{
			Environment: requestBody.Params.Environment,
			CurrentEmail: requestBody.Params.UpdateEmail.CurrentEmail,
			UpdatedEmail: requestBody.Params.UpdateEmail.UpdatedEmail,
		},
	)

	if errUpdateEmailUser != nil {
		errors.DefaultResponse(w, errUpdateEmailUser)
		return
	}

	if users != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Users: &users,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToUpdateEmailUser,
	})
}

func UpdatePassword(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUpdatePasswordUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	users, errUpdatePasswordUser := controller.UpdatePassword(
		&controller.UpdatePasswordParams{
			Environment: requestBody.Params.Environment,
			Email: requestBody.Params.UpdatePassword.Email,
			Password: requestBody.Params.UpdatePassword.Password,
		},
	)

	if errUpdatePasswordUser != nil {
		errors.DefaultResponse(w, errUpdatePasswordUser)
		return
	}

	if users != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Users: &users,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToUpdatePasswordUser,
	})
}

func Delete(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToDeleteUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	users, errDeleteUsers := controller.Delete(&controller.DeleteParams{
		Environment: requestBody.Params.Environment,
		Email: requestBody.Params.Delete.Email,
	})

	if errDeleteUsers != nil {
		errors.DefaultResponse(w, errDeleteUsers)
		return
	}

	if users != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Users: &users,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToDeleteUser,
	})
}

func Undelete(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUndeleteUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	users, errUndeleteUsers := controller.Undelete(&controller.UndeleteParams{
		Environment: requestBody.Params.Environment,
		Email: requestBody.Params.Undelete.Email,
	})

	if errUndeleteUsers != nil {
		errors.DefaultResponse(w, errUndeleteUsers)
		return
	}

	if users != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&responses.Body{
			Users: &users,
		})
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToUndeleteUser,
	})
}
