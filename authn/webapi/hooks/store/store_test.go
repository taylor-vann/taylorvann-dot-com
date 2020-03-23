package store

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"webapi/hooks/constants"
	"webapi/hooks/sessions"
	"webapi/hooks/store/mutations"
	"webapi/store"
)

func TestCreateUserBadRequest(t *testing.T) {
	resp, errResp := http.NewRequest(
		"POST",
		"/store/m/",
		nil,
	)
	if errResp != nil {
		t.Error("error making guest session request")
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	status := httpTest.Code
	if status != http.StatusBadRequest {
		t.Error("handler returned incorrect status code, should be 400")
	}
}

func TestCreateUserBadRequestBody(t *testing.T) {
	userParams := store.CreateUserParams{
		Email:    "brian.t.vann@gmail.com",
		Password: "passwerd",
	}

	requestBody := mutations.CreateUserRequestBody{
		Action: "CREATE_USER",
		Params: mutations.CreateUserParams{
			User: userParams,
		},
	}
	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)

	req, errReq := http.NewRequest(
		"POST",
		"/store/m/",
		marshalBytes,
	)
	if errReq != nil {
		t.Error("error making guest session request")
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, req)

	status := httpTest.Code
	if status != http.StatusBadRequest {
		t.Error("handler returned incorrect status code, should be 400")
	}
}

func TestCreateUser(t *testing.T) {
	requestBodyGuest := RequestBody{
		Action: sessions.CreateGuestSession,
	}

	marshalBytesGuest := new(bytes.Buffer)
	json.NewEncoder(marshalBytesGuest).Encode(requestBodyGuest)
	resp, errResp := http.NewRequest(
		"POST",
		"/sessions/m/",
		marshalBytesGuest,
	)
	if errResp != nil {
		t.Error("error making guest session request")
	}

	httpTestGuest := httptest.NewRecorder()
	handlerGuest := http.HandlerFunc(sessions.Mutation)
	handlerGuest.ServeHTTP(httpTestGuest, resp)

	statusGuest := httpTestGuest.Code
	if statusGuest != http.StatusOK {
		t.Error("handler returned incorrect status code")
	}

	userParams := store.CreateUserParams{
		Email:    "brian.t.vann@gmail.com",
		Password: "passwerd",
	}

	// not out of the head anymore
	sessionToken := httpTestGuest.Header().Get(constants.SessionTokenHeader)
	csrfToken := httpTestGuest.Header().Get(constants.CsrfTokenHeader)
	sessionParams := mutations.SessionParams{
		SessionToken: &sessionToken,
		CsrfToken: &csrfToken,
	}

	requestBody := mutations.CreateUserRequestBody{
		Action: "CREATE_USER",
		Params: mutations.CreateUserParams{
			User: userParams,
			Session: sessionParams,
		},
	}
	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)

	req, errReq := http.NewRequest(
		"POST",
		"/store/m/",
		marshalBytes,
	)
	if errReq != nil {
		t.Error("error making guest session request")
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, req)

	status := httpTest.Code
	if status != http.StatusOK {
		t.Error("handler returned incorrect status code, should be 200")
	}
	t.Error("fail because we're new")
}

// func TestUpdateUserEmail(t *testing.T) {
// 	// get gets headers
// 	requestBodyGuest := RequestBodyParams{
// 		Action: sessions.CreateGuestSession,
// 	}

// 	marshalBytesGuest := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytesGuest).Encode(requestBodyGuest)
// 	resp, errResp := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytesGuest,
// 	)
// 	if errResp != nil {
// 		t.Error("error making guest session request")
// 	}

// 	httpTestGuest := httptest.NewRecorder()
// 	handlerGuest := http.HandlerFunc(sessions.Mutation)
// 	handlerGuest.ServeHTTP(httpTestGuest, resp)

// 	statusGuest := httpTestGuest.Code
// 	if statusGuest != http.StatusOK {
// 		t.Error("handler returned incorrect status code")
// 	}

// 	userParams := store.CreateUserParams{
// 		Email:    "brian.t.vann@gmail.com",
// 		Password: "passwerd",
// 	}

// 	// Sign user into public Sessions
// 	requestBody := mutations.CreateUserRequestBodyParams{
// 		Action: sessions.CreatePublicSession,
// 		Params: userParams,
// 	}
// 	marshalBytes := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytes).Encode(requestBody)

