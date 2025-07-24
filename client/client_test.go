package client

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/travior/lol-sdk/types"
)

func setupClient(t *testing.T) *Client {
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

var regions = []types.Region{
	types.EUW1,
	types.EUN1,
	types.TR1,
	types.RU,
	types.KR,
	types.JP1,
	types.BR1,
	types.LA1,
	types.LA2,
	types.OC1,
	types.NA1,
}

func TestGetChallenger(t *testing.T) {
	client := setupClient(t)

	for _, region := range regions {
		t.Run(
			fmt.Sprintf("TestGetChallenger-%s", region.ToString()),
			func(t *testing.T) {
				ctx := context.Background()
				league, err := client.GetChallengerLeague(ctx, "RANKED_SOLO_5x5", region)
				if err != nil {
					t.Fatalf("API call failed: %v", err)
				}
				t.Logf("Fetched %d players", len(league.Entries))
			})
	}
}

func TestGetGrandmaster(t *testing.T) {
	client := setupClient(t)

	for _, region := range regions {
		t.Run(
			fmt.Sprintf("TestGetGrandmaster-%s", region.ToString()),
			func(t *testing.T) {
				ctx := context.Background()
				league, err := client.GetGrandMasterLeague(ctx, "RANKED_SOLO_5x5", region)
				if err != nil {
					t.Fatalf("API call failed: %v", err)
				}
				t.Logf("Fetched %d players", len(league.Entries))
			})
	}
}

func TestGetMaster(t *testing.T) {
	client := setupClient(t)

	for _, region := range regions {
		t.Run(
			fmt.Sprintf("TestGetMaster-%s", region.ToString()),
			func(t *testing.T) {
				ctx := context.Background()
				league, err := client.GetMasterLeague(ctx, "RANKED_SOLO_5x5", region)
				if err != nil {
					t.Fatalf("API call failed: %v", err)
				}
				t.Logf("Fetched %d players", len(league.Entries))
			})
	}
}

func TestGetDiamondI(t *testing.T) {
	client := setupClient(t)

	for _, region := range regions {
		t.Run(
			fmt.Sprintf("TestGetDiamondI-%s", region.ToString()),
			func(t *testing.T) {
				ctx := context.Background()
				entries, err := client.GetLeagueEntries(ctx, "RANKED_SOLO_5x5", "DIAMOND", "I", region)
				if err != nil {
					t.Fatalf("API call failed: %v", err)
				}
				t.Logf("Fetched %d players", len(entries))
			})
	}
}

func TestMatchHistoryAndData(t *testing.T) {
	client := setupClient(t)

	for _, region := range regions {
		t.Run(
			fmt.Sprintf("TestMatchHistoryAndData-%s", region.ToString()),
			func(t *testing.T) {
				ctx := context.Background()

				// Get challenger players first
				league, err := client.GetChallengerLeague(ctx, "RANKED_SOLO_5x5", region)
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
				matches, err := client.GetMatchHistoryByPUUID(ctx, playerPUUID, region, 5)
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
