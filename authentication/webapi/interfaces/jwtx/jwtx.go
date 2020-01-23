package jwtx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

// Header - Standard
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// Claims - Payload body of a JWT
type Claims struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"` // subject: public, internal, infra, auth
	Aud string `json:"aud"` // audience: username
	Exp string `json:"exp"`
	Iat string `json:"iat"`
}

// Signature - Unique hash of header and payload
type Signature = string

// Token - Contains contents of
type Token struct {
	Header    *string
	Payload   *string
	Signature *string
}

var headerDefaultParams = Header{
	Alg: "HS256",
	Typ: "JWT",
}

// headerDefaultBase64 - Default payload for all JWTs
var headerDefaultBase64 = func() *string {
	marhalledHeader, err := json.Marshal(headerDefaultParams)
	if err != nil {
		return nil
	}
	b64Header := base64.RawStdEncoding.EncodeToString(marhalledHeader)

	return &b64Header
}()

func createPayloadAsBase64(claims *Claims) (*string, error) {
	PayloadResult := Claims{
		Iss: claims.Iss,
		Sub: claims.Sub,
		Aud: claims.Aud,
		Exp: claims.Exp,
		Iat: claims.Iat,
	}

	marshalledPayload, err := json.Marshal(PayloadResult)
	if err != nil {
		return nil, err
	}

	marshalledPayloadBase64 := base64.RawStdEncoding.EncodeToString(
		marshalledPayload
	)

	return &marshalledPayloadBase64, nil
}

func generateSignature(
	header *string,
	payload *string,
	randomSecret *[]byte,
) *string {
	combinedHeaderAndPayload := *header + "." + *payload
	signature := hmac.New(sha256.New, *randomSecret)
	signature.Write([]byte(combinedHeaderAndPayload))
	signatureBase64 := base64.RawStdEncoding.EncodeToString(
		signature.Sum(nil)
	)

	return &signatureBase64
}

// Create - Return JWT Token
func Create(
	claims *Claims,
	randomHash *[]byte,
	randomSecret *[]byte,
) (*Token, error) {
	marshalledPayloadBase64, err := createPayloadAsBase64(
		claims,
	)
	if err != nil {
		return nil, err
	}

	signatureBase64 := generateSignature(
		headerDefaultBase64,
		marshalledPayloadBase64,
		randomSecret,
	)

	token := Token{
		Header:    headerDefaultBase64,
		Payload:   marshalledPayloadBase64,
		Signature: signatureBase64,
	}

	return &token, nil
}
