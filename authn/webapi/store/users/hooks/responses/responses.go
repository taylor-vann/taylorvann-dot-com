package responses

import "webapi/store/users/controller"

type Users = controller.Users

type Errors struct {
	Body		*string `json:"body"`
	Roles 	*string `json:"session"`
	Default *string `json:"default"`
}

type Body struct {
	Users  *Roles 	`json:"users"`
	Errors *Errors	`json:"errors"`
}
