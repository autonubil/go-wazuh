package rest

import (
	"errors"
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
	var fe *FailedItemsError
	if !errors.As(err, &fe) {
		t.Fatalf("err is %T, want *FailedItemsError", err)
	}
	if fe.Code != 1 || len(fe.FailedItems) != 1 {
		t.Fatalf("FailedItemsError = %+v", fe)
	}
	item := fe.FailedItems[0]
	if item.Error.Code != 42 || item.Error.Message != "item failed" || len(item.IDs) != 1 || item.IDs[0] != "007" {
		t.Fatalf("failed item = %+v", item)
	}
}
