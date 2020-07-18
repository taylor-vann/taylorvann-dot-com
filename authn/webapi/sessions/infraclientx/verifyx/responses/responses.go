package responses

type Errors struct {
	RequestBody	*string `json:"request_body"`
	Default 		*string `json:"default"`
}

type Body struct {
	Errors	*Errors	`json:"errors"`
}
