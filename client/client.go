// Copyright 2023 CodeMaker AI Inc. All rights reserved.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultConnectionTimeout = 5 * time.Second
	defaultRequestTimeout    = 50 * time.Second

	endpointUrl = "https://api.codemaker.ai"

	headerAuthorization = "Authorization"
	headerRequestId     = "X-Request-Id"
)

type Client interface {
	CreateProcess(request *CreateProcessRequest) (*CreateProcessResponse, error)
	GetProcessStatus(request *GetProcessStatusRequest) (*GetProcessStatusResponse, error)
	GetProcessOutput(request *GetProcessOutputRequest) (*GetProcessOutputResponse, error)
}

type HttpClient struct {
	Client
	config Config
	client *http.Client
}

func NewClient(config Config) Client {
	connectionTimeout := defaultConnectionTimeout
	if config.ConnectionTimeout != nil && *config.ConnectionTimeout > 0 {
		connectionTimeout = *config.ConnectionTimeout
	}

	requestTimeout := defaultRequestTimeout
	if config.RequestTimeout != nil && *config.RequestTimeout > 0 {
		requestTimeout = *config.RequestTimeout
	}

	return &HttpClient{
		config: config,
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: connectionTimeout,
				}).DialContext,
				TLSHandshakeTimeout: connectionTimeout,
			},
			Timeout: requestTimeout,
		},
	}
}

func (c *HttpClient) CreateProcess(request *CreateProcessRequest) (*CreateProcessResponse, error) {
	if request == nil {
		request = &CreateProcessRequest{}
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, NewClientErrorWithCause("failed to serialize request payload", err)
	}

	resp, err := c.doRequest(http.MethodPost, "/process", body)
	if err != nil {
		return nil, NewClientErrorWithCause("failed to make HTTP request", err)
	}
	defer resp.Body.Close()

	if !c.isSuccess(resp) {
		return nil, c.handleError(resp)
	}

	response, err := c.handleResponse(resp, &CreateProcessResponse{})
	if err != nil {
		return nil, err
	}
	return response.(*CreateProcessResponse), nil
}

func (c *HttpClient) GetProcessStatus(request *GetProcessStatusRequest) (*GetProcessStatusResponse, error) {
	if request == nil {
		request = &GetProcessStatusRequest{}
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, NewClientErrorWithCause("failed to serialize request payload", err)
	}

	resp, err := c.doRequest(http.MethodPost, "/process/status", body)
	if err != nil {
		return nil, NewClientErrorWithCause("failed to make HTTP request", err)
	}
	defer resp.Body.Close()

	if !c.isSuccess(resp) {
		return nil, c.handleError(resp)
	}

	response, err := c.handleResponse(resp, &GetProcessStatusResponse{})
	if err != nil {
		return nil, err
	}
	return response.(*GetProcessStatusResponse), nil
}

func (c *HttpClient) GetProcessOutput(request *GetProcessOutputRequest) (*GetProcessOutputResponse, error) {
	if request == nil {
		request = &GetProcessOutputRequest{}
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, NewClientErrorWithCause("failed to serialize request payload", err)
	}

	resp, err := c.doRequest(http.MethodPost, "/process/output", body)
	if err != nil {
		return nil, NewClientErrorWithCause("failed to make HTTP request", err)
	}
	defer resp.Body.Close()

	if !c.isSuccess(resp) {
		return nil, c.handleError(resp)
	}

	response, err := c.handleResponse(resp, &GetProcessOutputResponse{})
	if err != nil {
		return nil, err
	}
	return response.(*GetProcessOutputResponse), nil
}

func (c *HttpClient) isSuccess(resp *http.Response) bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

func (c *HttpClient) doRequest(method string, path string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, c.url(path), bytes.NewBuffer(body))
	if err != nil {
		return nil, NewClientErrorWithCause("failed to create HTTP request", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", fmt.Sprintf("CodeMakerSdkGo/%s", Version))
	req.Header.Add(headerAuthorization, fmt.Sprintf("Bearer %s", c.config.ApiKey))

	resp, err := c.client.Do(req)
	return resp, err
}

func (c *HttpClient) handleResponse(resp *http.Response, val interface{}) (interface{}, error) {
	reader, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, NewClientErrorWithCause("failed to read HTTP response", err)
	}

	err = json.Unmarshal(reader, val)
	if err != nil {
		return nil, NewClientErrorWithCause("failed to parse HTTP response", err)
	}
	return val, nil
}

func (c *HttpClient) handleError(resp *http.Response) error {
	var requestId = ""
	if id, ok := resp.Header[headerRequestId]; ok {
		requestId = id[0]
	}
	errorCode := c.tryUnmarshallError(resp)

	return NewClientError(
		fmt.Sprintf("(%s) request failed %d %s", requestId, resp.StatusCode, errorCode))
}

func (c *HttpClient) tryUnmarshallError(resp *http.Response) string {
	reader, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	result := &Error{}
	err = json.Unmarshal(reader, result)
	if err != nil {
		return ""
	}
	return result.Code
}

func (c *HttpClient) endpoint() string {
	if c.config.Endpoint != nil {
		return *c.config.Endpoint
	}
	return endpointUrl
}

func (c *HttpClient) url(path string) string {
	return fmt.Sprintf("%s/%s",
		strings.TrimSuffix(c.endpoint(), "/"),
		strings.TrimPrefix(path, "/"))
}
