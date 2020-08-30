package sessionsx

import (
	// "os"
	"testing"
	"webapi/store"

	"webapi/utils/jwtx"
)

const unitTests = "UNIT_TESTS"

var fakeSession = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJ5b2Rhd2ciLCJpYXQiOjE1ODIzMTU0NDYsImV4cCI6MTYxMzg1MTQ0NiwiYXVkIjoid3d3LmJsYWhibGFoLmNvbSIsInN1YiI6ImpvaG5ueUBxdWVzdC5jb20iLCJHaXZlbk5hbWUiOiJKb2hubnkiLCJTdXJuYW1lIjoiUXVlc3QiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20ifQ.hPmDUos2HzRgn-OfFcC3gzhi28xa5YEDAVwfxWYfvdY"
var storeSuccess, errStore = store.CreateRequiredTables()

var GuestSession *Session

func TestCreateGuestSession(t *testing.T) {
	session, errSession := Create(&CreateParams{
		Environment: unitTests,
		Claims:      CreateGuestSessionClaims(),
	})
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error(errSession.Error())
	}

	GuestSession = session
}

func TestReadSession(t *testing.T) {
	result, errResult := Read(&ReadParams{
		Environment: unitTests,
		Token:       GuestSession.Token,
	})
	if errResult != nil {
		t.Error(errResult)
	}
	if result == false {
		t.Error("bad values")
	}
}

func TestRemoveSession(t *testing.T) {
	// create public session
	session, errSession := Create(&CreateParams{
		Environment: unitTests,
		Claims:      CreateGuestSessionClaims(),
	})
	if session == nil {
		t.Error("Nil value returned")

	}
	if errSession != nil {
		t.Error(errSession.Error())
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		session.Token,
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error(errTokenDetails.Error())
	}

	entryRemoved, errEntryRemoved := Delete(&DeleteParams{
		Environment: unitTests,
		Signature:   tokenDetails.Signature,
	})
	if errEntryRemoved != nil {
		t.Error(errEntryRemoved.Error())
	}

	if entryRemoved == false {
		t.Error("failed to remove session")
	}
}

func TestRemoveSessionRespondsFalse(t *testing.T) {
	badSignature := "animal_crackers_with_nutella"
	entryRemoved, errEntryRemoved := Delete(&DeleteParams{
		Environment: unitTests,
		Signature:   badSignature,
	})
	if errEntryRemoved != nil {
		t.Error(errEntryRemoved.Error())
	}

	if entryRemoved == true {
		t.Error("failed to remove nonexistent session")
	}
}
