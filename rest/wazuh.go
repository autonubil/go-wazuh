/*
go client for the wazuh [rest api](https://documentation.wazuh.com/4.0/user-manual/api/reference.html)

it is generated from the OpenAPI 3.0 specifications. Thus it is not the most elegant API. Some effort has been put into an more go friendly interface by wrapping non successful results into errors and returning the `Data` objects instead of the raw result.


There are a few With... option functions that can be used to customize the API client:

- `WithBaseURL` custom base url
- `WithLogin` (username, password)
- `WithContext` (custom Context)
- `WithInsecure` allow insecure certificates
- `WithUserAgent` to set custom user agent

go-wazuh supports following environment variables for easy construction of a client:

- `WAZUH_URL`
- `WAZUH_USER`
- `WAZUH_PASSWORD`
- `WAZUH_INSECURE`

Construct a new Wazuh client, then use the various service on the client to access different parts of the wazuh API. For example, to list all agents:

	c := NewAPIClient("https://localhost:55000", WithLogin("wazuh", "wazuh"), WithInsecure(true))
	c.Authenticate()
	agents := c.AgentsController.GetAgents(&AgentsControllerGetAgentsParams{})
	fmt.Printf("Get Agents TotalAffectedItems %d\n", agents.AllItemsResponse.TotalAffectedItems)
	for i, agent := range agents.AffectedItems {
		fmt.Printf(" %d: %s on %s\n", i, *agent.Id, *agent.NodeName)
	}


Or use the environment to construct the client to get the server basic information:


	c, err := NewClientFromEnvironment(WithInsecure(true))
	if err != nil {
		panic(err)
	}
	// authenticate
	err = c.Authenticate()
	if err != nil {
		panic(err)
	}

	// call the DefaultInfo on the
	status, err := c.Default.DefaultInfo(&DefaultControllerDefaultInfoParams{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to %s on %s\n", *status.Title, *status.Hostname)

*/

package rest

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"reflect"
	"strings"
)

// The Client for the wazuh REST API
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	innerClient HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn

	ctx       context.Context
	userAgent string
	token     string
	user      string
	password  string
	insecure  bool
	trace     bool
}

// WithLogin specifies the credentials for
func WithLogin(user string, password string) ClientOption {
	return func(c *Client) error {
		c.user = user
		c.password = password
		return nil
	}
}

// WithContext specifies the credentials for
func WithContext(ctx context.Context) ClientOption {
	return func(c *Client) error {
		c.ctx = ctx
		return nil
	}
}

// WithInsecure accept all certificates
func WithInsecure(insecure bool) ClientOption {
	return func(c *Client) error {
		c.insecure = insecure
		return nil
	}
}

// WithTrace write all requests to the log
func WithTrace(trace bool) ClientOption {
	return func(c *Client) error {
		c.trace = trace
		return nil
	}
}

// WithUserAgent specify a user agent string to identify the client
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.userAgent = userAgent
		return nil
	}
}

// do execute and evaluate the request
func (c *Client) do(ctx context.Context, req *http.Request) error {
	// Headers for all request
	req.Header.Set("User-Agent", c.userAgent)
	if c.token == "" {
		encoded := base64.StdEncoding.EncodeToString([]byte(c.user + ":" + c.password))
		req.Header.Set("Authorization", "Basic "+encoded)
	} else {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	return nil
}

// APIClient extended client with less abstract api access
type APIClient struct {
	*ClientWithResponses

	ExperimentalController   ExperimentalControllerInterface
	SyscheckController       SyscheckControllerInterface
	AgentsController         AgentsControllerInterface
	CiscatController         CiscatControllerInterface
	ListsController          ListsControllerInterface
	ManagerController        ManagerControllerInterface
	MitreController          MitreControllerInterface
	ScaController            ScaControllerInterface
	DefaultController        DefaultControllerInterface
	OverviewController       OverviewControllerInterface
	RulesController          RulesControllerInterface
	SecurityController       SecurityControllerInterface
	SyscollectorController   SyscollectorControllerInterface
	ActiveResponseController ActiveResponseControllerInterface
	ClusterController        ClusterControllerInterface
	DecodersController       DecodersControllerInterface
}

// NewClientFromEnvironment creates a new client from default environment variables
func NewClientFromEnvironment(opts ...ClientOption) (*APIClient, error) {
	baseURL := os.Getenv("WAZUH_URL")
	user := os.Getenv("WAZUH_USER")
	password := os.Getenv("WAZUH_PASSWORD")
	opts = append(opts, WithLogin(user, password))
	if os.Getenv("WAZUH_INSECURE") == "true" {
		opts = append(opts, WithInsecure(true))
	}

	c, err := NewAPIClient(baseURL, opts...)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// NewAPIClient Create a new API (yes, naming is awkward)
func NewAPIClient(baseURL string, opts ...ClientOption) (*APIClient, error) {
	cl, err := NewClient(baseURL, opts...)
	cl.RequestEditor = cl.do
	if err != nil {
		return nil, err
	}
	clientWithResponses := &ClientWithResponses{cl}
	return &APIClient{
		ClientWithResponses: clientWithResponses,

		ExperimentalController:   &ExperimentalController{clientWithResponses},
		SyscheckController:       &SyscheckController{clientWithResponses},
		ListsController:          &ListsController{clientWithResponses},
		ManagerController:        &ManagerController{clientWithResponses},
		MitreController:          &MitreController{clientWithResponses},
		ScaController:            &ScaController{clientWithResponses},
		AgentsController:         &AgentsController{clientWithResponses},
		CiscatController:         &CiscatController{clientWithResponses},
		RulesController:          &RulesController{clientWithResponses},
		SecurityController:       &SecurityController{clientWithResponses},
		SyscollectorController:   &SyscollectorController{clientWithResponses},
		DefaultController:        &DefaultController{clientWithResponses},
		OverviewController:       &OverviewController{clientWithResponses},
		DecodersController:       &DecodersController{clientWithResponses},
		ActiveResponseController: &ActiveResponseController{clientWithResponses},
		ClusterController:        &ClusterController{clientWithResponses},
	}, nil
}

// ServerAddress return the Wazuh server address
func (c *APIClient) ServerAddress() string {
	u, err := url.Parse(c.ClientInterface.(*Client).Server)
	if err != nil {
		return "127.0.0.1"
	}
	lastSep := strings.LastIndex(u.Host, ":")
	if lastSep > 0 {
		return u.Host[:lastSep]
	}
	return u.Host
}

// NewClient returns a new wazuh API client
func NewClient(baseURL string, opts ...ClientOption) (*Client, error) {
	// remove trailing slash (if any) from base URL
	baseURL = strings.TrimRight(baseURL, "/")

	c := &Client{
		Server:    baseURL,
		userAgent: "go-wazuh",
	}

	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	if c.ctx == nil {
		c.ctx = context.Background()
	}

	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(c.Server, "/") {
		c.Server += "/"
	}
	// create httpClient, if not already present
	c.Client = c
	c.innerClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.insecure, // test server certificate is not trusted.
			},
		},
	}

	return c, nil
}

