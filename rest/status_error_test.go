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
	if got, want := err.Error(), "500 Internal Server Error"; got != want {
		t.Errorf("error = %q, want %q", got, want)
	}
	var re *RequestError
	if errors.As(err, &re) {
		t.Errorf("manager error body unexpectedly available: %+v", re)
	}
}
