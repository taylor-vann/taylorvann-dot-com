package requests

import (
	"webapi/store/users/controller"
)

type Create = controller.CreateParams
type Read = controller.ReadParams
type Validate = controller.ValidateParams
type Index = controller.IndexParams
type Search = controller.SearchParams
type Update = controller.UpdateParams
type UpdateEmail = controller.UpdateEmailParams
type UpdatePassword = controller.UpdatePasswordParams
type Delete = controller.DeleteParams
type Undelete = controller.UndeleteParams

type Body struct {
	Action string			 `json:"action"`
	Params interface{} `json:"payload"`
}