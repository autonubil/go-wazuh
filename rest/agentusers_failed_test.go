package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// A response with error != 0 still names, per item, which agents failed and
// why; the caller needs that to tell "no inventory" from a failing manager.
func TestGetAgentUsersFailedItems(t *testing.T) {
	const body = `{"data":{"affected_items":[],"total_affected_items":0,"total_failed_items":1,` +
		`"failed_items":[{"error":{"code":42,"message":"item failed"},"id":["007"]}]},` +
		`"message":"nothing returned","error":1}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(body))
	}))
	defer srv.Close()

	cl, err := NewClient(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	cl.token = "test"
	ctrl := &SyscollectorController{&ClientWithResponses{cl}}

	resp, err := ctrl.GetAgentUsers("007", nil)
	if resp != nil {
		t.Fatalf("resp = %+v, want nil", resp)
	}
	if err == nil || err.Error() != "API error 1: nothing returned" {
		t.Fatalf("err = %v, want %q", err, "API error 1: nothing returned")
	}
}
