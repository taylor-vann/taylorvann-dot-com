package queries

import (
	"encoding/json"
	"net/http"

	"log"
	// "webapi/store/validatesessionx/cookies"
	// "webapi/store/validatesessionx"
	// "webapi/store/infrax/client"

	"webapi/store/users/controller"
	"webapi/store/users/hooks/cache"
	"webapi/store/users/hooks/errors"
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"
)

const SessionCookieHeader = "briantaylorvann.com_session"

func writeUsersResponse(w http.ResponseWriter, users *controller.SafeUsers) {
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

func ValidateGuest(w http.ResponseWriter, r *http.Request, requestBody *requests.Body) {
	// session, errSession := cookies.GetInternalSessionFromRequest(r)
	// if errSession != nil {
	// 	errors.BadRequest(w, &responses.Errors{
	// 		Users: &errors.InvalidGuestSession,
	// 	})
	// 	return
	// }

	// isValidSession, errIsValidSession := validatesessionx.ValidateGuestSession(session)
	// if errIsValidSession != nil {
	// 	errors.DefaultResponse(w, errIsValidSession)
	// 	return
	// }
	// if isValidSession == false {
	// 	errors.BadRequest(w, &responses.Errors{
	// 		Users: &errors.InvalidGuestSession,
	// 	})
	// 	return
	// }

	// client.RemoteLog("made it to Validate guest!")

	log.Println(*r)
	log.Println(r.Host)
	log.Println(*r.URL)
	log.Println(r.RemoteAddr)


	log.Println(r.Cookies())
	log.Println(r.Header.Get("cookie"))
	for _, cookie := range r.Cookies() {
		log.Println(cookie.Value)
	}
	log.Println("Made it to validate guest!")
	if requestBody == nil || requestBody.Params == nil {
		log.Println("Error with request body!")
		log.Println(requestBody)
		if requestBody != nil {
			log.Println(requestBody.Params)
		}
		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToValidateUser,
			Body: &errors.BadRequestFail,
		})
		return
	}

	log.Println("about to marshal request body!")

	// check for guest session, you can only validate with a guest session
	bytes, _ := json.Marshal(requestBody.Params)
	var params requests.Validate
	errParamsMarshal := json.Unmarshal(bytes, &params)
	if errParamsMarshal != nil {

		errAsStr := errParamsMarshal.Error()
		log.Println("error marshalling request bodu!")
		log.Println(errAsStr)

		errors.BadRequest(w, &responses.Errors{
			Users: &errors.FailedToValidateUser,
			Body: &errors.BadRequestFail,
			Default: &errAsStr,
		})
		return
	}

	log.Println("made it past marshal!")

	log.Println("about to validate that user!")

	usersStore, errUserStore := controller.Validate(&requests.Validate{
		Environment: params.Environment,
		Email: params.Email,
		Password: params.Password,
	})
	log.Println("validated the user!!")

	if errUserStore != nil {
		log.Println("error validating the user!!")
		log.Println(errUserStore)

		errors.DefaultResponse(w, errUserStore)
		return
	}
	if usersStore != nil {
		log.Println("about to write a response!")

		writeUsersResponse(w, &usersStore)
		return
	}
	log.Println("reached the end!")

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

