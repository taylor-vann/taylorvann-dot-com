package infrax

// func TestInitHasCookiePresent(t *testing.T) {
// 	urlStr, errUrlStr := url.Parse(Domain)
// 	if errUrlStr != nil {
// 		t.Error(errUrlStr)
// 	}

// 	if len(Client.Jar.Cookies(urlStr)) < 1 {
// 		t.Error("no cookies present")
// 	}

// 	for _, cookie := range Client.Jar.Cookies(urlStr) {
// 		log.Println(cookie.Name, cookie.Value)
// 	}
	
// 	if Client.Jar == nil {
// 		t.Error("cookie jar is nil")
// 	}
// }

// func TestValidateUser(t *testing.T) {
// 	resp, errResp := FetchValidateUser()

// 	if errResp != nil {
// 		t.Error(errResp)
// 	}

// 	if resp == "" {
// 		t.Error("nil response returned")
// 	}

// 	urlStr, errUrlStr := url.Parse(Domain)
// 	log.Println(urlStr)
// 	if errUrlStr != nil {
// 		t.Error(errUrlStr)
// 	}

// 	t.Error("force an error ")
// 	for _, cookie := range Client.Jar.Cookies(urlStr) {
// 		log.Println(cookie.Name, cookie.Value)
// 	}
// 	if Client.Jar == nil {
// 		t.Error("cookie jar is nil")
// 	}
// }