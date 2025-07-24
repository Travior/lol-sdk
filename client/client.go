package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/time/rate"

	"github.com/travior/lol-sdk/types"
)

type Client struct {
	httpClient       *http.Client
	logger           *zerolog.Logger
	rateLimiters     map[string]*rate.Limiter
	rateLimiterMutex sync.RWMutex
	config           Config
}

type Config struct {
	APIKey         string
	RequestsPerMin int
	BurstSize      int
}

func NewClient(config Config, logger *zerolog.Logger) *Client {
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
	c.rateLimiterMutex.RLock()
	limiter, exists := c.rateLimiters[routingValue]
	c.rateLimiterMutex.RUnlock()

	if exists {
		return limiter
	}

	c.rateLimiterMutex.Lock()
	defer c.rateLimiterMutex.Unlock()

	if limiter, exists := c.rateLimiters[routingValue]; exists {
		return limiter
	}
	c.rateLimiters[routingValue] = rate.NewLimiter(rate.Limit(float64(c.config.RequestsPerMin)/60.0), 1)
	c.logger.Info().Str("routing_value", routingValue).Msg("Created new limiter")

	return c.rateLimiters[routingValue]
}

func (c *Client) makeRequest(ctx context.Context, url string, routingValue string) ([]byte, error) {
	limiter := c.getRateLimiter(routingValue)
	if err := limiter.Wait(ctx); err != nil {
		c.logger.Err(err).Str("routing_value", routingValue).Str("url", url).Msg("rate limiter wait failed")
		return nil, fmt.Errorf("rate limiter wait failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		c.logger.Err(err).Str("routing_value", routingValue).Str("url", url).Msg("failed to create request")
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Riot-Token", c.config.APIKey)
	req.Header.Set("User-Agent", "lol-sdk/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Err(err).Str("routing_value", routingValue).Str("url", url).Msg("failed to make request")
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Err(err).Str("routing_value", routingValue).Str("url", url).Msg("failed to read response body")
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		c.logger.Warn().Str("routing_value", routingValue).Str("url", url).Int("status", resp.StatusCode).Msg("Got non OK status code")
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func getAccountRouting(region types.Region) string {
	switch region {
	case types.BR1, types.LA1, types.LA2, types.OC1, types.NA1:
		return "americas"
	case types.KR, types.JP1:
		return "asia"
	case types.EUN1, types.EUW1, types.TR1, types.RU:
		return "europe"
	}
	return "" //switch is exhaustive
}

func (c *Client) GetSummonerByPUUID(ctx context.Context, puuid string, region types.Region) (*types.Summoner, error) {
	c.logger.Debug().Str("puuid", puuid).Str("region", region.ToString()).Msg("Fetching summoner")

	routingValue := getAccountRouting(region)
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s", routingValue, puuid)

	body, err := c.makeRequest(ctx, url, routingValue)
	if err != nil {
		c.logger.Err(err).Str("puuid", puuid).Str("region", region.ToString()).Msg("Failed to fetch summoner")
		return nil, err
	}

	var summoner types.Summoner
	if err := json.Unmarshal(body, &summoner); err != nil {
		c.logger.Err(err).Str("puuid", puuid).Str("region", region.ToString()).Msg("Failed to parse response")
		return nil, err
	}

	return &summoner, nil
}

func (c *Client) GetMatchHistoryByPUUID(ctx context.Context, puuid string, region types.Region, count int) ([]string, error) {
	c.logger.Debug().Str("puuid", puuid).Str("region", region.ToString()).Msg("Fetching match history")

	routingValue := getAccountRouting(region)
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=%d", routingValue, puuid, count)

	body, err := c.makeRequest(ctx, url, routingValue)
	if err != nil {
		c.logger.Err(err).Str("puuid", puuid).Str("region", region.ToString()).Msg("Failed to fetch match history")
		return nil, err
	}

	var matches []string
	if err := json.Unmarshal(body, &matches); err != nil {
		c.logger.Err(err).Str("puuid", puuid).Str("region", region.ToString()).Msg("Failed to parse response")
		return nil, err
	}

	c.logger.Debug().Str("puuid", puuid).Str("region", region.ToString()).Int("match_count", len(matches)).Interface("matches", matches).Msg("Match history fetched")
	return matches, err

}

func (c *Client) GetMatch(ctx context.Context, matchID string, region types.Region) (*types.Match, error) {
	c.logger.Debug().Str("matchID", matchID).Str("region", region.ToString()).Msg("Fetching match")

	routingValue := getAccountRouting(region)
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/%s", routingValue, matchID)

	body, err := c.makeRequest(ctx, url, routingValue)
	if err != nil {
		c.logger.Err(err).Str("matchID", matchID).Str("region", region.ToString()).Msg("Failed to fetch match")
		return nil, err
	}

	var match types.Match
	if err := json.Unmarshal(body, &match); err != nil {
		c.logger.Err(err).Str("matchID", matchID).Str("region", region.ToString()).Msg("Failed to parse match")
		return nil, err
	}

	return &match, nil
}

func (c *Client) GetMatchTimeline(ctx context.Context, matchID string, region types.Region) (*types.MatchTimeline, error) {
	c.logger.Debug().Str("matchID", matchID).Str("region", region.ToString()).Msg("Fetching match timeline")

	routingValue := getAccountRouting(region)
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/%s/timeline", routingValue, matchID)

	body, err := c.makeRequest(ctx, url, routingValue)
	if err != nil {
		c.logger.Err(err).Str("matchID", matchID).Str("region", region.ToString()).Msg("Failed to fetch match timeline")
		return nil, err
	}

	var timeline types.MatchTimeline
	if err := json.Unmarshal(body, &timeline); err != nil {
		c.logger.Err(err).Str("matchID", matchID).Str("region", region.ToString()).Msg("Failed to parse match timeline")
		return nil, err
	}

	return &timeline, nil
}

func (c *Client) GetChallengerLeague(ctx context.Context, queue string, region types.Region) (*types.LeagueList, error) {
	c.logger.Debug().Str("queue", queue).Str("region", region.ToString()).Msg("Fetching challengers")

	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/league/v4/challengerleagues/by-queue/%s", region.ToString(), queue)

	body, err := c.makeRequest(ctx, url, region.ToString())
	if err != nil {
		c.logger.Err(err).Str("queue", queue).Str("region", region.ToString()).Msg("Failed to get challenger league")
		return nil, err
	}

	var league types.LeagueList
	if err := json.Unmarshal(body, &league); err != nil {
		c.logger.Err(err).Str("queue", queue).Str("region", region.ToString()).Msg("Failed to parse challenger league response")
		return nil, err
	}

	return &league, nil
}

func (c *Client) GetGrandMasterLeague(ctx context.Context, queue string, region types.Region) (*types.LeagueList, error) {
	c.logger.Debug().Str("queue", queue).Str("region", region.ToString()).Msg("Fetching grandmasters")

	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/league/v4/grandmasterleagues/by-queue/%s", region.ToString(), queue)

	body, err := c.makeRequest(ctx, url, region.ToString())
	if err != nil {
		c.logger.Err(err).Str("queue", queue).Str("region", region.ToString()).Msg("Failed to get grandmaster league")
		return nil, err
	}

	var league types.LeagueList
	if err := json.Unmarshal(body, &league); err != nil {
		c.logger.Err(err).Str("queue", queue).Str("region", region.ToString()).Msg("Failed to parse grandmaster league response")
		return nil, err
	}

	return &league, nil
}

func (c *Client) GetMasterLeague(ctx context.Context, queue string, region types.Region) (*types.LeagueList, error) {
	c.logger.Debug().Str("queue", queue).Str("region", region.ToString()).Msg("Fetching masters")

	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/league/v4/masterleagues/by-queue/%s", region.ToString(), queue)

	body, err := c.makeRequest(ctx, url, region.ToString())
	if err != nil {
		c.logger.Err(err).Str("queue", queue).Str("region", region.ToString()).Msg("Failed to get master league")
		return nil, err
	}

	var league types.LeagueList
	if err := json.Unmarshal(body, &league); err != nil {
		c.logger.Err(err).Str("queue", queue).Str("region", region.ToString()).Msg("Failed to parse master league response")
		return nil, err
	}

	return &league, nil
}

func (c *Client) GetLeagueEntries(ctx context.Context, queue string, tier string, division string, region types.Region) ([]types.LeagueEntry, error) {
	c.logger.Debug().Str("queue", queue).Str("tier", tier).Str("division", division).Str("region", region.ToString()).Msg("Fetching league entries")

	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/league/v4/entries/%s/%s/%s", region.ToString(), queue, tier, division)

	body, err := c.makeRequest(ctx, url, region.ToString())
	if err != nil {
		c.logger.Err(err).Str("queue", queue).Str("tier", tier).Str("division", division).Str("region", region.ToString()).Msg("Failed to get league entries")
		return nil, err
	}

	var entries []types.LeagueEntry
	if err := json.Unmarshal(body, &entries); err != nil {
		c.logger.Err(err).Str("queue", queue).Str("tier", tier).Str("division", division).Str("region", region.ToString()).Msg("Failed to parse league entries response")
		return nil, err
	}

	return entries, nil
}
