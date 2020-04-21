package responses

import "webapi/controllers/users"

type Users = roles.Users

type Errors struct {
	Body		*string `json:"body"`
	Roles 	*string `json:"session"`
	Default *string `json:"default"`
}

type Body struct {
	Users  *Roles 	`json:"users"`
	Errors *Errors	`json:"errors"`
}
