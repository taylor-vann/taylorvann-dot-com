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
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToReadUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	users, errReadRole := controller.Read(&controller.ReadParams{
		Environment: requestBody.Params.Environment,
		Email: requestBody.Params.Read.Email,
	})

	if errReadRole != nil {
		errors.DefaultResponse(w, errReadRole)
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
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToIndexUsers,
			Body: &errors.BadRequestFail,
		})
	}

	users, errIndexRoles := controller.Index(&controller.IndexParams{
		Environment: requestBody.Params.Index.Environment,
		StartIndex: requestBody.Params.Index.StartIndex,
		Length: requestBody.Params.Index.Length,
	})

	if errIndexRoles != nil {
		errors.DefaultResponse(w, errIndexRoles)
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
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToSearchUsers,
			Body: &errors.BadRequestFail,
		})
		return
	}

	users, errSearchRoles := controller.Search(&controller.SearchParams{
		Environment: requestBody.Params.Environment,
		EmailSubstring: requestBody.Params.Search.EmailSubstring,
		StartIndex: requestBody.Params.Search.StartIndex,
		Length:	requestBody.Params.Search.Length,
	})

	if errSearchRoles != nil {
		errors.DefaultResponse(w, errSearchRoles)
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

