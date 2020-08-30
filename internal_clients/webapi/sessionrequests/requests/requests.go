package requests

type CreateSessionParams struct {
	Environment string `json:"environment"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type User struct {
	Environment string `json:"environment"`
	UserID      int64  `json:"user_id"`
}

type RemoveSessionRequestParams struct {
	Environment string `json:"environment"`
}

type RemoveSessionParams struct {
	Environment string `json:"environment"`
	Signature   string `json:"signature"`
}

// type RequestSessionBody struct {
// 	Action string			 `json:"action"`
// 	Params interface{} `json:"params"`
// }

type RemoveSessionBody struct {
	Action string               `json:"action"`
	Params *RemoveSessionParams `json:"params"`
}

type ValidateUserBody struct {
	Action string               `json:"action"`
	Params *CreateSessionParams `json:"params"`
}

type RequestSessionBody struct {
	Action string `json:"action"`
	Params *User  `json:"params"`
}
