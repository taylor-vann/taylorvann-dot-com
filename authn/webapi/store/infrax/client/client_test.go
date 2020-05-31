package client

import (
	"testing"
)

func TestInit(t *testing.T) {
	if Client == nil {
		t.Error("cookie jar is nil")
		return
	}
	if Client.Jar == nil {
		t.Error("cookie jar is nil")
	}
}
