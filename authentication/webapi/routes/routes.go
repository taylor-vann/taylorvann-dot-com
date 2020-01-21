package routes

import (
	"encoding/json"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello world, json!")
}

// CreateRoutes - add hooks to route callbacks
func CreateRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", helloWorld)
	mux.HandleFunc("/session", helloWorld)
	mux.HandleFunc("/user", helloWorld)

	return mux
}

// we are assuming jwt is unique
// in redis as a key, value store:
// sessions:user_id -> a map of jwts to their random keys and devices details
// jwt:jwt_id -> user_id, random key

// Session Controller
// login user - if no jwt, check password, issue jwt, whitelist jwt
// logout user - if jwt is whitelisted, remove jwt from whitelist

// User Controller
// create user - if no jwt, check if user does not exist, check payload, create user, issue jwt, whitelist jwt
// update user - if jwt, check payload, update password, invalidate all jwts associated with user

// verify jwt is whitelisted
// this method should only be accepted from dedicated ips
// this could return true and the code
