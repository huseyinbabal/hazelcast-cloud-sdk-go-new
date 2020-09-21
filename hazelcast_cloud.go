// Package hazelcastcloud is the Hazelcast Cloud API client for Go.
package hazelcastcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/utils/to"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	// This is the version of the Library.
	libraryVersion = "1.0.0"
	// This is the default base url.
	defaultBaseURL = "https://cloud.hazelcast.com/api/v1"
	// This is the User-Agent of Client.
	userAgent = "hazelcast-cloud-sdk-go/" + libraryVersion
	// This is the Http Accept type.
	mediaType = "application/json"
	// This is the header key of total Rate Limit.
	headerRateLimit = "X-RateLimit-Limit"
	// This is the header key of total Rate Remaining.
	headerRateRemaining = "X-RateLimit-Remaining"
	// This is the header key of total Rate Reset.
	headerRateReset = "X-RateLimit-Reset"
	// This is status code of Rate Limit.
	statusCodeRateLimit = 429
)

// This is the RequestCompletionCallback to intercept response when request complete.
type RequestCompletionCallback func(*http.Request, *http.Response)

// This is the main Client.
type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string
	Rate      Rate
	rateMutex sync.Mutex
	token     string

	StarterCluster    StarterClusterService
	EnterpriseCluster EnterpriseClusterService
	CloudProvider     CloudProviderService
	Region            RegionService
	AvailabilityZone  AvailabilityZoneService
	InstanceType      InstanceTypeService
	HazelcastVersion  HazelcastVersionService
	Auth              AuthService

	onRequestCompleted RequestCompletionCallback
}

// This is the OnRequestCompleted method to assign callback.
func (c *Client) OnRequestCompleted(rc RequestCompletionCallback) {
	c.onRequestCompleted = rc
}

type Option func(*Client)

func OptionEndpoint(e string) func(*Client) {
	return func(c *Client) {
		baseURL, err := url.Parse(e)
		if err != nil {
			panic(err)
		}
		c.BaseURL = baseURL
	}
}

// This creates a new client with provided http client.
func NewClient(httpClient *http.Client, options ...Option) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}

	c.Auth = NewAuthService(c)
	c.StarterCluster = NewStarterClusterService(c)
	c.EnterpriseCluster = NewEnterpriseClusterService(c)
	c.CloudProvider = NewCloudProviderService(c)
	c.Region = NewRegionService(c)
	c.AvailabilityZone = NewAvailabilityZoneService(c)
	c.InstanceType = NewInstanceTypeService(c)
	c.HazelcastVersion = NewHazelcastVersionService(c)

	for _, option := range options {
		option(c)
	}

	return c
}

// This function creates a new client.
func New(options ...Option) (*Client, *Response, error) {
	apiKey := os.Getenv("HZ_CLOUD_API_KEY")
	apiSecret := os.Getenv("HZ_CLOUD_API_SECRET")
	if len(strings.TrimSpace(apiKey)) == 0 || len(strings.TrimSpace(apiSecret)) == 0 {
		return nil, nil, &ErrorResponse{
			CorrelationId: "",
			Message:       "You need to provide HZ_CLOUD_API_KEY and HZ_CLOUD_API_SECRET in your environment variables.",
			Response:      nil,
		}
	}
	return NewFromCredentials(apiKey, apiSecret, options...)
}

// This function create new client with ApiKey and ApiSecret.
func NewFromCredentials(apiKey string, apiSecret string, options ...Option) (*Client, *Response, error) {
	client := NewClient(nil, options...)
	login, loginResp, loginErr := client.Auth.Login(context.Background(), &models.LoginInput{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	})

	if loginErr != nil {
		return nil, loginResp, loginErr
	}
	client.token = login.Token
	return client, loginResp, nil
}

//This function creates a new http request from graphql request. It returns an error if it can not build request.
func (c *Client) NewRequest(body *models.GraphqlRequest) (*http.Request, error) {
	graphqlBody := GraphQLQuery{
		OperationName: "",
		Query:         to.Query(body.Name, body.Operation, body.Input, body.Args, body.Response),
		Variables:     map[string]interface{}{"input": to.Variables(body.Input)},
	}

	buf := new(bytes.Buffer)
	encodeErr := json.NewEncoder(buf).Encode(graphqlBody)
	if encodeErr != nil {
		return nil, encodeErr
	}

	req, requestErr := http.NewRequest(http.MethodPost, c.BaseURL.String(), buf)
	if requestErr != nil {
		return nil, requestErr
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)
	if c.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	return req, nil
}

