// Package hooks - create_user
//
// brian taylor vann
// taylorvann dot com
package hooks

import (
	"encoding/json"
	"net/http"
)

// Request
// {username, email, password, verified_password}

// Response?
// action - CREATED_USER
// payload - {User: username}

// CreateUser - get information about our api
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// check cookies for httpOnly header

	// insert, if user exists or error, return fail

	// add password, if error, return fail
	// add has_password, if error, return fail

	json.NewEncoder(w).Encode(authnDetails)
}
