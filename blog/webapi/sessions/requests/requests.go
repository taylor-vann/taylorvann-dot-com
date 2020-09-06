package requests

type CreateGuestSessionParams struct {
	Environment string `json:"environment"`
}

type CreateClientSessionParams struct {
	Environment string `json:"environment"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type RemoveSessionParams = CreateGuestSessionParams

type RemoveSessionRequestParams struct {
	Environment string `json:"environment"`
	Signature   string `json:"signature"`
}

type User struct {
	Environment string `json:"environment"`
	UserID      int64  `json:"user_id"`
}

type RequestBody struct {
	Action string      `json:"action"`
	Params interface{} `json:"params"`
}
