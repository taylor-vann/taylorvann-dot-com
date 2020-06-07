package requests

type GuestSessionParams struct {
	Environment string
}

type ValidateParams struct {
	Environment string
	Token 			string
}

type Body struct {
	Action string 		 `json:"action"`
	Params interface{} `json:"params"`
}