# League of Legends SDK

A Go SDK for interacting with the Riot Games League of Legends API. This library provides a simple and efficient way to fetch summoner data, match history, match details, and league information.

## Features

- **Rate Limiting**: Built-in per-region rate limiting to comply with Riot API limits
- **Regional Support**: Support for all League of Legends regions
- **Comprehensive Data Types**: Full type definitions for matches, summoners, leagues, and timelines
- **Structured Logging**: Optional zerolog request logging
- **Typed Errors**: Non-200 responses are returned as `*client.APIError`
- **Context Support**: All API calls support Go context for cancellation and timeouts

## Requirements

Go 1.26 or newer.

## Installation

```bash
go get github.com/travior/lol-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/rs/zerolog"
    "github.com/travior/lol-sdk/client"
    "github.com/travior/lol-sdk/types"
)

func main() {
    logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
    
    riot := client.NewClient(client.Config{
        APIKey:         "YOUR_RIOT_API_KEY",
        RequestsPerMin: 50,
    }, &logger)

    ctx := context.Background()
    
    // Look up a player by Riot ID (gameName#tagLine)
    account, err := riot.GetAccountByRiotID(ctx, "gameName", "tagLine", types.EUW1)
    if err != nil {
        panic(err)
    }

    summoner, err := riot.GetSummonerByPUUID(ctx, account.PUUID, types.EUW1)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Summoner: %s#%s (Level %d)\n", account.GameName, account.TagLine, summoner.SummonerLevel)
}
```

## Supported Regions

- **Europe**: EUW1, EUN1, TR1, RU, ME1
- **Asia**: KR, JP1
- **Americas**: BR1, LA1, LA2, NA1
- **SEA**: OC1, SG2, TW2, VN2

All regions are listed in `types.Regions`. The zero value of `types.Region` is not valid, and requests made with it return `client.ErrUnknownRegion`.

## API Methods

### Account API
- `GetAccountByRiotID(ctx, gameName, tagLine, region)` - Get an account (PUUID) by Riot ID
- `GetAccountByPUUID(ctx, puuid, region)` - Get an account's Riot ID by PUUID

### Summoner API
- `GetSummonerByPUUID(ctx, puuid, region)` - Get summoner information by PUUID

### Match API
- `GetMatchHistoryByPUUID(ctx, puuid, region, opts)` - Get match IDs for a player. `MatchHistoryOptions` supports paging (`Start`, `Count`) and filtering by `Queue`, `Type`, `StartTime` and `EndTime`
- `GetMatch(ctx, matchID, region)` - Get detailed match information
- `GetMatchTimeline(ctx, matchID, region)` - Get match timeline data

### League API
- `GetChallengerLeague(ctx, queue, region)` - Get Challenger tier players
- `GetGrandmasterLeague(ctx, queue, region)` - Get Grandmaster tier players
- `GetMasterLeague(ctx, queue, region)` - Get Master tier players
- `GetLeagueEntries(ctx, queue, tier, division, page, region)` - Get one page of players in a tier/division below Master
- `GetLeagueEntriesByPUUID(ctx, puuid, region)` - Get a player's ranked entries

Queues, tiers and divisions are typed, e.g. `types.RankedSolo5x5`, `types.Diamond` and `types.DivisionI`.

## Error Handling

When Riot responds with a non-200 status, the error is a `*client.APIError` with the status code, Riot's error message and, for 429 responses, the `Retry-After` duration:

```go
match, err := riot.GetMatch(ctx, matchID, types.EUW1)
var apiErr *client.APIError
if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
    // match does not exist
}
```

## Configuration

`NewClient(config, logger)` accepts a `Config` struct and an optional `*zerolog.Logger`. Pass `nil` to disable logging; requests are logged at debug level.

- `APIKey`: Your Riot Games API key (required)
- `RequestsPerMin`: Requests per minute for each routing value (defaults to 50, a development key's limit)
- `BurstSize`: Burst size for rate limiting (defaults to 1)

## Testing

`go test ./...` runs the offline tests against a local fake server. To also run the tests against the live API, set your API key:

```bash
API_KEY=your_riot_api_key_here go test -v ./...
```

## Legal

This SDK is not affiliated with, endorsed, sponsored, or specifically approved by Riot Games and Riot Games is not responsible for it. This SDK uses the Riot Games API but is not endorsed or certified by Riot Games.

## License

This project is licensed under the MIT License.

**Legal Notice**: League of Legends® and Riot Games® are registered trademarks of Riot Games, Inc. This project is not affiliated with Riot Games.
