package requests

type Query struct {
	UserID int64 `json:"user_id"`
}

type QueryRange struct {
	StartIndex int64 `json:"start_index"`
	Length		 int64 `json:"length"`
}

type SearchRange struct {
	EmailSubstring string `json:"substring"`
	StartIndex 		 int64  `json:"start_index"`
	Length		 		 int64  `json:"length"`
}

type Params struct {
	Environment			string 			 `json:"environment"`
	UserCredentials	*Query 			 `json:"query"`
	QueryRange			*QueryRange	 `json:"query_range"`
	Search					*SearchRange `json:"search_range"`
}

type Body struct {
	Action string  `json:"action"`
	Params *Params `json:"params"`
}
