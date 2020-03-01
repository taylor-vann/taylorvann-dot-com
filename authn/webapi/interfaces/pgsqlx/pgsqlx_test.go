// brian taylor vann
// taylorvann dot com

package pgsqlx

import (
	"testing"
)

func TestCreate(t *testing.T) {
	conn, errConn := Create(nil)
	if conn != nil {
		t.Error("nil parameters should return nil")
	}
	if errConn == nil {
		t.Error("nil paramters should return error")
	}
}
