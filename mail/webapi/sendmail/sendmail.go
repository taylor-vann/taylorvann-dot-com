// brian taylor vann
// taylorvann dot com

package sendmail

import (
	"encoding/json"
	"net/http"
)

type EmailParams struct{
	Sender	 string
	Body 		 string
	Recpient string
}

func NoReply(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello world")
}
