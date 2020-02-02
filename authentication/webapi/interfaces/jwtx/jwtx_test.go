package jwtx

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

type JWTClaimTestPlan = []Claims

const dayAsMS = 24 * 60 * 60 * 1000

var HeaderParamsDoubleCheck = Header{
	Alg: "HS256",
	Typ: "JWT",
}

var randomJWTClaims = generateRandomJWTClaims("public", 5)

func getNowAsMS() MilliSecond {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func getLaterAsMS() MilliSecond {
	return (time.Now().UnixNano() + dayAsMS) / int64(time.Millisecond)
}

func generateRandomJWTClaims(sub string, num int) *JWTClaimTestPlan {
	jwtClaims := make(JWTClaimTestPlan, num)

	for index := range jwtClaims {
		nowAsMS := getNowAsMS()
		laterAsMS := getLaterAsMS()

		jwtClaims[index] = Claims{
			Iss: "taylorvann_dot_com",
			Sub: sub,
			Aud: "username" + strconv.Itoa(index),
			Iat: nowAsMS,
			Exp: laterAsMS,
		}
	}

	fmt.Println(jwtClaims)
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
	var tokens = make([]*Token, len(*randomJWTClaims))

	for index, claim := range *randomJWTClaims {
		token, errToken := CreateJWT(&claim)
		if errToken != nil {
			fmt.Println(errToken)
			fmt.Println(index)
			t.Error("Unable to create jwt")
		}
		tokens[index] = token
	}
}

func TestValidateJWT(t *testing.T) {
	for index, claim := range *randomJWTClaims {
		token, errToken := CreateJWT(&claim)
		if errToken != nil {
			fmt.Println(errToken)
			fmt.Println(index)
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
	var tokens = make([]*Token, tokenSetLength)

	for index, claim := range *randomJWTClaims {
		token, errToken := CreateJWT(&claim)
		if errToken != nil {
			t.Error("Unable to create jwt")
		}
		tokens[index] = token
	}

	for index, token := range tokens {
		offsetIndex := (index + 1) % tokenSetLength
		offsetToken := tokens[offsetIndex]

		badToken := Token{
			Header:       token.Header,
			Payload:      offsetToken.Payload,
			Signature:    token.Signature,
			RandomSecret: token.RandomSecret,
		}
	
		isValid := ValidateJWT(&badToken)
		if isValid {
			t.Error("A bad token should not be valid")
		}
	}
}