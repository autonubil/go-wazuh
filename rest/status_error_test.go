package rest

import (
	"errors"
	"net/http"
	"testing"
)

// A status the response type has no field for, answered with the manager's
// usual error body (shape from Wazuh's API error handler; values are minimal).
func unexpectedStatus(body string) *AgentControllerGetAgentConfigResponse {
	return &AgentControllerGetAgentConfigResponse{
		Body:         []byte(body),
		HTTPResponse: &http.Response{StatusCode: http.StatusInternalServerError, Status: "500 Internal Server Error"},
	}
}

func TestGetResponseObjectUnexpectedStatus(t *testing.T) {
	_, err := getResponseObject(unexpectedStatus(`{"title":"Wazuh Internal Error","detail":"probe detail","error":1000}`))
	if err == nil {
		t.Fatal("no error for a 500")
	}
	if got, want := err.Error(), "500 Internal Server Error: Wazuh Internal Error (probe detail)"; got != want {
		t.Errorf("error = %q, want %q", got, want)
	}
	var re *RequestError
	if !errors.As(err, &re) || re.RequestError == nil || *re.RequestError != 1000 {
		t.Fatalf("manager error body not available: %+v", re)
	}
	var se *StatusError
	if !errors.As(err, &se) || se.Status != "500 Internal Server Error" {
		t.Errorf("StatusError = %+v", se)
	}
}

// Without a manager error body the error stays the bare status, as before.
func TestGetResponseObjectUnexpectedStatusWithoutBody(t *testing.T) {
	for _, body := range []string{"", "<html><body>Bad Gateway</body></html>", `{"unrelated":true}`} {
		_, err := getResponseObject(unexpectedStatus(body))
		if err == nil || err.Error() != "500 Internal Server Error" {
			t.Errorf("body %q: error = %v, want the bare status", body, err)
		}
		var re *RequestError
		if errors.As(err, &re) {
			t.Errorf("body %q: unexpected RequestError %+v", body, re)
		}
	}
}
