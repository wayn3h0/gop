package http

import (
	"net/http"
	"testing"

	testing2 "github.com/wayn3h0/gop/testing"
)

func TestRequest(t *testing.T) {
	key1, key2 := 1, 2
	value1, value2 := 1, 2
	r, _ := http.NewRequest("GET", "http://localhost", nil)

	req := NewRequest(r)
	req.SetParameter(key1, value1)

	testing2.ExpectEqual(t, req.Parameter(key1), value1)
	testing2.ExpectEqual(t, len(req.Parameters()), 1)

	req.SetParameter(key2, value2)
	testing2.ExpectEqual(t, req.Parameter(key2), value2)
	testing2.ExpectEqual(t, len(req.Parameters()), 2)
}
