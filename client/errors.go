package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// ErrUnknownRegion is returned when a request is made with an invalid region.
var ErrUnknownRegion = errors.New("unknown region")

// APIError is returned when the Riot API responds with a non-200 status code.
// Use errors.As to inspect it.
type APIError struct {
	URL        string
	StatusCode int
	// Message is the error message from Riot's response body, if there was one.
	Message string
	Body    string
	// RetryAfter is set from the Retry-After header on 429 responses.
	RetryAfter time.Duration
}

func (e *APIError) Error() string {
	message := e.Message
	if message == "" {
		message = http.StatusText(e.StatusCode)
	}
	return fmt.Sprintf("riot api: %s returned %d: %s", e.URL, e.StatusCode, message)
}

func newAPIError(requestURL string, resp *http.Response, body []byte) *APIError {
	apiErr := &APIError{
		URL:        requestURL,
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}

	var parsed struct {
		Status struct {
			Message string `json:"message"`
		} `json:"status"`
	}
	if json.Unmarshal(body, &parsed) == nil {
		apiErr.Message = parsed.Status.Message
	}

	if seconds, err := strconv.Atoi(resp.Header.Get("Retry-After")); err == nil {
		apiErr.RetryAfter = time.Duration(seconds) * time.Second
	}

	return apiErr
}
