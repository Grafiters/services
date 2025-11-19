package lib

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// MakeRequest is a reusable function for making HTTP requests
func MakeRequest(method, url string, headers map[string]string, requestBody interface{}, responseBody interface{}) error {
	client := &http.Client{}

	// Marshal the request body into JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal the request data: %v", err)
	}

	fmt.Println("Url =>", url)

	// Create a new request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create the request: %v", err)
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make the request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the status code is 200 (OK)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %v", err)
	}

	fmt.Println("response body =>", string(body))

	err = json.Unmarshal(body, responseBody)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the response body: %v", err)
	}

	return nil
}

// For BrigateRequest
type AuthType int

const (
	NoAuth AuthType = iota
	BasicAuth
)

type RequestConfig struct {
	Method      string
	Url         string
	Headers     map[string]string
	Payload     interface{}
	ContentType string
	Auth        AuthType
	Username    string
	Password    string
}

func BuildRequest(cfg RequestConfig) (*http.Request, error) {
	var body io.Reader

	switch cfg.ContentType {
	case "application/json":
		jsonData, err := json.Marshal(cfg.Payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal the request data: %v", err)
		}
		body = bytes.NewBuffer(jsonData)

	case "application/x-www-form-urlencoded":
		formData, ok := cfg.Payload.(url.Values)
		if !ok {
			return nil, errors.New("payload must be url.Values for form content")
		}

		body = strings.NewReader(formData.Encode())

	default:
		body = nil
	}

	req, err := http.NewRequest(cfg.Method, cfg.Url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create the request: %v", err)
	}

	for k, v := range cfg.Headers {
		req.Header.Set(k, v)
	}

	// Set content-type
	if cfg.ContentType != "" {
		req.Header.Set("Content-Type", cfg.ContentType)
	}

	// Add Basic Auth if required
	if cfg.Auth == BasicAuth {
		authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(cfg.Username+":"+cfg.Password))
		req.Header.Set("Authorization", authHeader)
	}

	return req, nil
}

func SendRequest(client *http.Client, req *http.Request, result interface{}) error {
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// // Check for non-2xx
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("received non-success status code: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("response body =>", string(body))

	if err := json.Unmarshal(body, result); err != nil {
		return err
	}

	return nil
}