// 	req, errReq := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytes,
// 	)
// 	if errReq != nil {
// 		t.Error("error making guest session request")
// 	}
// 	req.Header.Set(
// 		constants.SessionTokenHeader,
// 		httpTestGuest.Header().Get(constants.SessionTokenHeader),
// 	)

// 	req.Header.Set(
// 		constants.CsrfTokenHeader,
// 		httpTestGuest.Header().Get(constants.CsrfTokenHeader),
// 	)

// 	httpTest := httptest.NewRecorder()
// 	handler := http.HandlerFunc(sessions.Mutation)
// 	handler.ServeHTTP(httpTest, req)

// 	status := httpTest.Code
// 	if status != http.StatusOK {
// 		t.Error("handler returned incorrect status code, should be 200")
// 		fmt.Println(httpTest.Body)
// 	}
// 	fmt.Println("signed the user in")
// 	t.Error("fail because we're new")

// 	// update request
// 	userUpdateParams := store.UpdateEmailParams{
// 		CurrentEmail: "brian.t.vann@gmail.com",
// 		UpdatedEmail: "brian.t.vann2@gmail.com",
// 	}

// 	requestBodyUpdate := mutations.CreateUpdateUserEmailRequestBodyParams{
// 		Action: "UPDATE_USER_EMAIL",
// 		Params: userUpdateParams,
// 	}
// 	marshalBytesUpdate := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytesUpdate).Encode(requestBodyUpdate)

// 	reqUpdate, errReqUpdate := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytesUpdate,
// 	)
// 	if errReqUpdate != nil {
// 		t.Error("error making guest session request")
// 	}
// 	reqUpdate.Header.Set(
// 		constants.SessionTokenHeader,
// 		httpTest.Header().Get(constants.SessionTokenHeader),
// 	)

// 	reqUpdate.Header.Set(
// 		constants.CsrfTokenHeader,
// 		httpTest.Header().Get(constants.CsrfTokenHeader),
// 	)

// 	httpTestUpdate := httptest.NewRecorder()
// 	// store mutation
// 	handlerUpdate := http.HandlerFunc(Mutation)
// 	handlerUpdate.ServeHTTP(httpTestUpdate, reqUpdate)

// 	statusUpdate := httpTestUpdate.Code
// 	if statusUpdate != http.StatusOK {
// 		t.Error("handler returned incorrect status code, should be 200")
// 		fmt.Println(httpTestUpdate.Body)
// 	}
// }

// // update password
// func TestUpdateUserPassword(t *testing.T) {
// 	// get gets headers
// 	requestBodyGuest := RequestBodyParams{
// 		Action: sessions.CreateGuestSession,
// 	}

// 	marshalBytesGuest := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytesGuest).Encode(requestBodyGuest)
// 	resp, errResp := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytesGuest,
// 	)
// 	if errResp != nil {
// 		t.Error("error making guest session request")
// 	}

// 	httpTestGuest := httptest.NewRecorder()
// 	handlerGuest := http.HandlerFunc(sessions.Mutation)
// 	handlerGuest.ServeHTTP(httpTestGuest, resp)

// 	statusGuest := httpTestGuest.Code
// 	if statusGuest != http.StatusOK {
// 		t.Error("handler returned incorrect status code")
// 	}

// 	userParams := store.CreateUserParams{
// 		Email:    "brian.t.vann@gmail.com",
// 		Password: "passwerd",
// 	}

// 	// Sign user into public Sessions
// 	requestBody := mutations.CreateUserRequestBodyParams{
// 		Action: sessions.CreatePublicSession,
// 		Params: userParams,
// 	}
// 	marshalBytes := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytes).Encode(requestBody)

// 	req, errReq := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytes,
// 	)
// 	if errReq != nil {
// 		t.Error("error making guest session request")
// 	}
// 	req.Header.Set(
// 		constants.SessionTokenHeader,
// 		httpTestGuest.Header().Get(constants.SessionTokenHeader),
// 	)

// 	req.Header.Set(
// 		constants.CsrfTokenHeader,
// 		httpTestGuest.Header().Get(constants.CsrfTokenHeader),
// 	)

// 	httpTest := httptest.NewRecorder()
// 	handler := http.HandlerFunc(sessions.Mutation)
// 	handler.ServeHTTP(httpTest, req)

// 	status := httpTest.Code
// 	if status != http.StatusOK {
// 		t.Error("handler returned incorrect status code, should be 200")
// 		fmt.Println(httpTest.Body)
// 	}
// 	fmt.Println("signed the user in")
// 	t.Error("fail because we're new")

