package client

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/travior/lol-sdk/types"
)

func setupLiveClient(t *testing.T) *Client {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		t.Skip("No API Key set")
	}

	logger := zerolog.New(zerolog.NewTestWriter(t)).
		With().
		Timestamp().
		Str("test", t.Name()).
		Logger().
		Level(zerolog.WarnLevel)

	return NewClient(Config{
		APIKey:         apiKey,
		RequestsPerMin: 60,
		BurstSize:      1,
	},
		&logger)
}

func TestGetChallenger(t *testing.T) {
	client := setupLiveClient(t)

	for _, region := range types.Regions {
		t.Run(
			fmt.Sprintf("TestGetChallenger-%s", region),
			func(t *testing.T) {
				ctx := context.Background()
				league, err := client.GetChallengerLeague(ctx, types.RankedSolo5x5, region)
				if err != nil {
					t.Fatalf("API call failed: %v", err)
				}
				t.Logf("Fetched %d players", len(league.Entries))
			})
	}
}

func TestGetGrandmaster(t *testing.T) {
	client := setupLiveClient(t)

	for _, region := range types.Regions {
		t.Run(
			fmt.Sprintf("TestGetGrandmaster-%s", region),
			func(t *testing.T) {
				ctx := context.Background()
				league, err := client.GetGrandmasterLeague(ctx, types.RankedSolo5x5, region)
				if err != nil {
					t.Fatalf("API call failed: %v", err)
				}
				t.Logf("Fetched %d players", len(league.Entries))
			})
	}
}

func TestGetMaster(t *testing.T) {
	client := setupLiveClient(t)

	for _, region := range types.Regions {
		t.Run(
			fmt.Sprintf("TestGetMaster-%s", region),
			func(t *testing.T) {
				ctx := context.Background()
				league, err := client.GetMasterLeague(ctx, types.RankedSolo5x5, region)
				if err != nil {
					t.Fatalf("API call failed: %v", err)
				}
				t.Logf("Fetched %d players", len(league.Entries))
			})
	}
}

func TestGetDiamondI(t *testing.T) {
	client := setupLiveClient(t)

	for _, region := range types.Regions {
		t.Run(
			fmt.Sprintf("TestGetDiamondI-%s", region),
			func(t *testing.T) {
				ctx := context.Background()
				entries, err := client.GetLeagueEntries(ctx, types.RankedSolo5x5, types.Diamond, types.DivisionI, 1, region)
				if err != nil {
					t.Fatalf("API call failed: %v", err)
				}
				t.Logf("Fetched %d players", len(entries))
			})
	}
}

func TestMatchHistoryAndData(t *testing.T) {
	client := setupLiveClient(t)

	for _, region := range types.Regions {
		t.Run(
			fmt.Sprintf("TestMatchHistoryAndData-%s", region),
			func(t *testing.T) {
				ctx := context.Background()

				// Get challenger players first
				league, err := client.GetChallengerLeague(ctx, types.RankedSolo5x5, region)
				if err != nil {
					t.Fatalf("Failed to get challenger league: %v", err)
				}

				if len(league.Entries) == 0 {
					t.Skip("No challenger players found")
				}

				// Get first player's PUUID
				playerPUUID := league.Entries[0].PUUID
				if playerPUUID == "" {
					t.Skip("Player PUUID not available")
				}

				// Get match history
				matches, err := client.GetMatchHistoryByPUUID(ctx, playerPUUID, region, MatchHistoryOptions{Count: 5})
				if err != nil {
					t.Fatalf("Failed to get match history: %v", err)
				}

				if len(matches) == 0 {
					t.Skip("No matches found")
				}

				t.Logf("Fetched %d matches", len(matches))

				// Get first match details
				match, err := client.GetMatch(ctx, matches[0], region)
				if err != nil {
					t.Fatalf("Failed to get match: %v", err)
				}

				t.Logf("Match ID: %s, Duration: %d seconds", match.Metadata.MatchID, match.Info.GameDuration)

				// Get match timeline
				timeline, err := client.GetMatchTimeline(ctx, matches[0], region)
				if err != nil {
					t.Fatalf("Failed to get match timeline: %v", err)
				}

				t.Logf("Timeline has %d frames", len(timeline.Info.Frames))
			})
	}
}

func TestPlayerLookups(t *testing.T) {
	client := setupLiveClient(t)

	for _, region := range types.Regions {
		t.Run(
			fmt.Sprintf("TestPlayerLookups-%s", region),
			func(t *testing.T) {
				ctx := context.Background()

				league, err := client.GetChallengerLeague(ctx, types.RankedSolo5x5, region)
				if err != nil {
					t.Fatalf("Failed to get challenger league: %v", err)
				}
				if len(league.Entries) == 0 {
					t.Skip("No challenger players found")
				}
				playerPUUID := league.Entries[0].PUUID

				summoner, err := client.GetSummonerByPUUID(ctx, playerPUUID, region)
				if err != nil {
					t.Fatalf("Failed to get summoner: %v", err)
				}
				t.Logf("Summoner level %d", summoner.SummonerLevel)

				entries, err := client.GetLeagueEntriesByPUUID(ctx, playerPUUID, region)
				if err != nil {
					t.Fatalf("Failed to get league entries by puuid: %v", err)
				}
				t.Logf("Player has %d league entries", len(entries))

				account, err := client.GetAccountByPUUID(ctx, playerPUUID, region)
				if err != nil {
					t.Fatalf("Failed to get account by puuid: %v", err)
				}
				if account.GameName == "" {
					t.Skip("Account has no game name")
				}

				byRiotID, err := client.GetAccountByRiotID(ctx, account.GameName, account.TagLine, region)
				if err != nil {
					t.Fatalf("Failed to get account by riot id: %v", err)
				}
				if byRiotID.PUUID != playerPUUID {
					t.Errorf("Riot ID lookup returned a different PUUID")
				}
			})
	}
}
