package responses

import "webapi/store/users/controller"

type Users = controller.SafeUsers

type Errors struct {
	RequestBody *string `json:"request_body"`
	Users       *string `json:"users"`
	Default     *string `json:"default"`
}

type Body struct {
	Users  *Users  `json:"users"`
	Errors *Errors `json:"errors"`
}
