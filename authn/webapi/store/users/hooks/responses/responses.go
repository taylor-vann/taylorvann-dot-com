package responses

import "webapi/store/users/controller"

type Users = controller.Users

type Errors struct {
	Body		*string `json:"body"`
	Users 	*string `json:"users"`
	Default *string `json:"default"`
}

type Body struct {
	Users  *Users 	`json:"users"`
	Errors *Errors	`json:"errors"`
}
