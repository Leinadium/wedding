package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// apiClient holds the underlying http.Client and base URL.
type apiClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewAPIClient creates a new client. It defaults to http.DefaultClient if nil is passed.
func NewAPIClient(baseURL string, client *http.Client) *apiClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &apiClient{
		BaseURL:    baseURL,
		HTTPClient: client,
	}
}

// doRequest is the internal helper that executes the HTTP request.
// reqBody is marshaled to JSON if provided. resBody is populated from the JSON response if provided.
func (c *apiClient) doRequest(ctx context.Context, method, path string, headers map[string]string, reqBody any, resBody any) error {
	var bodyReader io.Reader

	// 1. Marshal Request Payload (if any)
	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	// 2. Create the Request
	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// 3. Set Default and Custom Headers
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 4. Execute the Request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	// 5. Handle HTTP Errors
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Read the body to surface API error messages
		errBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(errBody))
	}

	// 6. Unmarshal Response Payload (if requested and content exists)
	if resBody != nil && resp.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(resp.Body).Decode(resBody); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// Get performs an HTTP GET request.
// Type parameter Res represents the expected JSON response struct.
func Get[Res any](c *apiClient, ctx context.Context, path string, headers map[string]string) (*Res, error) {
	var res Res
	err := c.doRequest(ctx, http.MethodGet, path, headers, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Post performs an HTTP POST request.
// Type Req is the payload struct, Res is the expected response struct.
func Post[Req any, Res any](c *apiClient, ctx context.Context, path string, headers map[string]string, payload Req) (*Res, error) {
	var res Res
	err := c.doRequest(ctx, http.MethodPost, path, headers, payload, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Patch performs an HTTP PATCH request.
func Patch[Req any, Res any](c *apiClient, ctx context.Context, path string, headers map[string]string, payload Req) (*Res, error) {
	var res Res
	err := c.doRequest(ctx, http.MethodPatch, path, headers, payload, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Delete performs an HTTP DELETE request.
func Delete(c *apiClient, ctx context.Context, path string, headers map[string]string) error {
	err := c.doRequest(ctx, http.MethodDelete, path, headers, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
