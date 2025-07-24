# League of Legends SDK

A Go SDK for interacting with the Riot Games League of Legends API. This library provides a simple and efficient way to fetch summoner data, match history, match details, and league information.

## Features

- **Rate Limiting**: Built-in per-region rate limiting to comply with Riot API limits
- **Regional Support**: Support for all League of Legends regions
- **Comprehensive Data Types**: Full type definitions for matches, summoners, leagues, and timelines
- **Structured Logging**: Uses zerolog for detailed request logging
- **Context Support**: All API calls support Go context for cancellation and timeouts

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
    
    client := client.NewClient(client.Config{
        APIKey:         "YOUR_RIOT_API_KEY",
        RequestsPerMin: 100,
        BurstSize:      20,
    }, &logger)

    ctx := context.Background()
    
    // Get summoner by PUUID
    summoner, err := client.GetSummonerByPUUID(ctx, "puuid", types.EUW1)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Summoner: %s (Level %d)\n", summoner.Name, summoner.SummonerLevel)
}
```

## Supported Regions

- **Europe**: EUW1, EUN1, TR1, RU
- **Asia**: KR, JP1  
- **Americas**: BR1, LA1, LA2, OC1, NA1

## API Methods

### Summoner API
- `GetSummonerByPUUID(ctx, puuid, region)` - Get summoner information by PUUID

### Match API
- `GetMatchHistoryByPUUID(ctx, puuid, region, count)` - Get match history for a summoner
- `GetMatch(ctx, matchID, region)` - Get detailed match information
- `GetMatchTimeline(ctx, matchID, region)` - Get match timeline data

### League API
- `GetChallengerLeague(ctx, queue, region)` - Get Challenger tier players
- `GetGrandMasterLeague(ctx, queue, region)` - Get Grandmaster tier players  
- `GetMasterLeague(ctx, queue, region)` - Get Master tier players
- `GetLeagueEntries(ctx, queue, tier, division, region)` - Get players in specific tier/division

## Configuration

The client accepts a `Config` struct with the following options:

- `APIKey`: Your Riot Games API key (required)
- `RequestsPerMin`: Rate limit per minute
- `BurstSize`: Burst size for rate limiting

## Testing

Set your API key as an environment variable and run the tests:

```bash
export API_KEY=your_riot_api_key_here go test -v
```

## Legal

This SDK is not affiliated with, endorsed, sponsored, or specifically approved by Riot Games and Riot Games is not responsible for it. This SDK uses the Riot Games API but is not endorsed or certified by Riot Games.

## License

This project is licensed under the MIT License.

**Legal Notice**: League of Legends® and Riot Games® are registered trademarks of Riot Games, Inc. This project is not affiliated with Riot Games.
