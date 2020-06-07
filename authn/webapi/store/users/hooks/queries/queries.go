package queries

import (
	"encoding/json"
	"net/http"

	"webapi/store/clientx/fetch"
	fetchRequests "webapi/store/clientx/fetch/requests"

	"webapi/store/users/controller"
	"webapi/store/users/hooks/cache"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"
)

func writeUsersResponse(w http.ResponseWriter, users *controller.SafeUsers) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&responses.Body{
		Users: users,
	})
}

func dropRequestNotValidBody(w http.ResponseWriter, requestBody *requests.Body) bool {
	if requestBody != nil && requestBody.Params != nil {
		return false
	}
	errors.BadRequest(w, &responses.Errors{
		RequestBody: &errors.BadRequestFail,
	})
	return true
}

func Read(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Read
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	users, errReadUserCache := cache.GetReadEntry(&params)
	if errReadUserCache != nil {
		errors.DefaultResponse(w, errReadUserCache)
		return
	}
	if users != nil {
		writeUsersResponse(w, users)
		return
	}

	usersStore, errUserStore := controller.Read(&params)
	if errUserStore != nil {
		errors.DefaultResponse(w, errUserStore)
		return
	}
	if users != nil {
		writeUsersResponse(w, &usersStore)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToReadUser,
	})
}

func ValidateGuest(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}

	// check for guest session, you can only validate with a guest session
	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Validate
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	// validate guest session
	resp, errResp := fetch.ValidateGuestSession(
		fetchRequests.GuestSession{
			Environment: params.Environment,
		},
		sessionCookie,
	)

	if errResp != nil {
		errors.DefaultResponse(w, errResp)
		return
	}

	if resp == "" {
		errors.BadRequest(w, &responses.Errors{
			Default: &errors.FailedToValidateGuestSession,
		})
		return
	}

	usersStore, errUserStore := controller.Validate(&requests.Validate{
		Environment: params.Environment,
		Email: params.Email,
		Password: params.Password,
	})

	if errUserStore != nil {
		errors.DefaultResponse(w, errUserStore)
		return
	}
	if usersStore != nil {
		writeUsersResponse(w, &usersStore)
		return
	}
	
	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToReadUser,
	})
}


func Index(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Index
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	users, errIndexUsers := controller.Index(&params)
	if errIndexUsers != nil {
		errors.DefaultResponse(w, errIndexUsers)
		return
	}
	if users != nil {
		writeUsersResponse(w, &users)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToIndexUsers,
	})
}

func Search(
	w http.ResponseWriter,
	sessionCookie *http.Cookie,
	requestBody *requests.Body,
) {
	if dropRequestNotValidBody(w, requestBody) {
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Search
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errors.DefaultResponse(w, errParamsMarshal)
		return
	}

	users, errSearchUsers := controller.Search(&params)
	if errSearchUsers != nil {
		errors.DefaultResponse(w, errSearchUsers)
		return
	}
	if users != nil {
		writeUsersResponse(w, &users)
		return
	}

	errors.BadRequest(w, &responses.Errors{
		Users: &errors.FailedToSearchUsers,
	})
}

