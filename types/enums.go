package types

import (
	"fmt"
	"slices"
	"strings"
)

// Region is a League of Legends platform, e.g. "euw1". The zero value is not a
// valid region.
type Region string

const (
	//europe
	EUW1 Region = "euw1"
	EUN1 Region = "eun1"
	TR1  Region = "tr1"
	RU   Region = "ru"
	ME1  Region = "me1"

	//asia
	KR  Region = "kr"
	JP1 Region = "jp1"

	//americas
	BR1 Region = "br1"
	LA1 Region = "la1"
	LA2 Region = "la2"
	NA1 Region = "na1"

	//sea
	OC1 Region = "oc1"
	SG2 Region = "sg2"
	TW2 Region = "tw2"
	VN2 Region = "vn2"
)

// Regions lists every supported region.
var Regions = []Region{
	EUW1, EUN1, TR1, RU, ME1,
	KR, JP1,
	BR1, LA1, LA2, NA1,
	OC1, SG2, TW2, VN2,
}

func (r Region) Valid() bool {
	return slices.Contains(Regions, r)
}

func (r Region) String() string {
	return string(r)
}

func (r Region) MarshalText() ([]byte, error) {
	if !r.Valid() {
		return nil, fmt.Errorf("unknown region: %q", string(r))
	}
	return []byte(r), nil
}

func (r *Region) UnmarshalText(text []byte) error {
	region := Region(strings.ToLower(string(text)))
	if !region.Valid() {
		return fmt.Errorf("unknown region: %q", string(text))
	}
	*r = region
	return nil
}

// Queue is a ranked queue type as used by the league endpoints.
type Queue string

const (
	RankedSolo5x5 Queue = "RANKED_SOLO_5x5"
	RankedFlexSR  Queue = "RANKED_FLEX_SR"
)

type Tier string

const (
	Iron        Tier = "IRON"
	Bronze      Tier = "BRONZE"
	Silver      Tier = "SILVER"
	Gold        Tier = "GOLD"
	Platinum    Tier = "PLATINUM"
	Emerald     Tier = "EMERALD"
	Diamond     Tier = "DIAMOND"
	Master      Tier = "MASTER"
	Grandmaster Tier = "GRANDMASTER"
	Challenger  Tier = "CHALLENGER"
)

type Division string

const (
	DivisionI   Division = "I"
	DivisionII  Division = "II"
	DivisionIII Division = "III"
	DivisionIV  Division = "IV"
)

// MatchType filters match history by the kind of game.
type MatchType string

const (
	MatchTypeRanked   MatchType = "ranked"
	MatchTypeNormal   MatchType = "normal"
	MatchTypeTourney  MatchType = "tourney"
	MatchTypeTutorial MatchType = "tutorial"
)
