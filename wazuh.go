package wazuh

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
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

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn

	ctx       context.Context
	userAgent string
	token     string
	user      string
	password  string
	insecure  bool
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
func WithInsecure() ClientOption {
	return func(c *Client) error {
		c.insecure = true
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
	c, err := NewAPIClient(baseURL, opts...)
	if err != nil {
		return nil, err
	}
	if os.Getenv("WAZUH_INSECURE") == "true" {
		err := WithInsecure()(c.ClientInterface.(*Client))
		if err != nil {
			return nil, err
		}
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
	if c.Client == nil {
		c.Client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: c.insecure, // test server certificate is not trusted.
				},
			},
		}
	}

	return c, nil
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
		if apiError, ok := s.(*ApiError); ok {
			return nil, fmt.Errorf("%d: %s (%s)", sr.StatusCode(), apiError.Title, apiError.Detail)
		}
		if requestError, ok := s.(*RequestError); ok {
			return nil, fmt.Errorf("%d: %s (%s)", sr.StatusCode(), requestError.Title, requestError.Detail)
		}
		v := reflect.ValueOf(s).Elem()
		if _, ok := v.Type().FieldByName("Data"); ok {
			d := v.FieldByName("Data").Interface()
			return d, nil
		}
		return v, nil
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

	if sr.StatusCode() > 399 {
		_, err := getResponseObject(sr)
		if err != nil {
			return err
		}
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
