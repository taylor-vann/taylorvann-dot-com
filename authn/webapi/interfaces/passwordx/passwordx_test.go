// Brian Taylor Vann
// taylorvann dot com

// Package passwordx - utility library for password hashing
package passwordx

import (
	"fmt"
	"testing"
)

type HashPasswordTestPlan []struct {
	Password string
}

type PasswordIsValidFailsTestPlan []struct {
	Password          string
	IncorrectPassword string
}

var HashPasswordPlan = HashPasswordTestPlan{
	{""},
	{"admin"},
	{"password"},
	{"1234567890"},
	{"hello, world"},
}

var PasswordFailsPlan = PasswordIsValidFailsTestPlan{
	{"", "hello, world"},
	{"admin", "1234567890"},
	{"password", "drowssap"},
	{"1234567890", "admin"},
	{"hello, world", ""},
}

var HashParamsDoubleCheck = HashParams{
	HashFunction: "argon2id",
	Memory:       64 * 1024,
	Iterations:   3,
	Parallelism:  2,
	SaltLength:   16,
	KeyLength:    32,
	BuildVersion: 0,
}

// Test Default Hash Params
func TestDefaultHashParams(t *testing.T) {
	if DefaultHashParams != HashParamsDoubleCheck {
		t.Error("Unrecognized change in default password hash params.")
	}
}

// Make sure Salt and Hash are not equal to a given password
// This is the point of hashing a password
func TestHashPassword(t *testing.T) {
	for index, test := range HashPasswordPlan {
		hashedPassword, err := HashPassword(test.Password, &DefaultHashParams)
		if err != nil {
			t.Error(
				fmt.Sprintf(
					"Error hashing password!\n%d: %s",
					index,
					test.Password,
				),
			)
			continue
		}

		if test.Password == hashedPassword.Hash {
			t.Error("Password failed to hash")
		}
	}
}

func TestPasswordIsValid(t *testing.T) {
	// hash a password
	// run is
	for index, test := range HashPasswordPlan {
		hashedPasswordResults, hashedPasswordErr := HashPassword(
			test.Password,
			&DefaultHashParams,
		)
		if hashedPasswordErr != nil {
			t.Error(
				fmt.Sprintf(
					"Error hashing password!\n%d: %s",
					index,
					test.Password,
				),
			)
			continue
		}

		passwordIsValid, passwordCheckErr := PasswordIsValid(
			test.Password,
			hashedPasswordResults,
		)

		if passwordCheckErr != nil {
			t.Error(
				fmt.Sprintf(
					"Error validating password!\n%d: %s",
					index,
					test.Password,
				),
			)
			continue
		}

		if !passwordIsValid {
			t.Error(
				fmt.Sprintf(
					"Assymetric password matching failed:\n%d: %s",
					index,
					test.Password,
				),
			)
		}
	}
}

func TestPasswordIsValidFails(t *testing.T) {
	for index, test := range PasswordFailsPlan {
		hashedPasswordResults, hashedPasswordErr := HashPassword(
			test.Password,
			&DefaultHashParams,
		)
		if hashedPasswordErr != nil {
			t.Error(
				fmt.Sprintf(
					"Error hashing password!\n%d: %s",
					index,
					test.Password,
				),
			)
			continue
		}

		passwordIsValid, passwordCheckErr := PasswordIsValid(
			test.IncorrectPassword,
			hashedPasswordResults,
		)

		if passwordCheckErr != nil {
			t.Error(
				fmt.Sprintf(
					"Error validating password!\n%d: %s",
					index,
					test.Password,
				),
			)
			continue
		}

		if passwordIsValid {
			t.Error(
				fmt.Sprintf(
					"Assymetric password matching succeeded where it shouldn't:\n%d: %s",
					index,
					test.Password,
				),
			)
		}
	}
}
