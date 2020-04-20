package sessions

import (
	"testing"
	"webapi/interfaces/jwtx"
	"webapi/store"
)

const unitTests = "UNIT_TESTS"
var fakeSession = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJ5b2Rhd2ciLCJpYXQiOjE1ODIzMTU0NDYsImV4cCI6MTYxMzg1MTQ0NiwiYXVkIjoid3d3LmJsYWhibGFoLmNvbSIsInN1YiI6ImpvaG5ueUBxdWVzdC5jb20iLCJHaXZlbk5hbWUiOiJKb2hubnkiLCJTdXJuYW1lIjoiUXVlc3QiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20ifQ.hPmDUos2HzRgn-OfFcC3gzhi28xa5YEDAVwfxWYfvdY"
var storeSuccess, errStore = store.CreateRequiredTables()

func TestCreateGuestSession(t *testing.T) {
	session, errSession := Create(&CreateParams{
		Environment: unitTests,
		Claims: *CreateGuestSessionClaims(),
	})
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error(errSession.Error())
	}
}

func TestRetrieveGuestSession(t *testing.T) {
	session, errSession := Create(&CreateParams{
		Environment: unitTests,
		Claims: *CreateGuestSessionClaims(),
	})

	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error(errSession.Error())
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		session.SessionToken,
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error("error interpreting token")
	}

	// update the token
	updatedSession, errUpdatedSession := Update(
		&UpdateParams{
			Environment:  unitTests,
			SessionToken: session.SessionToken,
		},
	)
	if errUpdatedSession != nil {
		t.Error(errUpdatedSession.Error())
	}
	if updatedSession == nil {
		t.Error("nil value returned instead of session")
		return
	}

	updatedTokenDetails, errUpdatedTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		updatedSession.SessionToken,
	)
	if updatedTokenDetails == nil {
		t.Error("token is nil")
	}
	if errUpdatedTokenDetails != nil {
		t.Error(errUpdatedTokenDetails.Error())
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

func TestUpdateSession(t *testing.T) {
	session, errSession := Create(&CreateParams{
		Environment: unitTests,
		Claims: *CreateGuestSessionClaims(),
	})
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error(errSession.Error())
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		session.SessionToken,
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error(errTokenDetails.Error())
	}

	// update the token
	ReadSession, errReadSession := Update(
		&UpdateParams{
			Environment: unitTests,
			SessionToken: session.SessionToken,
		},
	)
	if errReadSession != nil {
		t.Error(errReadSession.Error())
	}
	if ReadSession == nil {
		t.Error("nil value returned instead of session")
		return
	}
}

func TestValidateAndRemoveSession(t *testing.T) {
	session, errSession := Create(&CreateParams{
		Environment: unitTests,
		Claims: *CreateGuestSessionClaims(),
	})
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error("Error creating public JWT")
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		session.SessionToken,
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error(errTokenDetails.Error())
	}

	// update the token
	removedSession, errRemovedSession := ValidateAndRemove(
		&UpdateParams{
			Environment: unitTests,
			SessionToken: session.SessionToken,
		},
	)
	if errRemovedSession != nil {
		t.Error(errRemovedSession.Error())
	}
	if removedSession == nil {
		t.Error("could not find session")
		return
	}

	// update the token
	reReadSession, errReReadSession := Read(
		&ReadParams{
			Environment: unitTests,
			SessionToken: session.SessionToken,
		},
	)
	if errReReadSession != nil {
		t.Error(errReReadSession.Error())
	}
	if reReadSession != false {
		t.Error("should not have found session")
		return
	}
}

func TestCheckBadSession(t *testing.T) {
	signature := fakeSession
	readSession, errReadSession := Update(
		&UpdateParams{
			Environment: unitTests,
			SessionToken: signature,
		},
	)
	if errReadSession != nil {
		t.Error(errReadSession.Error())
	}
	if readSession != nil {
		t.Error("value returned instead of nil")
	}
}

// Test Create Public JWT
func TestRetrieveUserSession(t *testing.T) {
	sessionClaims, errSessionClaims := CreateUserSessionClaims(
		&CreateUserClaimsParams{
			UserID: -1,
		},
	)
	if errSessionClaims != nil {
		t.Error(errSessionClaims.Error())
		return
	}
	if sessionClaims == nil {
		t.Error("nil session claims")
		return
	}

	// create session
	session, errSession := Create(&CreateParams{
		Environment: unitTests,
		Claims: *sessionClaims,
	})
	if session == nil {
		t.Error("Nil value returned")
	}
	if errSession != nil {
		t.Error(errSession.Error())
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		session.SessionToken,
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error(errTokenDetails.Error())
	}

	// update the token
	updatedSession, errUpdatedSession := Update(
		&UpdateParams{
			Environment: unitTests,
			SessionToken: session.SessionToken,
		},
	)
	if errUpdatedSession != nil {
		t.Error(errUpdatedSession.Error())
	}
	if updatedSession == nil {
		t.Error("nil value returned instead of session")
	}

	// exit test if updated session is nil
	if updatedSession == nil {
		return
	}

	updatedTokenDetails, errUpdatedTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		updatedSession.SessionToken,
	)
	if updatedTokenDetails == nil {
		t.Error("toke is nil")
	}
	if errUpdatedTokenDetails != nil {
		t.Error(errUpdatedTokenDetails.Error())
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
	session, errSession := Create(&CreateParams{
		Environment: unitTests,
		Claims: *CreateGuestSessionClaims(),
	})
	if session == nil {
		t.Error("Nil value returned")

	}
	if errSession != nil {
		t.Error(errSession.Error())
	}

	// get the session token
	tokenDetails, errTokenDetails := jwtx.RetrieveTokenDetailsFromString(
		session.SessionToken,
	)
	if tokenDetails == nil {
		t.Error("token is nil")
	}
	if errTokenDetails != nil {
		t.Error(errTokenDetails.Error())
	}

	entryRemoved, errEntryRemoved := Remove(&RemoveParams{
		Environment: unitTests,
		Signature: tokenDetails.Signature,
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
	entryRemoved, errEntryRemoved := Remove(&RemoveParams{
		Environment: unitTests,
		Signature: badSignature,
	})
	if errEntryRemoved != nil {
		t.Error(errEntryRemoved.Error())
	}

	if entryRemoved == true {
		t.Error("failed to remove nonexistent session")
	}
}
