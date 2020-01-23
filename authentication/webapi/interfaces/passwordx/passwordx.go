package passwordx

import (
	"crypto/rand"
	"encoding/base64"
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/argon2"
)

// Argon2IdParams - Parameters for generating password hashes
type HashParams struct {
	HashFunction string `json:"hash_function"`
	Memory       uint32 `json:"memory"`
	Iterations   uint32 `json:"iterations"`
	Parallelism  uint8  `json:"paralleslism"`
	SaltLength   uint32 `json:"salt_length"`
	KeyLength    uint32 `json:"key_length"`
}

// HashResults - Store for salt, hash, and paramters post hash
type HashResults struct {
	Salt   string `json:"salt"`
	Hash   string `json:"hash"`
	Params string `json:"params"`
}

// DefaultArgon2IdParams - Our default settings for Argon2id
var DefaultArgon2IdParams = func () *Argon2Params {
	params := Argon2IdParams{
		HashFunction: ,
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	return &params
}()

func generateSaltRandomBytes(n uint32) ([]byte, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// HashPassword - Hash password and return hash results
func HashPassword(password string, p *Argon2Params) (*HashResults, error) {
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		p.Iterations,
		p.Memory, 
		p.Parallelism,
		p.KeyLength,
	)

	// Base64 encode the salt and hashed password.
	saltBase64 := base64.RawStdEncoding.EncodeToString(salt)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hash)

	params, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	paramsBase64 := base64.RawStdEncoding.EncodeToString(params)
	encodedHash := SaltHashPayload{
		Salt:   saltBase64,
		Hash:   hashBase64,
		Params: paramsBase64,
	}

	return &encodedHash, nil
}

// IsPasswordValid - Compare a password to a hash result
func IsPasswordValid(givenPassword string, comparator *HashResults) (bool, error) {
	salt, err = base64.RawStdEncoding.DecodeString(comparator.Params.Salt)
	if err != nil {
			return false, err
	}

	comparatorHash, err = base64.RawStdEncoding.DecodeString(comparator.Params.Hash)
	if err != nil {
			return false, err
	}

	contrastHash := argon2.IDKey(
		[]byte(givenPassword),
		salt,
		comparator.Params.Iterations,
		comparator.Params.Memory,
		comparator.Params.Parallelism,
		comparator.Params.KeyLength,
	)

	if subtle.ConstantTimeCompare(comparatorHash, contrashHash) == 1 {
		return true, nil
	}

	return false, nil
}