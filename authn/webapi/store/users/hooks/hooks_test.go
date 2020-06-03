// brian taylor vann
// taylorvann-dot-com

package hooks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"

	"webapi/store/users/controller"	
	"webapi/store/users/hooks/requests"
	"webapi/store/users/hooks/responses"	
	"webapi/store/validatesessionx"
)

var httpTestClient = getTestClient()

type Row struct {
	ID					 int64     `json:"id"`
	UserID    	 int64     `json:"user_id"`
	Organization string    `json:"organization"`
	ReadAccess	 bool			 `json:"read_access"`
	WriteAccess	 bool			 `json:"write_access"`
	IsDeleted		 bool			 `json:"is_deleted"`
	CreatedAt		 time.Time `json:"created_at"`
	UpdatedAt		 time.Time `json:"updated_at"`
}

var createTable = controller.CreateTableParams{
	Environment: "LOCAL",
}

var user1 = requests.Create{
	Environment: "LOCAL",
	Email: "test_user_willhelm_dallas_truday@unit_tests.com",
	Password: "Pazzwerd",
}

var user1Search = requests.Search{
	Environment: "LOCAL",
	EmailSubstring: "willhelm_dallas_truday",
	StartIndex: 0,
	Length: 10,
}

var user1Updated = requests.Update{
	Environment: "LOCAL",
	CurrentEmail: "test_user_willhelm_dallas_truday@unit_tests.com",
	UpdatedEmail: "test_user_willhelm_dallas_thursday@unit_tests.com",
	Password: "Pazzw3rd",
	IsDeleted: false,
}

var user1UpdatedEmail = requests.UpdateEmail{
	Environment: "LOCAL",
	CurrentEmail: "test_user_willhelm_dallas_thursday@unit_tests.com",
	UpdatedEmail: "test_user_willhelm_dallas_thursdays@unit_tests.com",
}

var user1UpdatedPassword = requests.UpdatePassword{
	Environment: "LOCAL",
	Email: "test_user_willhelm_dallas_thursdays@unit_tests.com",
	Password: "PAZZw3rd",
}

// var user1Updated = requests.Update{
// 	Environment: "LOCAL",
// 	Email: "willhelm_dallas_truday@unit_tests.com",
// 	Email: "updatedtest_user_willhelm_dallas_truday@unit_tests.com",
// 	Password: "Pazzwerd",
// }

func getTestClient() *http.Client {
	jar, errJar := cookiejar.New(nil)
	if errJar != nil {
		return nil
	}
	return &http.Client{
		Jar: jar,
	}
}

func TestCreateTable(t *testing.T) {
	results, err := controller.CreateTable(&createTable)
	if err != nil {
		t.Error(err.Error())
	}
	if results == nil {
		t.Error("no results were returned from CreateTable.")
	}
}

func TestCreate(t *testing.T) {
	requestBody := requests.Body{
		Action: Create,
		Params: user1,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/q/users/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Users == nil {
		t.Error("nil users returned")
	}

	if httpTest.Code != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestValidateGuest(t *testing.T) {
	guestSession, errGuestSession := validatesessionx.FetchGuestSession()
	if errGuestSession != nil {
		t.Error("couldn't get guest session")
		return
	}
	requestBody := requests.Body{
		Action: ValidateGuest,
		Params: user1,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	req := httptest.NewRequest(
		"POST",
		"/q/users/",
		marshalBytes,
	)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)

	http.SetCookie(
		recorder,
		&http.Cookie{
			Name: "briantaylorvann.com_internal_session",
			Value: guestSession,
			MaxAge:		10000,
			Domain:   "briantaylorvann.com",
			Path:     "/",
			Secure:		true,
			HttpOnly:	true,
			SameSite:	3,
		},
	)

	handler.ServeHTTP(recorder, req)
	

	if req.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(recorder.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Users == nil {
		t.Error("nil users returned")
	}

	if recorder.Code != http.StatusOK {
		t.Error(*responseBody.Errors)
	}
}

func TestRead(t *testing.T) {
	requestBody := requests.Body{
		Action: Read,
		Params: requests.Read{
			Environment: "LOCAL",
			Email: user1.Email,
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/q/users/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Users == nil {
		t.Error("nil users returned")
		return
	}

	if len(*responseBody.Users) == 0 {
		t.Error("zero users returned")
		return
	}

	if httpTest.Code != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestIndex(t *testing.T) {
	requestBody := requests.Body{
		Action: Index,
		Params: requests.Index{
			Environment: "LOCAL",
			StartIndex: 0,
			Length: 10,
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/q/users/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Users == nil {
		t.Error("nil users returned")
		return
	}

	if len(*responseBody.Users) == 0 {
		t.Error("zero users returned")
		return
	}

	if httpTest.Code != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestSearch(t *testing.T) {
	requestBody := requests.Body{
		Action: Search,
		Params: user1Search,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/q/users/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Query)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Users == nil {
		t.Error("nil users returned")
		return
	}

	if len(*responseBody.Users) == 0 {
		t.Error("zero users returned")
		return
	}

	if httpTest.Code != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestUpdate(t *testing.T) {
	requestBody := requests.Body{
		Action: Update,
		Params: user1Updated,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/users/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Users == nil {
		t.Error("nil users returned")
		return
	}

	if len(*responseBody.Users) == 0 {
		t.Error("zero users returned")
		return
	}

	if httpTest.Code != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestUpdateEmail(t *testing.T) {
	requestBody := requests.Body{
		Action: UpdateEmail,
		Params: user1UpdatedEmail,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/users/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Users == nil {
		t.Error("nil users returned")
		return
	}

	if len(*responseBody.Users) == 0 {
		t.Error("zero users returned")
		return
	}

	if httpTest.Code != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestUpdatePassword(t *testing.T) {
	requestBody := requests.Body{
		Action: UpdatePassword,
		Params: user1UpdatedPassword,
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/users/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Users == nil {
		t.Error("nil users returned")
		return
	}

	if len(*responseBody.Users) == 0 {
		t.Error("zero users returned")
		return
	}

	if httpTest.Code != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestDelete(t *testing.T) {
	requestBody := requests.Body{
		Action: Delete,
		Params: requests.Delete{
			Environment: "LOCAL",
			Email: user1UpdatedPassword.Email,
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/users/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Users == nil {
		t.Error("nil users returned")
		return
	}

	if len(*responseBody.Users) == 0 {
		t.Error("zero users returned")
		return
	}

	if httpTest.Code != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestUndelete(t *testing.T) {
	requestBody := requests.Body{
		Action: Undelete,
		Params: requests.Undelete{
			Environment: "LOCAL",
			Email: user1UpdatedPassword.Email,
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := http.NewRequest(
		"POST",
		"/m/users/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}

	httpTest := httptest.NewRecorder()
	handler := http.HandlerFunc(Mutation)
	handler.ServeHTTP(httpTest, resp)

	if resp.Body == nil {
		t.Error("response body is nil")
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(httpTest.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Users == nil {
		t.Error("nil users returned")
		return
	}

	if len(*responseBody.Users) == 0 {
		t.Error("zero users returned")
		return
	}

	if httpTest.Code != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestDangerouslyDropUnitTestsTable(t *testing.T) {
	result, err := controller.DangerouslyDropUnitTestsTable()
	if result == nil {
		t.Error("Failed to drop table")
	}
	if err != nil {
		t.Error(err.Error())
	}
}