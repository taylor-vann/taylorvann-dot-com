// Brian Taylor Vann
// toolbox-go

package passwordx

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

type HashParams struct {
	HashFunction string `json:"hash_function"`
	Memory       uint32 `json:"memory"`
	Iterations   uint32 `json:"iterations"`
	Parallelism  uint8  `json:"paralleslism"`
	SaltLength   uint32 `json:"salt_length"`
	KeyLength    uint32 `json:"key_length"`
	BuildVersion uint32 `json:"build_version"`
}

type HashResults struct {
	Salt   string     `json:"salt"`
	Hash   string     `json:"hash"`
	Params HashParams `json:"params"`
}

var DefaultHashParams = HashParams{
	HashFunction: "argon2id",
	Memory:       64 * 1024,
	Iterations:   3,
	Parallelism:  2,
	SaltLength:   16,
	KeyLength:    32,
	BuildVersion: 0,
}

func generateSaltRandomBytes(n uint32) (*[]byte, error) {
	token := make([]byte, n)
	_, err := rand.Read(token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func HashPassword(password string, p *HashParams) (*HashResults, error) {
	salt, err := generateSaltRandomBytes(p.SaltLength)
	if err != nil {
		return nil, err
	}

	hash := argon2.IDKey(
		[]byte(password),
		*salt,
		p.Iterations,
		p.Memory,
		p.Parallelism,
		p.KeyLength,
	)

	saltBase64 := base64.RawStdEncoding.EncodeToString(*salt)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := HashResults{
		Salt:   saltBase64,
		Hash:   hashBase64,
		Params: *p,
	}

	return &encodedHash, nil
}

func PasswordIsValid(givenPassword string, comparator *HashResults) (bool, error) {
	if comparator == nil {
		return false, errors.New("nil hash results given")
	}
	salt, err := base64.RawStdEncoding.DecodeString(comparator.Salt)
	if err != nil {
		return false, err
	}

	comparatorHash, err := base64.RawStdEncoding.DecodeString(comparator.Hash)
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

	if subtle.ConstantTimeCompare(comparatorHash, contrastHash) == 1 {
		return true, nil
	}

	return false, nil
}
