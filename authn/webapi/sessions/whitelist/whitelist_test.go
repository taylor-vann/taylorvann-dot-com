package whitelist

import (
	"strconv"
	"testing"
	"time"

	"webapi/constants"
	"webapi/interfaces/jwtx"
)

type JWTClaimTestPlan = []jwtx.Claims

const TestEnvironment = "UNIT_TESTS"
var randomJWTClaims = generateRandomJWTClaims("public", 5)

func getNowAsMS() MilliSeconds {
	return time.Now().UnixNano() / int64(time.Millisecond)
}


func getLaterAsMS() MilliSeconds {
	return (time.Now().UnixNano() + DayAsMS) / int64(time.Millisecond)
}

func generateRandomJWTClaims(sub string, num int) *JWTClaimTestPlan {
	jwtClaims := make(JWTClaimTestPlan, num)

	for index := range jwtClaims {
		nowAsMS := getNowAsMS()
		laterAsMS := getLaterAsMS()

		jwtClaims[index] = jwtx.Claims{
			Iss: constants.TaylorVannDotCom,
			Sub: sub,
			Aud: strconv.Itoa(index),
			Iat: nowAsMS,
			Exp: laterAsMS,
		}
	}

	return &jwtClaims
}

func TestCreateEntry(t *testing.T) {
	var tokens = make([]*jwtx.TokenPayload, len(*randomJWTClaims))

	for index, claim := range *randomJWTClaims {
		token, errToken := jwtx.CreateJWT(&claim)
		if errToken != nil {
			t.Error("Unable to create jwt")
		}
		tokens[index] = token
		entry, errEntry := CreateEntry(&CreateEntryParams{
			Environment: TestEnvironment,
			CreatedAt:  claim.Iat,
			Lifetime:   DayAsMS,
			SessionKey: *token.RandomSecret,
			Signature:  token.Token.Signature,
		})

		if errEntry != nil {
			t.Error(errEntry.Error())
		}

		if entry == nil {
			t.Error("nil entry returned")
		}
	}
}

func TestReadEntry(t *testing.T) {
	var tokens = make([]*jwtx.TokenPayload, len(*randomJWTClaims))

	for index, claim := range *randomJWTClaims {
		token, errToken := jwtx.CreateJWT(&claim)
		if errToken != nil {
			t.Error(errToken.Error())
		}
		tokens[index] = token
		entry, errEntry := CreateEntry(&CreateEntryParams{
			Environment: TestEnvironment,
			CreatedAt:  claim.Iat,
			Lifetime:   DayAsMS,
			SessionKey: *token.RandomSecret,
			Signature:  token.Token.Signature,
		})

		if errEntry != nil {
			t.Error(errEntry.Error())
		}

		if entry == nil {
			t.Error("nil entry returned")
		}
	}

	// check entries for accuracy
	for _, token := range tokens {
		readEntry, errReadEntry := ReadEntry(&ReadEntryParams{
			Environment: TestEnvironment,
			Signature: token.Token.Signature,
		})
		if errReadEntry != nil {
			t.Error(errReadEntry.Error())
		}

		randomSecret := *(token.RandomSecret)
		if len(randomSecret) != len(readEntry.SessionKey) {
			t.Error("mismatching secret key lengths")
		}
		for index, bit := range randomSecret {
			if bit != readEntry.SessionKey[index] {
				t.Error("mismatching secret keys")
			}
			break
		}
	}
}

func TestRemoveEntry(t *testing.T) {
	var tokens = make([]*jwtx.TokenPayload, len(*randomJWTClaims))

	for index, claim := range *randomJWTClaims {
		token, errToken := jwtx.CreateJWT(&claim)
		if errToken != nil {
			t.Error("Unable to create jwt")
		}
		tokens[index] = token
		entry, errEntry := CreateEntry(&CreateEntryParams{
			Environment: TestEnvironment,
			CreatedAt:  claim.Iat,
			Lifetime:   DayAsMS,
			SessionKey: *token.RandomSecret,
			Signature:  token.Token.Signature,
		})

		if errEntry != nil {
			t.Error(errEntry.Error())
		}

		if entry == nil {
			t.Error("nil entry returned")
		}
	}

	// check entries for accuracy
	for _, token := range tokens {
		removeEntry, errRemoveEntry := RemoveEntry(&RemoveEntryParams{
			Environment: TestEnvironment,
			Signature: token.Token.Signature,
		})
		if errRemoveEntry != nil {
			t.Error(errRemoveEntry.Error())
		}
		if removeEntry == false {
			t.Error("couldn't remove entry")
		}
	}
}
