// brian taylor vann
// taylorvann dot com

// Package jwtx - utility library for JWTs
package jwtx

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"webapi/utils"
)

type JWTClaimTestPlan = []Claims

var HeaderParamsDoubleCheck = Header{
	Alg: "HS256",
	Typ: "JWT",
}

var randomJWTClaims = generateRandomJWTClaims("client", 5)

func getLaterAsMS() MilliSeconds {
	return (time.Now().UnixNano() + DayAsMS) / int64(time.Millisecond)
}

func generateRandomJWTClaims(sub string, num int) *JWTClaimTestPlan {
	jwtClaims := make(JWTClaimTestPlan, num)

	for index := range jwtClaims {
		nowAsMS := utils.GetNowAsMS()
		laterAsMS := getLaterAsMS()

		jwtClaims[index] = Claims{
			Iss: "taylorvann_dot_com",
			Sub: sub,
			Aud: "username" + strconv.Itoa(index),
			Iat: nowAsMS,
			Exp: laterAsMS,
		}
	}

	return &jwtClaims
}

func TestDefaultHeaderParams(t *testing.T) {
	if HeaderBase64 == nil {
		t.Error("HeaderBase64 is nil")
	}

	marhalledHeader, err := json.Marshal(HeaderParamsDoubleCheck)
	if err != nil {
		t.Error("Error marshalling header as json")
	}
	headerTest := base64.RawStdEncoding.EncodeToString(marhalledHeader)

	if headerTest != *HeaderBase64 {
		t.Error("Unrecongnized change in JWT Header")
	}
}

func TestCreateJWT(t *testing.T) {
	var tokens = make([]*TokenPayload, len(*randomJWTClaims))

	for index, claim := range *randomJWTClaims {
		token, errToken := CreateJWT(&claim)
		if errToken != nil {
			t.Error("Unable to create jwt")
		}
		tokens[index] = token
	}
}

func TestValidateJWT(t *testing.T) {
	for _, claim := range *randomJWTClaims {
		token, errToken := CreateJWT(&claim)
		if errToken != nil {
			t.Error("Unable to create jwt")
		}

		isValid := ValidateJWT(token)
		if !isValid {
			t.Error("Unable to authenticate a valid JWT")
		}
	}
}

func TestFailValidateJWT(t *testing.T) {
	var tokenSetLength = len(*randomJWTClaims)
	var tokenPayloads = make([]*TokenPayload, tokenSetLength)

	for index, claim := range *randomJWTClaims {
		token, errToken := CreateJWT(&claim)
		if errToken != nil {
			t.Error("Unable to create jwt")
		}
		tokenPayloads[index] = token
	}

	for index, tokenPayload := range tokenPayloads {
		offsetIndex := (index + 1) % tokenSetLength
		offsetToken := tokenPayloads[offsetIndex]

		badToken := Token{
			Header:    tokenPayload.Token.Header,
			Payload:   offsetToken.Token.Payload,
			Signature: tokenPayload.Token.Signature,
		}
		badTokenPayload := TokenPayload{
			Token:        &badToken,
			RandomSecret: tokenPayload.RandomSecret,
		}

		isValid := ValidateJWT(&badTokenPayload)
		if isValid {
			t.Error("A bad token should not be valid")
		}
	}
}

func TestConvertTokenToString(t *testing.T) {
	for _, claim := range *randomJWTClaims {
		tokenPayload, errToken := CreateJWT(&claim)
		if errToken != nil {
			t.Error("Unable to create jwt")
		}

		strippedToken := Token{
			Header:    tokenPayload.Token.Header,
			Payload:   tokenPayload.Token.Payload,
			Signature: tokenPayload.Token.Signature,
		}

		_, errJwt := ConvertTokenToString(&strippedToken)
		if errJwt != nil {
			t.Error("Unable to convert to string")
		}
	}
}

func TestRetrieveTokenDetailsFromString(t *testing.T) {
	for _, claim := range *randomJWTClaims {
		tokenPayload, errToken := CreateJWT(&claim)
		if errToken != nil {
			t.Error("Unable to create jwt")
		}

		strippedToken := Token{
			Header:    tokenPayload.Token.Header,
			Payload:   tokenPayload.Token.Payload,
			Signature: tokenPayload.Token.Signature,
		}

		tokenStr, errJwt := ConvertTokenToString(&strippedToken)
		if errJwt != nil {
			t.Error("Unable to convert to string")
		}

		retrievedToken, errRetrieve := RetrieveTokenFromString(tokenStr)
		if errRetrieve != nil {
			t.Error("error retrieving token")
		}

		if retrievedToken.Header != tokenPayload.Token.Header {
			t.Error("mismatching headers")
		}

		if retrievedToken.Payload != tokenPayload.Token.Payload {
			t.Error("mismatching payloads")
		}

		if retrievedToken.Signature != tokenPayload.Token.Signature {
			t.Error("mismatching signatures")
		}
	}
}
