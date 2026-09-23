package client

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/travior/lol-sdk/types"
)

// rewriteTransport sends every request to target, keeping the original host
// in the X-Original-Host header.
type rewriteTransport struct {
	target *url.URL
}

func (t rewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.Header.Set("X-Original-Host", req.URL.Host)
	req.URL.Scheme = t.target.Scheme
	req.URL.Host = t.target.Host
	return http.DefaultTransport.RoundTrip(req)
}

func setupFakeClient(t *testing.T, handler http.HandlerFunc) *Client {
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	target, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	client := NewClient(Config{APIKey: "test-key", RequestsPerMin: 6000, BurstSize: 10}, nil)
	client.httpClient.Transport = rewriteTransport{target: target}
	return client
}

func TestRouting(t *testing.T) {
	matchRoutingByRegion := map[types.Region]string{
		types.EUW1: "europe", types.EUN1: "europe", types.TR1: "europe", types.RU: "europe", types.ME1: "europe",
		types.KR: "asia", types.JP1: "asia",
		types.BR1: "americas", types.LA1: "americas", types.LA2: "americas", types.NA1: "americas",
		types.OC1: "sea", types.SG2: "sea", types.TW2: "sea", types.VN2: "sea",
	}
	validAccountRouting := map[string]bool{"americas": true, "asia": true, "europe": true}

	if len(matchRoutingByRegion) != len(types.Regions) {
		t.Fatalf("routing table covers %d regions, expected %d", len(matchRoutingByRegion), len(types.Regions))
	}
	for _, region := range types.Regions {
		if got, err := matchRouting(region); err != nil || got != matchRoutingByRegion[region] {
			t.Errorf("matchRouting(%s) = %q, %v; want %q", region, got, err, matchRoutingByRegion[region])
		}
		if got, err := accountRouting(region); err != nil || !validAccountRouting[got] {
			t.Errorf("accountRouting(%s) = %q, %v; not a valid account-v1 routing value", region, got, err)
		}
		if got, err := platformRouting(region); err != nil || got != string(region) {
			t.Errorf("platformRouting(%s) = %q, %v", region, got, err)
		}
	}

	for _, routing := range []func(types.Region) (string, error){matchRouting, accountRouting, platformRouting} {
		if _, err := routing(""); !errors.Is(err, ErrUnknownRegion) {
			t.Errorf("expected ErrUnknownRegion for empty region, got %v", err)
		}
	}
}

func TestRegionText(t *testing.T) {
	for _, region := range types.Regions {
		text, err := region.MarshalText()
		if err != nil {
			t.Fatalf("MarshalText(%s) failed: %v", region, err)
		}
		var parsed types.Region
		if err := parsed.UnmarshalText(text); err != nil {
			t.Fatalf("UnmarshalText(%q) failed: %v", text, err)
		}
		if parsed != region {
			t.Errorf("round trip of %s gave %s", region, parsed)
		}
	}

	var parsed types.Region
	if err := parsed.UnmarshalText([]byte("EUW1")); err != nil || parsed != types.EUW1 {
		t.Errorf("UnmarshalText should be case-insensitive, got %q, %v", parsed, err)
	}
	if err := parsed.UnmarshalText([]byte("nowhere")); err == nil {
		t.Error("expected an error for an unknown region")
	}
	if _, err := types.Region("").MarshalText(); err == nil {
		t.Error("expected an error when marshalling the zero value")
	}
}

func TestRequest(t *testing.T) {
	var gotRequest *http.Request
	client := setupFakeClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotRequest = r
		w.Write([]byte(`["EUW1_1", "EUW1_2"]`))
	})

	start := time.Unix(1700000000, 0)
	matches, err := client.GetMatchHistoryByPUUID(context.Background(), "some/puuid", types.EUW1, MatchHistoryOptions{
		Count:     5,
		Queue:     420,
		Type:      types.MatchTypeRanked,
		StartTime: start,
	})
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if len(matches) != 2 || matches[0] != "EUW1_1" {
		t.Errorf("unexpected matches: %v", matches)
	}
	if host := gotRequest.Header.Get("X-Original-Host"); host != "europe.api.riotgames.com" {
		t.Errorf("request went to %q", host)
	}
	if key := gotRequest.Header.Get("X-Riot-Token"); key != "test-key" {
		t.Errorf("API key header = %q", key)
	}
	if path := gotRequest.URL.EscapedPath(); path != "/lol/match/v5/matches/by-puuid/some%2Fpuuid/ids" {
		t.Errorf("path = %q", path)
	}
	wantQuery := url.Values{"count": {"5"}, "queue": {"420"}, "type": {"ranked"}, "startTime": {"1700000000"}}
	if query := gotRequest.URL.Query(); query.Encode() != wantQuery.Encode() {
		t.Errorf("query = %q, want %q", query.Encode(), wantQuery.Encode())
	}
}

func TestAPIError(t *testing.T) {
	client := setupFakeClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "7")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"status":{"message":"Rate limit exceeded","status_code":429}}`))
	})

	_, err := client.GetSummonerByPUUID(context.Background(), "puuid", types.KR)

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected an APIError, got %v", err)
	}
	if apiErr.StatusCode != http.StatusTooManyRequests {
		t.Errorf("StatusCode = %d", apiErr.StatusCode)
	}
	if apiErr.Message != "Rate limit exceeded" {
		t.Errorf("Message = %q", apiErr.Message)
	}
	if apiErr.RetryAfter != 7*time.Second {
		t.Errorf("RetryAfter = %s", apiErr.RetryAfter)
	}
}

func TestUnknownRegionSkipsRequest(t *testing.T) {
	client := setupFakeClient(t, func(w http.ResponseWriter, r *http.Request) {
		t.Error("no request should be sent for an unknown region")
	})

	if _, err := client.GetMatch(context.Background(), "EUW1_1", ""); !errors.Is(err, ErrUnknownRegion) {
		t.Errorf("expected ErrUnknownRegion, got %v", err)
	}
}
