package responses

import "webapi/store/roles/controller"

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
