package requests

import (
	"webapi/store/roles/controller"
)

type Read struct = contoller.ReadPara
type Index = controller.IndexParams
type Search = controller.SearchParams
type Update = controller.UpdateParams
type UpdateAccess = controller.UpdateAccessParams
type Delete = controller.DeleteParams
type Undelete = controller.UndeleteParams

type Params struct {
	Environment			string 			 `json:"environment"`
	Payload					interface{}  `json:"payload"`
}


type Body struct {
	Action string  `json:"action"`
	Params *Params `json:"params"`
}
