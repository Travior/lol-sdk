package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

// defaultRequestsPerMin matches a development key's limit of 100 requests per 2 minutes.
const defaultRequestsPerMin = 50

type Client struct {
	httpClient       *http.Client
	logger           *zerolog.Logger
	rateLimiters     map[string]*rate.Limiter
	rateLimiterMutex sync.Mutex
	config           Config
}

type Config struct {
	APIKey string
	// RequestsPerMin is the request rate allowed per routing value. Defaults to 50.
	RequestsPerMin int
	// BurstSize is the number of requests that may be sent at once. Defaults to 1.
	BurstSize int
}

// NewClient creates a client. The logger is optional; pass nil to disable logging.
func NewClient(config Config, logger *zerolog.Logger) *Client {
	if logger == nil {
		nop := zerolog.Nop()
		logger = &nop
	}
	if config.RequestsPerMin <= 0 {
		config.RequestsPerMin = defaultRequestsPerMin
	}
	if config.BurstSize <= 0 {
		config.BurstSize = 1
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		logger:       logger,
		rateLimiters: make(map[string]*rate.Limiter),
		config:       config,
	}
}

func (c *Client) getRateLimiter(routingValue string) *rate.Limiter {
	c.rateLimiterMutex.Lock()
	defer c.rateLimiterMutex.Unlock()

	limiter, exists := c.rateLimiters[routingValue]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(float64(c.config.RequestsPerMin)/60.0), c.config.BurstSize)
		c.rateLimiters[routingValue] = limiter
		c.logger.Debug().Str("routing_value", routingValue).Msg("Created new limiter")
	}

	return limiter
}

func (c *Client) makeRequest(ctx context.Context, requestURL string, routingValue string) ([]byte, error) {
	if err := c.getRateLimiter(routingValue).Wait(ctx); err != nil {
		return nil, fmt.Errorf("waiting for rate limiter: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("X-Riot-Token", c.config.APIKey)
	req.Header.Set("User-Agent", "lol-sdk")

	start := time.Now()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request to %s: %w", requestURL, err)
	}
	defer resp.Body.Close()

	c.logger.Debug().
		Str("routing_value", routingValue).
		Str("url", requestURL).
		Int("status", resp.StatusCode).
		Dur("duration", time.Since(start)).
		Msg("Riot API request")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response from %s: %w", requestURL, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, newAPIError(requestURL, resp, body)
	}

	return body, nil
}

// get requests path from the given routing value and decodes the JSON response into T.
func get[T any](ctx context.Context, c *Client, routingValue string, path string, query url.Values) (T, error) {
	var result T

	requestURL := "https://" + routingValue + ".api.riotgames.com" + path
	if len(query) > 0 {
		requestURL += "?" + query.Encode()
	}

	body, err := c.makeRequest(ctx, requestURL, routingValue)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("decoding response from %s: %w", requestURL, err)
	}

	return result, nil
}
