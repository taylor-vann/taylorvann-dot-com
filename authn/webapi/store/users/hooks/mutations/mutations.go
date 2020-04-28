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
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToCreateUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Create)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToCreateUser,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	users, errCreateUser := controller.Create(&params)
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
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUpdateUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Update)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUpdateUser,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	users, errUpdateUsers := controller.Update(&params)
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
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUpdateEmailUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.UpdateEmail)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUpdateEmailUser,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	users, errUpdateEmailUser := controller.UpdateEmail(&params)
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
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUpdatePasswordUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.UpdatePassword)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUpdatePasswordUser,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	users, errUpdatePasswordUser := controller.UpdatePassword(&params)
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
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToDeleteUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Delete)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToDeleteUser,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	users, errDeleteUsers := controller.Delete(&params)
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
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUndeleteUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Undelete)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToUndeleteUser,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	users, errUndeleteUsers := controller.Undelete(&params)
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