// 	// update request
// 	userUpdateParams := mutations.UpdatePasswordRequestParams{
// 		Email: "brian.t.vann2@gmail.com",
// 		Password: "passwerd",
// 		UpdatedPassword: "passw3rd",
// 	}

// 	requestBodyUpdate := mutations.UpdateUserPasswordRequestBody{
// 		Action: "UPDATE_USER_PASSWORD",
// 		Params: userUpdateParams,
// 	}
// 	marshalBytesUpdate := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytesUpdate).Encode(requestBodyUpdate)

// 	reqUpdate, errReqUpdate := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytesUpdate,
// 	)
// 	if errReqUpdate != nil {
// 		t.Error("error making guest session request")
// 	}
// 	reqUpdate.Header.Set(
// 		constants.SessionTokenHeader,
// 		httpTest.Header().Get(constants.SessionTokenHeader),
// 	)

// 	reqUpdate.Header.Set(
// 		constants.CsrfTokenHeader,
// 		httpTest.Header().Get(constants.CsrfTokenHeader),
// 	)

// 	httpTestUpdate := httptest.NewRecorder()
// 	// store mutation
// 	handlerUpdate := http.HandlerFunc(Mutation)
// 	handlerUpdate.ServeHTTP(httpTestUpdate, reqUpdate)

// 	statusUpdate := httpTestUpdate.Code
// 	if statusUpdate != http.StatusOK {
// 		t.Error("handler returned incorrect status code, should be 200")
// 		fmt.Println(httpTestUpdate.Body)
// 	}
// }

// // remove user (soft delete user)
// func TestRemoveUser(t *testing.T) {
// 	// get gets headers
// 	requestBodyGuest := RequestBodyParams{
// 		Action: sessions.CreateGuestSession,
// 	}

// 	marshalBytesGuest := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytesGuest).Encode(requestBodyGuest)
// 	resp, errResp := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytesGuest,
// 	)
// 	if errResp != nil {
// 		t.Error("error making guest session request")
// 	}

// 	httpTestGuest := httptest.NewRecorder()
// 	handlerGuest := http.HandlerFunc(sessions.Mutation)
// 	handlerGuest.ServeHTTP(httpTestGuest, resp)

// 	statusGuest := httpTestGuest.Code
// 	if statusGuest != http.StatusOK {
// 		t.Error("handler returned incorrect status code")
// 	}

// 	userParams := store.CreateUserParams{
// 		Email:    "brian.t.vann2@gmail.com",
// 		Password: "passw3rd",
// 	}

// 	// Sign user into public Sessions
// 	requestBody := mutations.CreateUserRequestBodyParams{
// 		Action: sessions.CreatePublicSession,
// 		Params: userParams,
// 	}
// 	marshalBytes := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytes).Encode(requestBody)

// 	req, errReq := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytes,
// 	)
// 	if errReq != nil {
// 		t.Error("error making guest session request")
// 	}
// 	req.Header.Set(
// 		constants.SessionTokenHeader,
// 		httpTestGuest.Header().Get(constants.SessionTokenHeader),
// 	)

// 	req.Header.Set(
// 		constants.CsrfTokenHeader,
// 		httpTestGuest.Header().Get(constants.CsrfTokenHeader),
// 	)

// 	httpTest := httptest.NewRecorder()
// 	handler := http.HandlerFunc(sessions.Mutation)
// 	handler.ServeHTTP(httpTest, req)

// 	status := httpTest.Code
// 	if status != http.StatusOK {
// 		t.Error("handler returned incorrect status code, should be 200")
// 		fmt.Println(httpTest.Body)
// 	}
// 	fmt.Println("signed the user in")
// 	t.Error("fail because we're new")

// 	// update request
// 	userRemoveParams := mutations.RemoveUserRequestParams{
// 		Email: "brian.t.vann2@gmail.com",
// 		Password: "passw3rd",
// 	}

// 	requestBodyRemove := mutations.RemoveUserRequestBody{
// 		Action: "REMOVE_USER",
// 		Params: userRemoveParams,
// 	}
// 	marshalBytesUpdate := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytesUpdate).Encode(requestBodyRemove)

// 	reqRemove, errReqRemove := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytesUpdate,
// 	)
// 	if errReqRemove != nil {
// 		t.Error("error making guest session request")
// 	}
// 	reqRemove.Header.Set(
// 		constants.SessionTokenHeader,
// 		httpTest.Header().Get(constants.SessionTokenHeader),
// 	)

