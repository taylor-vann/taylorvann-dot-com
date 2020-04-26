package requests

import (
	"webapi/store/users/controller"
)

type Create = controller.CreateParams
type Read = controller.ReadParams
type Index = controller.IndexParams
type Search = controller.SearchParams
type Update = controller.UpdateParams
type UpdateEmail = controller.UpdateEmailParams
type UpdatePassword = controller.UpdatePasswordParams
type Delete = controller.DeleteParams
type Undelete = controller.UndeleteParams

type Params struct {
	Environment			string					`json:"environment"`
	Create					*Create					`json:"create"`
	Read						*Read						`json:"read"`
	Index						*Index					`json:"index"`
	Search  				*Search					`json:"search"`
	Update  				*Update					`json:"update"`
	UpdateEmail  	 	*UpdateEmail		`json:"update_email"`
	UpdatePassword	*UpdatePassword	`json:"update_password"`
	Delete  				*Delete					`json:"delete"`
	Undelete  			*Undelete				`json:"undelete"`
}

type Body struct {
	Action string  `json:"action"`
	Params *Params `json:"params"`
}