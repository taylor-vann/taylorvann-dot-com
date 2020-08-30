package responses

import "webapi/store/roles/controller"

type Roles = controller.Roles

type Errors struct {
	RequestBody *string `json:"request_body"`
	Roles       *string `json:"roles"`
	Default     *string `json:"default"`
}

type Body struct {
	Roles  *Roles  `json:"roles"`
	Errors *Errors `json:"errors"`
}