// Do wrap the doer for tracing
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	r, e := c.innerClient.Do(req)
	if c.trace {
		var reqStr = ""
		dump, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			reqStr = strings.ReplaceAll(strings.TrimRight(string(dump), "\r\n"), "\n", "\n                            ")
		}
		if r == nil {
			dump = nil
			err = nil
		} else {
			dump, err = httputil.DumpResponse(r, true)
		}
		if err == nil {
			c.Tracef("%s\n\n                            %s\n", reqStr, strings.ReplaceAll(strings.TrimRight(string(dump), "\r\n"), "\n", "\n                            "))
		}
	}
	return r, e

}

// Errorf logs errors
func (c *Client) Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR] %s", fmt.Sprintf(format, v...))
}

// Warnf logs warings
func (c *Client) Warnf(format string, v ...interface{}) {
	log.Printf("[WARN] %s", fmt.Sprintf(format, v...))
}

// Debugf logs debug info
func (c *Client) Debugf(format string, v ...interface{}) {
	log.Printf("[DEBUG] %s", fmt.Sprintf(format, v...))
}

// Tracef logs trace info
func (c *Client) Tracef(format string, v ...interface{}) {
	log.Printf("[TRACE] %s", fmt.Sprintf(format, v...))
}

// RawAPIResponse generic response wrapper
type RawAPIResponse interface {
	Status() string
	StatusCode() int
}

func getResponseObject(sr RawAPIResponse) (interface{}, error) {
	fldForCode := fmt.Sprintf("JSON%d", sr.StatusCode())
	v := reflect.ValueOf(sr).Elem()
	if _, ok := v.Type().FieldByName(fldForCode); ok {
		s := v.FieldByName(fldForCode).Interface()
		if apiError, ok := s.(*ApiError); ok && (apiError.ApiCode != nil && *apiError.ApiCode != 0) {
			return nil, apiError
		} else if requestError, ok := s.(*RequestError); ok && (requestError.RequestError != nil && *requestError.RequestError != 0) {
			return nil, requestError
		} else if apiResponse, ok := s.(*ApiResponse); ok && (apiResponse.Code() > 0) {
			return nil, apiResponse
		} else {
			v := reflect.ValueOf(s).Elem()
			if _, ok := v.Type().FieldByName("Data"); ok {
				d := v.FieldByName("Data").Interface()
				return d, nil
			}
			if _, ok := s.(ApiResponse); ok {
				fmt.Println("do")
			}
			return s, nil
		}
	}
	return sr, nil
}

//Authenticate login using basic auth to optain a token
func (c *ClientWithResponses) Authenticate() error {
	// Authenticate
	c.ClientInterface.(*Client).token = ""
	var raw Raw = false
	params := &SecurityControllerLoginUserParams{Raw: &raw}
	sr, err := c.SecurityControllerLoginUserWithResponse(c.ClientInterface.(*Client).ctx, params)
	if err != nil {
		return err
	}
	if sr == nil {
		return fmt.Errorf("Authentication failed")
	}
	if sr.StatusCode() > 399 {
		if sr != nil {
			_, err = getResponseObject(sr)
		}
		if err != nil {
			return err
		}
		return errors.New(sr.Status())
	}

	if sr.JSON200.Data.Token == nil {
		return errors.New("Nil token!?")
	}
	c.ClientInterface.(*Client).token = *sr.JSON200.Data.Token
	return nil
}

func (c *ClientWithResponses) evaluateResponse(response RawAPIResponse, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	// log.Printf("[TRACE] %s  %v", response.Request.URL, reflect.ValueOf(response.Request.Result).Elem().FieldByName("Data").Interface())

	return getResponseObject(response)
}
