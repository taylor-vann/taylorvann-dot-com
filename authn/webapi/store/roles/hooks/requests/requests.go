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

type Params struct {
	Environment		string				`json:"environment"`
	ActionParams	interface{}		`json:"action_params`
	// Create				*Create				`json:"create"`
	// Read					*Read					`json:"read"`
	// Index					*Index				`json:"index"`
	// Search  			*Search				`json:"search"`
	// Update  			*Update				`json:"update"`
	// UpdateAccess  *UpdateAccess	`json:"update_access"`
	// Delete  			*Delete				`json:"delete"`
	// Undelete  		*Undelete			`json:"undelete"`
}

type Body struct {
	Action string  `json:"action"`
	Params *Params `json:"params"`
}
