package responses

type Errors struct {
	Headers	*string `json:"headers"`
	Body		*string `json:"body"`
	Mail	  *string `json:"mail"`
	Default *string `json:"default"`
}

type Body struct {
	Mail 	 *string	`json:"session"`
	Errors *Errors	`json:"errors"`
}