//This function sends http request to the server. Then it creates a response according to type interface provided as v.
//It returns error if response is not successful.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, respErr := DoRequestWithClient(ctx, c.client, req)

	if respErr != nil {
		return nil, respErr
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer resp.Body.Close()

	response := newResponse(resp)
	c.rateMutex.Lock()
	c.Rate = response.Rate
	c.rateMutex.Unlock()

	responseData, checkRespErr := AugmentResponse(resp)
	if checkRespErr != nil {
		return response, checkRespErr
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, copyErr := io.Copy(w, resp.Body)
			if copyErr != nil {
				return response, copyErr
			}
		} else {
			var objectMap map[string]interface{}
			if err := json.Unmarshal(responseData, &objectMap); err != nil {
				log.Fatal(err)
			}
			dataMarshall, _ := json.Marshal(objectMap["data"].(map[string]interface{})["response"])
			decodeErr := json.NewDecoder(bytes.NewReader(dataMarshall)).Decode(v)
			if decodeErr != nil {
				return response, decodeErr
			}
		}
	}

	return response, nil
}

// This function augments the response and returns and error if response has
func AugmentResponse(response *http.Response) ([]byte, error) {
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == statusCodeRateLimit {
		return nil, &ErrorResponse{
			CorrelationId: response.Header.Get("X-B3-TraceId"),
			Message:       "Too Many Request",
			Response:      response,
		}
	}

	var objectMap map[string]interface{}
	if err := json.Unmarshal(responseData, &objectMap); err != nil {
		log.Fatal(err)
	}

	errorObject, errorKeyFound := objectMap["errors"]
	if !errorKeyFound {
		return responseData, nil
	}

	errorObjectJson, errorObjectJsonErr := json.Marshal(errorObject)
	if errorObjectJsonErr != nil {
		log.Fatal(errorObjectJsonErr)
	}

	var errorResponse []ErrorResponse
	if errorUnmarshal := json.Unmarshal(errorObjectJson, &errorResponse); errorUnmarshal != nil {
		log.Fatal(errorUnmarshal)
	}

	firstError := &errorResponse[0]
	firstError.Response = response
	firstError.CorrelationId = response.Header.Get("X-B3-TraceId")
	return nil, firstError
}

//This method returns an Error string of ErrorResponse.
func (r *ErrorResponse) Error() string {
	if r.CorrelationId != "" {
		return fmt.Sprintf("Method:%v URL:%v: Status:%d TraceId:%s Message:%v",
			r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.CorrelationId, r.Message)
	}
	return fmt.Sprintf("Method:%v URL:%v: Status:%d Message:%v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

//Error response is the main type of response for the errors this library handles
type ErrorResponse struct {
	CorrelationId string `json:"correlation_id,omitempty"`
	Message       string `json:"message"`
	Response      *http.Response
}

//This function creates a new response with populating rate from http response.
func newResponse(r *http.Response) *Response {
	response := Response{Response: r}
	response.populateRate()

	return &response
}

//This functions populates the rates from http headers
func (r *Response) populateRate() {
	if limit := r.Header.Get(headerRateLimit); limit != "" {
		r.Rate.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		r.Rate.Remaining, _ = strconv.Atoi(remaining)
	}
	if reset := r.Header.Get(headerRateReset); reset != "" {
		resetTime, _ := strconv.ParseInt(reset, 10, 64)
		r.Rate.Reset = resetTime
	}
}

//This function does request with the request and client provided.
func DoRequestWithClient(
	ctx context.Context,
	client *http.Client,
	req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return client.Do(req)
}

//Type of the main response
type Response struct {
	*http.Response
	Rate
}

//Type of the rate limiting params we follow
type Rate struct {
	Limit     int   `json:"limit"`
	Remaining int   `json:"remaining"`
	Reset     int64 `json:"reset"`
}

//Type of graphql query
type GraphQLQuery struct {
	OperationName string                 `json:"operationName,omitempty"`
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}
