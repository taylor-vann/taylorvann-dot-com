package requests

type GuestSessionParams struct {
	Environment string  `json:"environment"`
}

type ValidateSessionParams struct {
	Environment string `json:"environment"`
	Token 			string `json:"token"`
}

type ValidateUserParams struct {
	Environment string `json:"environment"`
	Email				string `json:"email"`
	Password		string `json:"password"`
}

type Body struct {
	Action string 		 `json:"action"`
	Params interface{} `json:"params"`
}