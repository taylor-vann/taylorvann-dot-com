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
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToReadRole,
			Body: &errors.BadRequestFail,
		})
		return
	}

	roles, errReadSession := controller.Read(&controller.ReadParams{
		Environment: requestBody.Params.Environment,
		UserID: requestBody.Params.Read.UserID,
		Organization: requestBody.Params.Read.Organization,
	})

	if errReadSession != nil {
		errors.DefaultResponse(w, errReadSession)
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
		Roles: &errors.FailedToReadRole,
	})
}

func Index(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToIndexRoles,
			Body: &errors.BadRequestFail,
		})
		return
	}

	roles, errIndexRoles := controller.Index(&controller.IndexParams{
		Environment: requestBody.Params.Index.Environment,
		StartIndex: requestBody.Params.Index.StartIndex,
		Length: requestBody.Params.Index.Length,
	})

	if errIndexRoles != nil {
		errors.DefaultResponse(w, errIndexRoles)
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
		Roles: &errors.FailedToIndexRoles,
	})
}

func Search(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil {
		errors.BadRequest(w, &responses.Errors{
			Roles: &errors.FailedToSearchRoles,
			Body: &errors.BadRequestFail,
		})
		return
	}

	roles, errSearchRoles := controller.Search(&controller.SearchParams{
		Environment: requestBody.Params.Search.Environment,
		UserID: requestBody.Params.Search.UserID,
		StartIndex: requestBody.Params.Search.StartIndex,
		Length:	requestBody.Params.Search.Length,
	})

	if errSearchRoles != nil {
		errors.DefaultResponse(w, errSearchRoles)
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
		Roles: &errors.FailedToSearchRoles,
	})
}