// 	reqRemove.Header.Set(
// 		constants.CsrfTokenHeader,
// 		httpTest.Header().Get(constants.CsrfTokenHeader),
// 	)

// 	httpTestUpdate := httptest.NewRecorder()
// 	// store mutation
// 	handlerUpdate := http.HandlerFunc(Mutation)
// 	handlerUpdate.ServeHTTP(httpTestUpdate, reqRemove)

// 	statusUpdate := httpTestUpdate.Code
// 	if statusUpdate != http.StatusOK {
// 		t.Error("handler returned incorrect status code, should be 200")
// 		fmt.Println(httpTestUpdate.Body)
// 	}
// }

// // ReviveUser
// func TestReviveUser(t *testing.T) {
// 	// get gets headers
// 	requestBodyGuest := RequestBodyParams{
// 		Action: sessions.CreateGuestSession,
// 	}

// 	marshalBytesGuest := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytesGuest).Encode(requestBodyGuest)
// 	resp, errResp := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytesGuest,
// 	)
// 	if errResp != nil {
// 		t.Error("error making guest session request")
// 	}

// 	httpTestGuest := httptest.NewRecorder()
// 	handlerGuest := http.HandlerFunc(sessions.Mutation)
// 	handlerGuest.ServeHTTP(httpTestGuest, resp)

// 	statusGuest := httpTestGuest.Code
// 	if statusGuest != http.StatusOK {
// 		t.Error("handler returned incorrect status code")
// 	}

// 	userParams := store.CreateUserParams{
// 		Email:    "brian.t.vann2@gmail.com",
// 		Password: "passw3rd",
// 	}

// 	// Sign user into public Sessions
// 	requestBody := mutations.CreateUserRequestBodyParams{
// 		Action: sessions.CreatePublicSession,
// 		Params: userParams,
// 	}
// 	marshalBytes := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytes).Encode(requestBody)

// 	req, errReq := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytes,
// 	)
// 	if errReq != nil {
// 		t.Error("error making guest session request")
// 	}
// 	req.Header.Set(
// 		constants.SessionTokenHeader,
// 		httpTestGuest.Header().Get(constants.SessionTokenHeader),
// 	)

// 	req.Header.Set(
// 		constants.CsrfTokenHeader,
// 		httpTestGuest.Header().Get(constants.CsrfTokenHeader),
// 	)

// 	httpTest := httptest.NewRecorder()
// 	handler := http.HandlerFunc(sessions.Mutation)
// 	handler.ServeHTTP(httpTest, req)

// 	status := httpTest.Code
// 	if status != http.StatusOK {
// 		t.Error("handler returned incorrect status code, should be 200")
// 		fmt.Println(httpTest.Body)
// 	}
// 	fmt.Println("signed the user in")
// 	t.Error("fail because we're new")

// 	// update request
// 	userReviveParams := mutations.RemoveUserRequestParams{
// 		Email: "brian.t.vann2@gmail.com",
// 		Password: "passw3rd",
// 	}

// 	requestBodyRemove := mutations.RemoveUserRequestBody{
// 		Action: "REVIVE_USER",
// 		Params: userReviveParams,
// 	}
// 	marshalBytesUpdate := new(bytes.Buffer)
// 	json.NewEncoder(marshalBytesUpdate).Encode(requestBodyRemove)

// 	reqRemove, errReqRemove := http.NewRequest(
// 		"POST",
// 		"/sessions/m/",
// 		marshalBytesUpdate,
// 	)
// 	if errReqRemove != nil {
// 		t.Error("error making guest session request")
// 	}
// 	reqRemove.Header.Set(
// 		constants.SessionTokenHeader,
// 		httpTest.Header().Get(constants.SessionTokenHeader),
// 	)

// 	reqRemove.Header.Set(
// 		constants.CsrfTokenHeader,
// 		httpTest.Header().Get(constants.CsrfTokenHeader),
// 	)

// 	httpTestUpdate := httptest.NewRecorder()
// 	// store mutation
// 	handlerUpdate := http.HandlerFunc(Mutation)
// 	handlerUpdate.ServeHTTP(httpTestUpdate, reqRemove)

// 	statusUpdate := httpTestUpdate.Code
// 	if statusUpdate != http.StatusOK {
// 		t.Error("handler returned incorrect status code, should be 200")
// 		fmt.Println(httpTestUpdate.Body)
// 	}
// }
