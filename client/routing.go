package client

import (
	"fmt"

	"github.com/travior/lol-sdk/types"
)

// platformRouting returns the routing value for platform endpoints such as
// summoner-v4 and league-v4.
func platformRouting(region types.Region) (string, error) {
	if !region.Valid() {
		return "", fmt.Errorf("%w: %q", ErrUnknownRegion, string(region))
	}
	return region.String(), nil
}

// accountRouting returns the regional routing value for account-v1, which
// only supports americas, asia and europe.
func accountRouting(region types.Region) (string, error) {
	switch region {
	case types.BR1, types.LA1, types.LA2, types.NA1, types.OC1:
		return "americas", nil
	case types.KR, types.JP1, types.SG2, types.TW2, types.VN2:
		return "asia", nil
	case types.EUW1, types.EUN1, types.TR1, types.RU, types.ME1:
		return "europe", nil
	}
	return "", fmt.Errorf("%w: %q", ErrUnknownRegion, string(region))
}

// matchRouting returns the regional routing value for match-v5.
func matchRouting(region types.Region) (string, error) {
	switch region {
	case types.BR1, types.LA1, types.LA2, types.NA1:
		return "americas", nil
	case types.KR, types.JP1:
		return "asia", nil
	case types.EUW1, types.EUN1, types.TR1, types.RU, types.ME1:
		return "europe", nil
	case types.OC1, types.SG2, types.TW2, types.VN2:
		return "sea", nil
	}
	return "", fmt.Errorf("%w: %q", ErrUnknownRegion, string(region))
}
