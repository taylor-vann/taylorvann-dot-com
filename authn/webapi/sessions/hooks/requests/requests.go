package requests

import (
	"webapi/sessions/sessionsx"
)

type Create = sessionsx.CreateParams
type Read = sessionsx.ReadParams
type Validate = sessionsx.ReadParams
type ValidateGuest = sessionsx.ReadParams
type Update = sessionsx.UpdateParams
type Delete = sessionsx.DeleteParams

type Guest struct {
	Environment string	`json:"environment"`
}

type User struct {
	Environment string `json:"environment"`
	UserID      int64  `json:"user_id"`
}

type Ancillary struct {
	Environment string `json:"environment"`
	Email      string  `json:"email"`
}

type InfraUser struct {
	Environment string `json:"environment"`
	Email				string `json:"email"`
	Password		string `json:"password"`
}

type Body struct {
	Action string 		 `json:"action"`
	Params interface{} `json:"params"`
}
