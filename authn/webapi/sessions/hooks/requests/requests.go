package requests

import (
	"webapi/sessions/sessionsx"
)

type Create = sessionsx.CreateParams
type Read = sessionsx.ReadParams
type Update = sessionsx.UpdateParams
type Delete = sessionsx.DeleteParams

type UserParams struct {
	Environment string
	UserID			int64
}

type AccountParams struct {
	Environment string
	Email				string
}

type SessionParams struct {
	Environment string
}

type Body struct {
	Action string 		 `json:"action"`
	Params interface{} `json:"params"`
}
