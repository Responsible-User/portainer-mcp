package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// doJSONAPIRequest performs an API request and decodes the JSON response into the target.
// If the response status code is not in the 2xx range, it returns an error with the response body.
func (c *PortainerClient) doJSONAPIRequest(method, path string, body io.Reader, target any) error {
	resp, err := c.DoAPIRequest(method, path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	if target != nil {
		if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
			return fmt.Errorf("failed to decode API response: %w", err)
		}
	}

	return nil
}

// doAPIDelete performs a DELETE request and checks for a successful status code.
func (c *PortainerClient) doAPIDelete(path string) error {
	resp, err := c.DoAPIRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}
