package headersessionx

import (
	// "net/http"
	// "net/http/httptest"
	"testing"
)

func TestFetchGuestSession(t *testing.T) {
	resp, errResp := FetchGuestSession()

	if errResp != nil {
		t.Error(errResp)
	}
 
	if resp == "" {
		t.Error("empty session string returned")
	}
}