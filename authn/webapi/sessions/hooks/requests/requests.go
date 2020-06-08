package requests

import (
	"webapi/sessions/sessionsx"
)

type Create = sessionsx.CreateParams
type Read = sessionsx.ReadParams
type Validate = sessionsx.ReadParams
type ValidateGuest = sessionsx.CreateParams
type Update = sessionsx.UpdateParams
type Delete = sessionsx.DeleteParams
type User = sessionsx.UserParams
type Account = sessionsx.AccountParams

type Guest struct {
	Environment string
}

type InfraUser struct {
	Environment string `json:"environment"`
	Email				string `json:"email"`
	Password		string `json:"password"`
}

type ClientUser = InfraUser

type Body struct {
	Action string 		 `json:"action"`
	Params interface{} `json:"params"`
}
