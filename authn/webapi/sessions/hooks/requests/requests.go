package requests

import (
	"webapi/sessions/sessionsx"
)

type Create = sessionsx.CreateParams
type Read = sessionsx.ReadParams
type Update = sessionsx.UpdateParams
type Delete = sessionsx.DeleteParams
type UserParams = sessionsx.UserParams
type AccountParams = sessionsx.AccountParams

type SessionParams struct {
	Environment string
}

type Body struct {
	Action string 		 `json:"action"`
	Params interface{} `json:"params"`
}
