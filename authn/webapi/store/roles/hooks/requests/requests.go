package requests

import (
	"webapi/store/roles/controller"
)

type Create = controller.CreateParams
type Read = controller.ReadParams
type Index = controller.IndexParams
type Search = controller.SearchParams
type Update = controller.UpdateParams
type UpdateAccess = controller.UpdateAccessParams
type Delete = controller.DeleteParams
type Undelete = controller.UndeleteParams

type ValidateInfra struct {
	Environment string `json:"environment"`
	Email				string `json:"email"`
	Password		string `json:"password"`
}

type Body struct {
	Action string  		 `json:"action"`
	Params interface{} `json:"params"`
}
