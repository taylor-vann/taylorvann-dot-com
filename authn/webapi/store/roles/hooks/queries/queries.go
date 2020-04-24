package queries

import (
	"encoding/json"
	"net/http"
	"webapi/store/roles/hooks/errors"
	"webapi/store/roles/hooks/requests"
	"webapi/store/roles/hooks/responses"
	"webapi/store/roles/controller"
)

func Read(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, responses.Errors{
			Roles: errors.FailedToReadRole,
			Body: errors.BadRequestFail
		})
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

	errors.BadRequest(w, responses.Errors{
		Roles: errors.FailedToReadRole,
	})
}

func Index(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, responses.Errors{
			Roles: errors.FailedToIndexRole,
			Body: errors.BadRequestFail
		})
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

	errors.BadRequest(w, responses.Errors{
		Roles: errors.FailedToIndexRole,
	})
}

func Search(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, responses.Errors{
			Roles: errors.FailedToSearchRoles,
			Body: errors.BadRequestFail
		})
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

	errors.BadRequest(w, responses.Errors{
		Roles: errors.FailedToSearchRoles,
	})
}

