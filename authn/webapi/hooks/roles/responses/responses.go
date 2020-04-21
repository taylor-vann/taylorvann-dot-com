package responses

import "webapi/controllers/roles"

type Roles = roles.Roles

type Errors struct {
	Body		*string `json:"body"`
	Roles 	*string `json:"session"`
	Default *string `json:"default"`
}

type Body struct {
	Roles  *Roles 	`json:"roles"`
	Errors *Errors	`json:"errors"`
}
