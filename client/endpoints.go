package client

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/travior/lol-sdk/types"
)

func (c *Client) GetAccountByRiotID(ctx context.Context, gameName string, tagLine string, region types.Region) (*types.Account, error) {
	routing, err := accountRouting(region)
	if err != nil {
		return nil, err
	}
	path := "/riot/account/v1/accounts/by-riot-id/" + url.PathEscape(gameName) + "/" + url.PathEscape(tagLine)
	return get[*types.Account](ctx, c, routing, path, nil)
}

func (c *Client) GetAccountByPUUID(ctx context.Context, puuid string, region types.Region) (*types.Account, error) {
	routing, err := accountRouting(region)
	if err != nil {
		return nil, err
	}
	return get[*types.Account](ctx, c, routing, "/riot/account/v1/accounts/by-puuid/"+url.PathEscape(puuid), nil)
}

func (c *Client) GetSummonerByPUUID(ctx context.Context, puuid string, region types.Region) (*types.Summoner, error) {
	routing, err := platformRouting(region)
	if err != nil {
		return nil, err
	}
	return get[*types.Summoner](ctx, c, routing, "/lol/summoner/v4/summoners/by-puuid/"+url.PathEscape(puuid), nil)
}

// MatchHistoryOptions filters and pages match history. Zero values leave a
// filter unset.
type MatchHistoryOptions struct {
	Start int
	// Count is the number of match IDs to return, up to 100. Riot defaults to 20.
	Count int
	// Queue filters by queue ID, e.g. 420 for ranked solo/duo.
	Queue int
	Type  types.MatchType
	// StartTime only works for matches played after June 16th, 2021.
	StartTime time.Time
	EndTime   time.Time
}

func (o MatchHistoryOptions) query() url.Values {
	query := url.Values{}
	if o.Start > 0 {
		query.Set("start", strconv.Itoa(o.Start))
	}
	if o.Count > 0 {
		query.Set("count", strconv.Itoa(o.Count))
	}
	if o.Queue > 0 {
		query.Set("queue", strconv.Itoa(o.Queue))
	}
	if o.Type != "" {
		query.Set("type", string(o.Type))
	}
	if !o.StartTime.IsZero() {
		query.Set("startTime", strconv.FormatInt(o.StartTime.Unix(), 10))
	}
	if !o.EndTime.IsZero() {
		query.Set("endTime", strconv.FormatInt(o.EndTime.Unix(), 10))
	}
	return query
}

// GetMatchHistoryByPUUID returns match IDs for a player, most recent first.
func (c *Client) GetMatchHistoryByPUUID(ctx context.Context, puuid string, region types.Region, opts MatchHistoryOptions) ([]string, error) {
	routing, err := matchRouting(region)
	if err != nil {
		return nil, err
	}
	return get[[]string](ctx, c, routing, "/lol/match/v5/matches/by-puuid/"+url.PathEscape(puuid)+"/ids", opts.query())
}

func (c *Client) GetMatch(ctx context.Context, matchID string, region types.Region) (*types.Match, error) {
	routing, err := matchRouting(region)
	if err != nil {
		return nil, err
	}
	return get[*types.Match](ctx, c, routing, "/lol/match/v5/matches/"+url.PathEscape(matchID), nil)
}

func (c *Client) GetMatchTimeline(ctx context.Context, matchID string, region types.Region) (*types.MatchTimeline, error) {
	routing, err := matchRouting(region)
	if err != nil {
		return nil, err
	}
	return get[*types.MatchTimeline](ctx, c, routing, "/lol/match/v5/matches/"+url.PathEscape(matchID)+"/timeline", nil)
}

func (c *Client) GetChallengerLeague(ctx context.Context, queue types.Queue, region types.Region) (*types.LeagueList, error) {
	return c.getApexLeague(ctx, "challengerleagues", queue, region)
}

func (c *Client) GetGrandmasterLeague(ctx context.Context, queue types.Queue, region types.Region) (*types.LeagueList, error) {
	return c.getApexLeague(ctx, "grandmasterleagues", queue, region)
}

func (c *Client) GetMasterLeague(ctx context.Context, queue types.Queue, region types.Region) (*types.LeagueList, error) {
	return c.getApexLeague(ctx, "masterleagues", queue, region)
}

func (c *Client) getApexLeague(ctx context.Context, league string, queue types.Queue, region types.Region) (*types.LeagueList, error) {
	routing, err := platformRouting(region)
	if err != nil {
		return nil, err
	}
	return get[*types.LeagueList](ctx, c, routing, "/lol/league/v4/"+league+"/by-queue/"+url.PathEscape(string(queue)), nil)
}

// GetLeagueEntries returns one page of players in a tier and division below
// Master. Pages start at 1; a page of 0 is treated as 1.
func (c *Client) GetLeagueEntries(ctx context.Context, queue types.Queue, tier types.Tier, division types.Division, page int, region types.Region) ([]types.LeagueEntry, error) {
	routing, err := platformRouting(region)
	if err != nil {
		return nil, err
	}
	path := "/lol/league/v4/entries/" + url.PathEscape(string(queue)) + "/" + url.PathEscape(string(tier)) + "/" + url.PathEscape(string(division))
	var query url.Values
	if page > 1 {
		query = url.Values{"page": {strconv.Itoa(page)}}
	}
	return get[[]types.LeagueEntry](ctx, c, routing, path, query)
}

func (c *Client) GetLeagueEntriesByPUUID(ctx context.Context, puuid string, region types.Region) ([]types.LeagueEntry, error) {
	routing, err := platformRouting(region)
	if err != nil {
		return nil, err
	}
	return get[[]types.LeagueEntry](ctx, c, routing, "/lol/league/v4/entries/by-puuid/"+url.PathEscape(puuid), nil)
}
