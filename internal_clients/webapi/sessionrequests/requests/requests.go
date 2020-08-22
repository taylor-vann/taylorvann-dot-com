package requests

type CreateParams struct {
	Environment string `json:"environment"`
	Email    		string `json:"email"`	
	Password 		string `json:"password"`
}

type User struct {
	Environment string `json:"environment"`
	UserID      int64  `json:"user_id"`
}

type RemoveSessionParams struct {
	Environment	string `json:"environment"`
	Signature		string `json:"signature"`
}

type Body struct {
	Action string			 `json:"action"`
	Params interface{} `json:"params"`
}