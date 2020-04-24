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
		errors.CustomErrorResponse(w, errors.FailedToReadRole)
		return
	}

	roles, errReadSession := controller.Read(&controller.ReadParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.UserID,
		Organization: requestBody.Params.Organization,
	})

	if errReadSession != nil {
		errors.DefaultErrorResponse(w, errReadSession)
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

	errors.CustomErrorResponse(w, errors.FailedToReadRole)
}

func Index(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, errors.FailedToIndexRole)
		return
	}

	roles, errIndexRoles := controller.Index(controller.IndexParams{
		Environment: requestBody.Params.Environment,
		StartIndex: requestBody.Params.StartIndex,
	})

	if errIndexRoles != nil {
		errors.DefaultErrorResponse(w, errIndexRoles)
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

	errors.CustomErrorResponse(w, errors.FailedToIndexRole)
}

func Search(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.CustomErrorResponse(w, errors.FailedToSearchRole)
		return
	}

	roles, errSearchRoles := controller.Search(&controller.SearchParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.UserID,
		StartIndex: requestBody.Params.StartIndex,
		Length:	requestBody.Params.Length,
	})

	if errSearchRoles != nil {
		errors.DefaultErrorResponse(w, errSearchRoles)
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

	errors.CustomErrorResponse(w, errors.FailedToSearchRole)
}

