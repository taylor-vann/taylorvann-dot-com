package passwordx

import (
	"crypto/rand"
	"encoding/base64"
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Argon2Params struct {
	HashFunction string `json:"hash_function"`
	Memory       uint32 `json:"memory"`
	Iterations   uint32 `json:"iterations"`
	Parallelism  uint8  `json:"paralleslism"`
	SaltLength   uint32 `json:"salt_length"`
	KeyLength    uint32 `json:"key_length"`
}

type SaltHashPayload struct {
	Salt   string        `json:"salt"`
	Hash   string        `json:"hash"`
	Params *Argon2Params `json:"params"`
}

func getArgon2Params() *Argon2Params {
	params := Argon2Params{
		HashFunction: "argon2id"
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	return &params
}

func generateSaltRandomBytes(n uint32) ([]byte, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func EncodeHashPayload(password string, p *Argon2Params) (*SaltHashPayload, error) {
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
		p.KeyLength
	)

	// Base64 encode the salt and hashed password.
	saltB64 := base64.RawStdEncoding.EncodeToString(salt)
	hashB64 := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := EncodedHash{
		Salt:   saltB64,
		Hash:   hashB64,
		Params: p,
	}

	return &encodedHash, nil
}

func PasswordIsValid(givenPassword string, comparator *SaltHashPayload) (bool, error) {
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
		comparator.Params.KeyLength
	)

	if subtle.ConstantTimeCompare(comparatorHash, contrashHash) == 1 {
		return true, nil
	}

	return false, nil
}