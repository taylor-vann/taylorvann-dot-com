// brian taylor vann
// taylorvann-dot-com

package hooks

import (
	"bytes"
	"encoding/json"
	"net/http"
	// "net/http/resp"
	"testing"
	"time"

	"toolbox/clientx"

	"webapi/store/roles/controller"	
	"webapi/store/roles/hooks/requests"
	"webapi/store/roles/hooks/responses"	
)

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
	UserID: -1,
	Organization: "STORE_ROLES_UNIT_TESTS",
	ReadAccess: false,
	WriteAccess: false,
}

var user1Updated = requests.Update{
	Environment: "LOCAL",
	UserID: -1,
	Organization: "STORE_ROLES_UNIT_TESTS",
	ReadAccess: true,
	WriteAccess: true,
	IsDeleted: false,
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
	resp, errResp := clientx.Do(
		"https://authn.briantaylorvann.com/m/roles/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}
	if resp == nil {
		t.Error("resp is nil")
		return
	}

	if resp.Body == nil {
		t.Error("response body is nil")
		return
	}
	
	var responseBody responses.Body
	errJSON := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Roles == nil {
		t.Error("nil roles returned")
	}

	if resp.StatusCode != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestRead(t *testing.T) {
	requestBody := requests.Body{
		Action: Read,
		Params: requests.Read{
			Environment: "LOCAL",
			UserID: user1.UserID,
			Organization: user1.Organization,
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := clientx.Do(
		"https://authn.briantaylorvann.com/q/roles/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}
	if resp == nil {
		t.Error("resp is nil")
		return
	}

	if resp.Body == nil {
		t.Error("response body is nil")
		return
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Roles == nil {
		t.Error("nil roles returned")
		return
	}

	if len(*responseBody.Roles) == 0 {
		t.Error("zero roles returned")
		return
	}

	if resp.StatusCode != http.StatusOK {
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
	resp, errResp := clientx.Do(
		"https://authn.briantaylorvann.com/q/roles/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}
	if resp == nil {
		t.Error("resp is nil")
		return
	}

	if resp.Body == nil {
		t.Error("response body is nil")
		return
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Roles == nil {
		t.Error("nil roles returned")
		return
	}

	if len(*responseBody.Roles) == 0 {
		t.Error("zero roles returned")
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestSearch(t *testing.T) {
	requestBody := requests.Body{
		Action: Search,
		Params: requests.Search{
			Environment: "LOCAL",
			UserID: user1.UserID,
			StartIndex: 0,
			Length: 10,
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := clientx.Do(
		"https://authn.briantaylorvann.com/q/roles/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}
	if resp == nil {
		t.Error("resp is nil")
		return
	}

	if resp.Body == nil {
		t.Error("response body is nil")
		return
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Roles == nil {
		t.Error("nil roles returned")
		return
	}

	if len(*responseBody.Roles) == 0 {
		t.Error("zero roles returned")
		return
	}

	if resp.StatusCode != http.StatusOK {
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
	resp, errResp := clientx.Do(
		"https://authn.briantaylorvann.com/m/roles/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}
	if resp == nil {
		t.Error("resp is nil")
		return
	}

	if resp.Body == nil {
		t.Error("response body is nil")
		return
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Roles == nil {
		t.Error("nil roles returned")
		return
	}

	if len(*responseBody.Roles) == 0 {
		t.Error("zero roles returned")
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestUpdateAccess(t *testing.T) {
	requestBody := requests.Body{
		Action: UpdateAccess,
		Params: requests.UpdateAccess{
			Environment: "LOCAL",
			UserID: -1,
			Organization: "STORE_ROLES_UNIT_TESTS",
			ReadAccess: false,
			WriteAccess: false,
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := clientx.Do(
		"https://authn.briantaylorvann.com/m/roles/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}
	if resp == nil {
		t.Error("resp is nil")
		return
	}

	if resp.Body == nil {
		t.Error("response body is nil")
		return
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Roles == nil {
		t.Error("nil roles returned")
		return
	}

	if len(*responseBody.Roles) == 0 {
		t.Error("zero roles returned")
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}

func TestDelete(t *testing.T) {
	requestBody := requests.Body{
		Action: Delete,
		Params: requests.Delete{
			Environment: "LOCAL",
			UserID: -1,
			Organization: "STORE_ROLES_UNIT_TESTS",
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := clientx.Do(
		"https://authn.briantaylorvann.com/m/roles/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}
	if resp == nil {
		t.Error("resp is nil")
		return
	}

	if resp.Body == nil {
		t.Error("response body is nil")
		return
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Roles == nil {
		t.Error("nil roles returned")
		return
	}

	if len(*responseBody.Roles) == 0 {
		t.Error("zero roles returned")
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Error(*responseBody.Errors.Default)
	}
}


func TestUndelete(t *testing.T) {
	requestBody := requests.Body{
		Action: Undelete,
		Params: requests.Undelete{
			Environment: "LOCAL",
			UserID: -1,
			Organization: "STORE_ROLES_UNIT_TESTS",
		},
	}

	marshalBytes := new(bytes.Buffer)
	json.NewEncoder(marshalBytes).Encode(requestBody)
	resp, errResp := clientx.Do(
		"https://authn.briantaylorvann.com/m/roles/",
		marshalBytes,
	)
	if errResp != nil {
		t.Error(errResp.Error())
	}
	if resp == nil {
		t.Error("resp is nil")
		return
	}

	if resp.Body == nil {
		t.Error("response body is nil")
		return
	}
	var responseBody responses.Body
	errJSON := json.NewDecoder(resp.Body).Decode(&responseBody)
	if errJSON != nil {
		t.Error(errJSON.Error())
	}
	if responseBody.Roles == nil {
		t.Error("nil roles returned")
		return
	}

	if len(*responseBody.Roles) == 0 {
		t.Error("zero roles returned")
		return
	}

	if resp.StatusCode != http.StatusOK {
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