package queries

import (
	"encoding/json"
	"net/http"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"
	"webapi/store/users/controller"
)

func Read(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToReadUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Read)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToReadUser,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	users, errReadUser := controller.Read(&params)
	if errReadUser != nil {
		errors.DefaultResponse(w, errReadUser)
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
		Users: &errors.FailedToReadUser,
	})
}

func Index(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToIndexUsers,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Index)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToIndexUsers,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	users, errIndexUsers := controller.Index(&params)
	if errIndexUsers != nil {
		errors.DefaultResponse(w, errIndexUsers)
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
		Users: &errors.FailedToIndexUsers,
	})
}

func Search(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToSearchUsers,
			Body: &errors.BadRequestFail,
		})
		return
	}

	params, errParams := requestBody.Params.(requests.Search)
	if errParams == false {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToSearchUsers,
			Body: &errors.BadRequestFail,
			Default: &errors.UnrecognizedParams,
		})
		return
	}

	users, errSearchUsers := controller.Search(&params)
	if errSearchUsers != nil {
		errors.DefaultResponse(w, errSearchUsers)
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
		Users: &errors.FailedToSearchUsers,
	})
}

