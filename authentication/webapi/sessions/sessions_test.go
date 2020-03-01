package sessions

import (
	"fmt"
	"testing"
	"webapi/interfaces/jwtx"
	"webapi/store"
)

var storeSuccess, errStore = store.CreateRequiredTables()

func TestCreateGuestSession(t *testing.T) {
	session, errSession := Create(
		ComposeCreateGuestSessionParams(),
	)
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error("Error creating public JWT")
	}
}

func TestRetrieveGuestSession(t *testing.T) {
	session, errSession := Create(
		ComposeCreateGuestSessionParams(),
	)
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error("Error creating public JWT")
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		&(session.SessionToken),
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error("error interpreting token")
	}

	// update the token
	updatedSession, errUpdatedSession := UpdateIfExists(
		&UpdateParams{
			SessionToken: &(session.SessionToken),
			CsrfToken:    &(session.CsrfToken),
		},
	)
	if errUpdatedSession != nil {
		t.Error("error updated session")
	}
	if updatedSession == nil {
		t.Error("nil value returned instead of session")
		return
	}

	updatedTokenDetails, errUpdatedTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		&(updatedSession.SessionToken),
	)
	if updatedTokenDetails == nil {
		t.Error("token is nil")
	}
	if errUpdatedTokenDetails != nil {
		t.Error("error interpreting token")
	}

	if tokenDetails.Payload.Iss != updatedTokenDetails.Payload.Iss {
		t.Error("mismatching issuer")
	}
	if tokenDetails.Payload.Sub != updatedTokenDetails.Payload.Sub {
		t.Error("mismatching subject")
	}
	if tokenDetails.Payload.Aud != updatedTokenDetails.Payload.Aud {
		t.Error("mismatching audience")
	}
}

func TestCheckSession(t *testing.T) {
	session, errSession := Create(
		ComposeCreateGuestSessionParams(),
	)
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error("Error creating public JWT")
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		&(session.SessionToken),
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error("error interpreting token")
	}

	// update the token
	ReadSession, errReadSession := ReadIfExists(
		&ReadParams{
			SessionToken: &(session.SessionToken),
		},
	)
	if errReadSession != nil {
		t.Error("error updated session")
	}
	if ReadSession == nil {
		t.Error("nil value returned instead of session")
		return
	}

	readTokenDetails, errReadTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		ReadSession.SessionToken,
	)
	if readTokenDetails == nil {
		t.Error("token is nil")
	}
	if errReadTokenDetails != nil {
		t.Error("error interpreting token")
	}

	if tokenDetails.Payload.Iss != readTokenDetails.Payload.Iss {
		t.Error("mismatching issuer")
	}
	if tokenDetails.Payload.Sub != readTokenDetails.Payload.Sub {
		t.Error("mismatching subject")
	}
	if tokenDetails.Payload.Aud != readTokenDetails.Payload.Aud {
		t.Error("mismatching audience")
	}
}

func TestDoubleCheckSession(t *testing.T) {
	session, errSession := Create(
		ComposeCreateGuestSessionParams(),
	)
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error("Error creating public JWT")
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		&(session.SessionToken),
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error("error interpreting token")
	}

	// update the token
	readSession, errReadSession := ReadIfExists(
		&ReadParams{
			SessionToken: &(session.SessionToken),
		},
	)
	if errReadSession != nil {
		t.Error("error updated session")
	}
	if readSession == nil {
		t.Error("nil value returned instead of session")
		return
	}

	// update the token
	reReadSession, errReReadSession := ReadIfExists(
		&ReadParams{
			SessionToken: &(session.SessionToken),
		},
	)
	if errReReadSession != nil {
		t.Error("error updated session")
	}
	if reReadSession == nil {
		t.Error("nil value returned instead of session")
		return
	}
}

func TestCheckBadSession(t *testing.T) {
	signature := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJ5b2Rhd2ciLCJpYXQiOjE1ODIzMTU0NDYsImV4cCI6MTYxMzg1MTQ0NiwiYXVkIjoid3d3LmJsYWhibGFoLmNvbSIsInN1YiI6ImpvaG5ueUBxdWVzdC5jb20iLCJHaXZlbk5hbWUiOiJKb2hubnkiLCJTdXJuYW1lIjoiUXVlc3QiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20ifQ.hPmDUos2HzRgn-OfFcC3gzhi28xa5YEDAVwfxWYfvdY"
	readSession, errReadSession := ReadIfExists(
		&ReadParams{
			SessionToken: &signature,
		},
	)
	if errReadSession != nil {
		fmt.Println(errReadSession)
		t.Error("error check bad session")
	}
	if readSession != nil {
		t.Error("value returned instead of nil")
	}
}

// Test Create Public JWT
func TestRetrievePublicSession(t *testing.T) {
	userParams := store.CreateUserParams{
		Email:    "testorino@booyakasha.com",
		Password: "pesswerd",
	}
	// create user
	_, errUserRow := store.CreateUser(&userParams)
	if errUserRow != nil {
		t.Error("error creating user row")
	}

	sessionParams, errSessionParams := ComposeCreatePublicSessionParams(
		&CreatePublicJWTParams{
			Email:    &userParams.Email,
			Password: &userParams.Password,
		},
	)
	if errSessionParams != nil {
		fmt.Println(errSessionParams)
		t.Error("error creating session params")
		return
	}
	// create session
	session, errSession := Create(sessionParams)
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error("Error creating public JWT")
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		&(session.SessionToken),
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error("error interpreting token")
	}

	// update the token
	updatedSession, errUpdatedSession := UpdateIfExists(
		&UpdateParams{
			SessionToken: &(session.SessionToken),
			CsrfToken:    &(session.CsrfToken),
		},
	)
	if errUpdatedSession != nil {
		t.Error("error updated session")
	}
	if updatedSession == nil {
		t.Error("nil value returned instead of session")
	}

	// exit test if updated session is nil
	if updatedSession == nil {
		return
	}

	updatedTokenDetails, errUpdatedTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		&(updatedSession.SessionToken),
	)
	if updatedTokenDetails == nil {
		t.Error("toke is nil")
	}
	if errUpdatedTokenDetails != nil {
		t.Error("error interpreting token")
	}

	if tokenDetails.Payload.Iss != updatedTokenDetails.Payload.Iss {
		t.Error("mismatching issuer")
	}
	if tokenDetails.Payload.Sub != updatedTokenDetails.Payload.Sub {
		t.Error("mismatching subject")
	}
	if tokenDetails.Payload.Aud != updatedTokenDetails.Payload.Aud {
		t.Error("mismatching audience")
	}
}

func TestRemoveSession(t *testing.T) {
	// create public session
	session, errSession := Create(
		ComposeCreateGuestSessionParams(),
	)
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error("Error creating public JWT")
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		&(session.SessionToken),
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error("error interpreting token")
	}

	entryRemoved, errEntryRemoved := RemoveSession(&RemoveParams{
		Signature: tokenDetails.Signature,
	})
	if errEntryRemoved != nil {
		t.Error("error removing session")
	}

	if entryRemoved == false {
		t.Error("failed to remove session")
	}
}

func TestRemoveSessionRespondsFalse(t *testing.T) {
	badSignature := "animal_crackers_with_nutella"
	entryRemoved, errEntryRemoved := RemoveSession(&RemoveParams{
		Signature: &badSignature,
	})
	if errEntryRemoved != nil {
		t.Error("error removing session")
	}

	if entryRemoved == true {
		t.Error("failed to remove nonexistent session")
	}
}