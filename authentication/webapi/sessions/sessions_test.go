package sessions

import (
	"testing"
	"webapi/interfaces/jwtx"
	"webapi/store"
)

var storeSuccess, errStore = store.CreateRequiredTables()

func TestCreatePublicSession(t *testing.T) {
	session, errSession := CreateSession(
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
	session, errSession := CreateSession(
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
	updatedSession, errUpdatedSession := UpdateSessionIfExists(
		&UpdateSessionParams{
			SessionToken: &(session.SessionToken),
			CsrfToken: &(session.CsrfToken),
		},
	)
	if errUpdatedSession != nil {
		t.Error("error updated session")
	}
	if updatedSession == nil {
		t.Error("nil value returned instead of session")
	}
	
	if updatedSession != nil {
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
}

// Test Create Public JWT
func TestRetrievePublicSession(t *testing.T) {
	userParams := store.CreateUserParams{
		Email: "testorino@booyakasha.com",
		Password: "pesswerd",
	}
	// create user
	store.CreateUser(&userParams)

	sessionParams, errSessionParams := ComposeCreatePublicSessionParams(&CreatePublicJWTParams{
		Email: &userParams.Email,
		Password: &userParams.Password,
	})
	if errSessionParams != nil {
		t.Error("error creating session params")
		return
	}
	// create session
	session, errSession := CreateSession(sessionParams)
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
	updatedSession, errUpdatedSession := UpdateSessionIfExists(
		&UpdateSessionParams{
			SessionToken: &(session.SessionToken),
			CsrfToken: &(session.CsrfToken),
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
	session, errSession := CreateSession(
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

	entryRemoved, errEntryRemoved := RemoveSession(&RemoveSessionParams{
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
	entryRemoved, errEntryRemoved := RemoveSession(&RemoveSessionParams{
		Signature: &badSignature,
	})
	if errEntryRemoved != nil {
		t.Error("error removing session")
	}

	if entryRemoved == true {
		t.Error("failed to remove nonexistent session")
	}
}
