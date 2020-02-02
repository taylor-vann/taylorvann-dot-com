package jwtx

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
)

// Header - Standard
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type MilliSecond = int64

// Claims - Payload body of a JWT
type Claims struct {
	Iss string      `json:"iss"` // taylorvann-dot-com
	Sub string      `json:"sub"` // subject: public, internal, infra, auth
	Aud string      `json:"aud"` // audience: username
	Iat MilliSecond `json:"iat"` // timestamp
	Exp MilliSecond `json:"exp"` // timestamp
}

// Signature - Unique hash of header and payload
type Signature = string

// Token - Contains contents of
type Token struct {
	Header       *string
	Payload      *string
	Signature    *string
	RandomSecret *[]byte
}

// RandomSecretLength - minimum secret byte array length
const RandomSecretLength = 512

var headerDefaultParams = Header{
	Alg: "HS256",
	Typ: "JWT",
}

// HeaderBase64 - Default payload for all JWTs
var HeaderBase64 = createDefaultHeaderAsBase64(&headerDefaultParams)

func generateRandomBytes(n uint32) (*[]byte, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func createDefaultHeaderAsBase64(h *Header) *string {
	marhalledHeader, err := json.Marshal(h)
	if err != nil {
		return nil
	}
	b64Header := base64.RawStdEncoding.EncodeToString(marhalledHeader)

	return &b64Header
}

func createPayloadAsBase64(claims *Claims) (*string, error) {
	if HeaderBase64 == nil {
		return nil, errors.New("createPayloadAsBase64 - header is nil")
	}

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
		marshalledPayload,
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
		signature.Sum(nil),
	)

	return &signatureBase64
}

// CreateJWT - Return JWT Token
func CreateJWT(
	claims *Claims,
) (*Token, error) {
	marshalledPayloadBase64, errPayload := createPayloadAsBase64(claims)
	if errPayload != nil {
		return nil, errPayload
	}

	randomSecret, errRandomSecret := generateRandomBytes(512)
	if errRandomSecret != nil {
		return nil, errRandomSecret
	}

	signatureBase64 := generateSignature(
		HeaderBase64,
		marshalledPayloadBase64,
		randomSecret,
	)

	token := Token{
		Header:       HeaderBase64,
		Payload:      marshalledPayloadBase64,
		Signature:    signatureBase64,
		RandomSecret: randomSecret,
	}

	return &token, nil
}

// ValidateJWT - take token payload and verify signature
func ValidateJWT(token *Token) bool {
	signatureBase64 := generateSignature(
		token.Header,
		token.Payload,
		token.RandomSecret,
	)

	if *token.Signature == *signatureBase64 {
		return true
	}

	return false
}
