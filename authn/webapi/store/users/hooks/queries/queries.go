package queries

import (
	"encoding/json"
	"net/http"

	// "webapi/store/validatesessionx/cookies"
	"webapi/store/validatesessionx"

	"webapi/store/users/controller"
	"webapi/store/users/hooks/cache"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"
)

const SessionCookieHeader = "briantaylorvann.com_internal_session"

func writeUsersResponse(w http.ResponseWriter, users *controller.Users) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&responses.Body{
		Users: users,
	})
}

func Read(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToReadUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Read
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToReadUser,
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
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

func Validate(w http.ResponseWriter, r *http.Request, requestBody *requests.Body) {
	session, errSession := cookies.GetInternalSessionFromRequest(r)
	if errSession != nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.InvalidGuestSession,
		})
		return
	}

	isValidSession, errIsValidSession := validatesessionx.ValidateGuestSession(session)
	if isValidSession == false {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.InvalidGuestSession,
		})
		return
	}

	if errIsValidSession != nil {
		errors.DefaultResponse(w, errIsValidSession)
		return
	}

	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToValidateUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	// check for guest session, you can only validate with a guest session
	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Validate
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToValidateUser,
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
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


func Index(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToIndexUsers,
			Body: &errors.BadRequestFail,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Index
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToIndexUsers,
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
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

func Search(w http.ResponseWriter, requestBody *requests.Body) {
	if requestBody == nil || requestBody.Params == nil {
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToSearchUsers,
			Body: &errors.BadRequestFail,
		})
		return
	}

	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Search
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {
		errAsStr := errParamsMarshal.Error()
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToSearchUsers,
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
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

