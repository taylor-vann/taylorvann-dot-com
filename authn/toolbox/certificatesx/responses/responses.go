package responses

type Certificates struct {
	Cert []byte	`json:"cert"`
	Key	 []byte `json:"key"`
}

type Errors struct {
	Body					string
	Certificates	string
	Default				string
}

type Body struct {
	Errors 			 *Errors
	Certificates *Certificates
}