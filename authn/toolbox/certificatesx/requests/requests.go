package requests

type Params struct {
	Email		 string
	Password string
}

type Body struct {
	Action	string
	Params	*Params
}
