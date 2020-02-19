// brian Taylor Vann
// taylorvann dot com

// Package jwtx - utility library for creating JWT based Session Tokens
package jwtx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"webapi/utils"
)

// MilliSeconds -
type MilliSeconds = int64

// Header - Standard
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// Claims - Payload body of a JWT
type Claims struct {
	Iss string      				`json:"iss"` // taylorvann-dot-com
	Sub string      				`json:"sub"` // subject: public, internal, infra, auth
	Aud string      				`json:"aud"` // audience: guest, public
	Iat utils.MilliSeconds	`json:"iat"` // timestamp
	Exp utils.MilliSeconds	`json:"exp"` // timestamp
}

// TokenDetails -
type TokenDetails struct {
	Header			*Header
	Payload 		*Claims
	Signature		*string
}

// Signature - Unique hash of header and payload
type Signature = string

// Token - Contains contents of
type Token struct {
	Header    string
	Payload   string
	Signature string
}

// TokenPayload -
type TokenPayload struct {
	Token					*Token
	RandomSecret	*[]byte
}

var headerDefaultParams = Header{
	Alg: "HS256",
	Typ: "JWT",
}

// DayAsMS -
var DayAsMS = int64(1000 * 60 * 60 * 24)

// ThreeDaysAsMS -
var ThreeDaysAsMS = 3 * DayAsMS

// HeaderBase64 - Default payload for all JWTs
var HeaderBase64 = createDefaultHeaderAsBase64(&headerDefaultParams)

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
		return nil, errors.New("jwtx.createPayloadAsBase64() - header is nil")
	}

	payloadResult := Claims{
		Iss: claims.Iss,
		Sub: claims.Sub,
		Aud: claims.Aud,
		Exp: claims.Exp,
		Iat: claims.Iat,
	}

	marshalledPayload, err := json.Marshal(payloadResult)
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
) (*TokenPayload, error) {
	marshalledPayloadBase64, errPayload := createPayloadAsBase64(claims)
	if errPayload != nil {
		return nil, errPayload
	}

	randomSecret, errRandomSecret := utils.GenerateRandomByteArray(512)
	if errRandomSecret != nil {
		return nil, errRandomSecret
	}

	signatureBase64 := generateSignature(
		HeaderBase64,
		marshalledPayloadBase64,
		randomSecret,
	)

	token := Token{
		Header:       *HeaderBase64,
		Payload:      *marshalledPayloadBase64,
		Signature:    *signatureBase64,
	}

	tokenPayload := TokenPayload{
		Token: &token,
		RandomSecret: randomSecret,
	}

	return &tokenPayload, nil
}

// ValidateJWT - take token payload and verify signature
func ValidateJWT(p *TokenPayload) bool {
	signatureBase64 := generateSignature(
		&p.Token.Header,
		&p.Token.Payload,
		p.RandomSecret,
	)

	if p.Token.Signature == *signatureBase64 {
		return true
	}

	return false
}

// ConvertTokenToString -
func ConvertTokenToString(token *Token) (*string, error) {
	if token == nil {
		return nil, errors.New("jwtx.ConvertTokenToString() - nil token provided")
	}

	tokenStr := fmt.Sprintf(
		"%s.%s.%s",
		token.Header,
		token.Payload,
		token.Signature,
	)

	return &tokenStr, nil
}

// RetrieveTokenFromString -
func RetrieveTokenFromString(tokenStr *string) (*Token, error) {
	if tokenStr == nil {
		return nil, errors.New("jwtx.RetrieveTokenFromString() - nil token provided")
	}

	bricks := strings.Split(*tokenStr, ".")
	if len(bricks) != 3 {
		return nil, errors.New("jwtx.RetrieveTokenFromString() - invalid token")
	}

	token := Token{
		Header:    bricks[0],
		Payload:   bricks[1],
		Signature: bricks[2],
	}

	return &token, nil
}

// RetrieveTokenDetailsFromString -
func RetrieveTokenDetailsFromString(tokenStr *string) (*TokenDetails, error) {
	if tokenStr == nil {
		return nil, errors.New("jwtx.RetrieveTokenFromString() - nil token provided")
	}

	token, errToken := RetrieveTokenFromString(tokenStr)
	if errToken != nil {
		return nil, errToken
	}

	headerDecoded, errHeaderDecoded := base64.RawStdEncoding.DecodeString(
		token.Header,
	)
	if errHeaderDecoded != nil {
		return nil, errHeaderDecoded
	}
	var header Header
	errHeaderMarshal := json.Unmarshal(headerDecoded, &header)
	if errHeaderMarshal != nil {
		return nil, errHeaderMarshal
	}

	payloadDecoded, errPayloadDecoded := base64.RawStdEncoding.DecodeString(
		token.Payload,
	)
	if errPayloadDecoded != nil {
		return nil, errPayloadDecoded
	}

	var payload Claims
	errPayloadMarshal := json.Unmarshal(payloadDecoded, &payload)
	if errPayloadMarshal != nil {
		return nil, errPayloadMarshal
	}

	tokenDetails := TokenDetails{
		Header: &header,
		Payload: &payload,
		Signature: &token.Signature,
	}

	return &tokenDetails, nil
}
