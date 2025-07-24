package types

import (
	"fmt"
	"strings"
)

type Region int

const (
	//europe
	EUW1 Region = iota
	EUN1
	TR1
	RU

	//asia
	KR
	JP1

	//america
	BR1
	LA1
	LA2
	OC1
	NA1
)

func (r *Region) UnmarshalText(text []byte) error {
	name := string(text)
	switch n := strings.ToLower(name); n {
	case "euw1":
		*r = EUW1
	case "eun1":
		*r = EUN1
	case "tr1":
		*r = TR1
	case "ru":
		*r = RU
	case "kr":
		*r = KR
	case "jp1":
		*r = JP1
	case "br1":
		*r = BR1
	case "LA1":
		*r = LA1
	case "LA2":
		*r = LA2
	case "OC1":
		*r = OC1
	case "NA1":
		*r = NA1
	default:
		return fmt.Errorf("Unknown region: %s", name)
	}
	return nil
}

func (r Region) ToString() string {
	switch r {
	case EUW1:
		return "euw1"
	case EUN1:
		return "eun1"
	case TR1:
		return "tr1"
	case RU:
		return "ru"
	case KR:
		return "kr"
	case JP1:
		return "jp1"
	case BR1:
		return "br1"
	case LA1:
		return "la1"
	case LA2:
		return "la2"
	case OC1:
		return "oc1"
	case NA1:
		return "na1"
	}
	return ""
}

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}
type Match struct {
	Metadata MatchMetadata `json:"metadata"`
	Info     MatchInfo     `json:"info"`
}

type MatchMetadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type MatchInfo struct {
	GameCreation       int64         `json:"gameCreation"`
	GameDuration       int           `json:"gameDuration"`
	GameEndTimestamp   int64         `json:"gameEndTimestamp"`
	GameID             int64         `json:"gameId"`
	GameMode           string        `json:"gameMode"`
	GameName           string        `json:"gameName"`
	GameStartTimestamp int64         `json:"gameStartTimestamp"`
	GameType           string        `json:"gameType"`
	GameVersion        string        `json:"gameVersion"`
	MapID              int           `json:"mapId"`
	Participants       []Participant `json:"participants"`
	PlatformID         string        `json:"platformId"`
	QueueID            int           `json:"queueId"`
	Teams              []Team        `json:"teams"`
	TournamentCode     string        `json:"tournamentCode"`
}

type Participant struct {
	AllInPings                     int              `json:"allInPings"`
	AssistMePings                  int              `json:"assistMePings"`
	Assists                        int              `json:"assists"`
	BaronKills                     int              `json:"baronKills"`
	BountyLevel                    int              `json:"bountyLevel"`
	ChampExperience                int              `json:"champExperience"`
	ChampLevel                     int              `json:"champLevel"`
	ChampionID                     int              `json:"championId"`
	ChampionName                   string           `json:"championName"`
	ChampionTransform              int              `json:"championTransform"`
	ConsumablesPurchased           int              `json:"consumablesPurchased"`
	DamageDealtToBuildings         int              `json:"damageDealtToBuildings"`
	DamageDealtToObjectives        int              `json:"damageDealtToObjectives"`
	DamageDealtToTurrets           int              `json:"damageDealtToTurrets"`
	DamageSelfMitigated            int              `json:"damageSelfMitigated"`
	Deaths                         int              `json:"deaths"`
	DetectorWardsPlaced            int              `json:"detectorWardsPlaced"`
	DoubleKills                    int              `json:"doubleKills"`
	DragonKills                    int              `json:"dragonKills"`
	EligibleForProgression         bool             `json:"eligibleForProgression"`
	EnemyMissingPings              int              `json:"enemyMissingPings"`
	EnemyVisionPings               int              `json:"enemyVisionPings"`
	FirstBloodAssist               bool             `json:"firstBloodAssist"`
	FirstBloodKill                 bool             `json:"firstBloodKill"`
	FirstTowerAssist               bool             `json:"firstTowerAssist"`
	FirstTowerKill                 bool             `json:"firstTowerKill"`
	GameEndedInEarlySurrender      bool             `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender           bool             `json:"gameEndedInSurrender"`
	GetBackPings                   int              `json:"getBackPings"`
	GoldEarned                     int              `json:"goldEarned"`
	GoldSpent                      int              `json:"goldSpent"`
	HoldPings                      int              `json:"holdPings"`
	IndividualPosition             string           `json:"individualPosition"`
	InhibitorKills                 int              `json:"inhibitorKills"`
	InhibitorTakedowns             int              `json:"inhibitorTakedowns"`
	InhibitorsLost                 int              `json:"inhibitorsLost"`
	Item0                          int              `json:"item0"`
	Item1                          int              `json:"item1"`
	Item2                          int              `json:"item2"`
	Item3                          int              `json:"item3"`
	Item4                          int              `json:"item4"`
	Item5                          int              `json:"item5"`
	Item6                          int              `json:"item6"`
	ItemsPurchased                 int              `json:"itemsPurchased"`
	KillingSprees                  int              `json:"killingSprees"`
	Kills                          int              `json:"kills"`
	Lane                           string           `json:"lane"`
	LargestCriticalStrike          int              `json:"largestCriticalStrike"`
	LargestKillingSpree            int              `json:"largestKillingSpree"`
	LargestMultiKill               int              `json:"largestMultiKill"`
	LongestTimeSpentLiving         int              `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int              `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int              `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int              `json:"magicDamageTaken"`
	NeedVisionPings                int              `json:"needVisionPings"`
	NeutralMinionsKilled           int              `json:"neutralMinionsKilled"`
	NexusKills                     int              `json:"nexusKills"`
	NexusLost                      int              `json:"nexusLost"`
	NexusTakedowns                 int              `json:"nexusTakedowns"`
	ObjectivesStolen               int              `json:"objectivesStolen"`
	ObjectivesStolenAssists        int              `json:"objectivesStolenAssists"`
	OnMyWayPings                   int              `json:"onMyWayPings"`
	ParticipantID                  int              `json:"participantId"`
	PentaKills                     int              `json:"pentaKills"`
	Perks                          ParticipantPerks `json:"perks"`
	PhysicalDamageDealt            int              `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int              `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int              `json:"physicalDamageTaken"`
	ProfileIcon                    int              `json:"profileIcon"`
	PushPings                      int              `json:"pushPings"`
	PUUID                          string           `json:"puuid"`
	QuadraKills                    int              `json:"quadraKills"`
	RiotIDGameName                 string           `json:"riotIdGameName"`
	RiotIDName                     string           `json:"riotIdName"`
	RiotIDTagline                  string           `json:"riotIdTagline"`
	Role                           string           `json:"role"`
	SightWardsBoughtInGame         int              `json:"sightWardsBoughtInGame"`
	Spell1Casts                    int              `json:"spell1Casts"`
	Spell2Casts                    int              `json:"spell2Casts"`
	Spell3Casts                    int              `json:"spell3Casts"`
	Spell4Casts                    int              `json:"spell4Casts"`
	Summoner1Casts                 int              `json:"summoner1Casts"`
	Summoner1ID                    int              `json:"summoner1Id"`
	Summoner2Casts                 int              `json:"summoner2Casts"`
	Summoner2ID                    int              `json:"summoner2Id"`
	SummonerID                     string           `json:"summonerId"`
	SummonerLevel                  int              `json:"summonerLevel"`
	SummonerName                   string           `json:"summonerName"`
	TeamEarlySurrendered           bool             `json:"teamEarlySurrendered"`
	TeamID                         int              `json:"teamId"`
	TeamPosition                   string           `json:"teamPosition"`
	TimeCCingOthers                int              `json:"timeCCingOthers"`
	TimePlayed                     int              `json:"timePlayed"`
	TotalDamageDealt               int              `json:"totalDamageDealt"`
	TotalDamageDealtToChampions    int              `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates int              `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken               int              `json:"totalDamageTaken"`
	TotalHeal                      int              `json:"totalHeal"`
	TotalHealsOnTeammates          int              `json:"totalHealsOnTeammates"`
	TotalMinionsKilled             int              `json:"totalMinionsKilled"`
	TotalTimeCCDealt               int              `json:"totalTimeCCDealt"`
	TotalTimeSpentDead             int              `json:"totalTimeSpentDead"`
	TotalUnitsHealed               int              `json:"totalUnitsHealed"`
	TripleKills                    int              `json:"tripleKills"`
	TrueDamageDealt                int              `json:"trueDamageDealt"`
	TrueDamageDealtToChampions     int              `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                int              `json:"trueDamageTaken"`
	TurretKills                    int              `json:"turretKills"`
	TurretTakedowns                int              `json:"turretTakedowns"`
	TurretsLost                    int              `json:"turretsLost"`
	UnrealKills                    int              `json:"unrealKills"`
	VisionClearedPings             int              `json:"visionClearedPings"`
	VisionScore                    int              `json:"visionScore"`
	VisionWardsBoughtInGame        int              `json:"visionWardsBoughtInGame"`
	WardsKilled                    int              `json:"wardsKilled"`
	WardsPlaced                    int              `json:"wardsPlaced"`
	Win                            bool             `json:"win"`
}

type ParticipantPerks struct {
	StatPerks ParticipantStatPerks `json:"statPerks"`
	Styles    []ParticipantStyle   `json:"styles"`
}

type ParticipantStatPerks struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

type ParticipantStyle struct {
	Description string                      `json:"description"`
	Selections  []ParticipantStyleSelection `json:"selections"`
	Style       int                         `json:"style"`
}

type ParticipantStyleSelection struct {
	Perk int `json:"perk"`
	Var1 int `json:"var1"`
	Var2 int `json:"var2"`
	Var3 int `json:"var3"`
}

type Team struct {
	Bans       []TeamBan      `json:"bans"`
	Objectives TeamObjectives `json:"objectives"`
	TeamID     int            `json:"teamId"`
	Win        bool           `json:"win"`
}

type TeamBan struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type TeamObjectives struct {
	Baron      TeamObjective `json:"baron"`
	Champion   TeamObjective `json:"champion"`
	Dragon     TeamObjective `json:"dragon"`
	Inhibitor  TeamObjective `json:"inhibitor"`
	RiftHerald TeamObjective `json:"riftHerald"`
	Tower      TeamObjective `json:"tower"`
}

type TeamObjective struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}

type MatchTimeline struct {
	Metadata TimelineMetadata `json:"metadata"`
	Info     TimelineInfo     `json:"info"`
}

type TimelineMetadata struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type TimelineInfo struct {
	FrameInterval int                   `json:"frameInterval"`
	Frames        []TimelineFrame       `json:"frames"`
	GameID        int64                 `json:"gameId"`
	Participants  []TimelineParticipant `json:"participants"`
}

type TimelineFrame struct {
	Events            []TimelineEvent                     `json:"events"`
	ParticipantFrames map[string]TimelineParticipantFrame `json:"participantFrames"`
	Timestamp         int                                 `json:"timestamp"`
}

type TimelineEvent struct {
	RealTimestamp           int64                  `json:"realTimestamp"`
	Timestamp               int                    `json:"timestamp"`
	Type                    string                 `json:"type"`
	ItemID                  int                    `json:"itemId,omitempty"`
	ParticipantID           int                    `json:"participantId,omitempty"`
	LevelUpType             string                 `json:"levelUpType,omitempty"`
	SkillSlot               int                    `json:"skillSlot,omitempty"`
	CreatorID               int                    `json:"creatorId,omitempty"`
	WardType                string                 `json:"wardType,omitempty"`
	Level                   int                    `json:"level,omitempty"`
	AssistingParticipantIDs []int                  `json:"assistingParticipantIds,omitempty"`
	BountyLevel             int                    `json:"bountyLevel,omitempty"`
	KillStreakLength        int                    `json:"killStreakLength,omitempty"`
	KillerID                int                    `json:"killerId,omitempty"`
	Position                TimelinePosition       `json:"position"`
	VictimDamageDealt       []TimelineVictimDamage `json:"victimDamageDealt,omitempty"`
	VictimDamageReceived    []TimelineVictimDamage `json:"victimDamageReceived,omitempty"`
	VictimID                int                    `json:"victimId,omitempty"`
	KillType                string                 `json:"killType,omitempty"`
	LaneType                string                 `json:"laneType,omitempty"`
	TeamID                  int                    `json:"teamId,omitempty"`
	MonsterType             string                 `json:"monsterType,omitempty"`
	MonsterSubType          string                 `json:"monsterSubType,omitempty"`
	BuildingType            string                 `json:"buildingType,omitempty"`
	TowerType               string                 `json:"towerType,omitempty"`
	AfterID                 int                    `json:"afterId,omitempty"`
	BeforeID                int                    `json:"beforeId,omitempty"`
	GoldGain                int                    `json:"goldGain,omitempty"`
	GameID                  int64                  `json:"gameId,omitempty"`
	WinningTeam             int                    `json:"winningTeam,omitempty"`
}

type TimelinePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type TimelineVictimDamage struct {
	Basic          bool   `json:"basic"`
	MagicDamage    int    `json:"magicDamage"`
	Name           string `json:"name"`
	ParticipantID  int    `json:"participantId"`
	PhysicalDamage int    `json:"physicalDamage"`
	SpellName      string `json:"spellName"`
	SpellSlot      int    `json:"spellSlot"`
	TrueDamage     int    `json:"trueDamage"`
	Type           string `json:"type"`
}

type TimelineParticipantFrame struct {
	ChampionStats            TimelineChampionStats `json:"championStats"`
	CurrentGold              int                   `json:"currentGold"`
	DamageStats              TimelineDamageStats   `json:"damageStats"`
	GoldPerSecond            int                   `json:"goldPerSecond"`
	JungleMinionsKilled      int                   `json:"jungleMinionsKilled"`
	Level                    int                   `json:"level"`
	MinionsKilled            int                   `json:"minionsKilled"`
	ParticipantID            int                   `json:"participantId"`
	Position                 TimelinePosition      `json:"position"`
	TimeEnemySpentControlled int                   `json:"timeEnemySpentControlled"`
	TotalGold                int                   `json:"totalGold"`
	XP                       int                   `json:"xp"`
}

type TimelineChampionStats struct {
	AbilityHaste         int `json:"abilityHaste"`
	AbilityPower         int `json:"abilityPower"`
	Armor                int `json:"armor"`
	ArmorPen             int `json:"armorPen"`
	ArmorPenPercent      int `json:"armorPenPercent"`
	AttackDamage         int `json:"attackDamage"`
	AttackSpeed          int `json:"attackSpeed"`
	BonusArmorPenPercent int `json:"bonusArmorPenPercent"`
	BonusMagicPenPercent int `json:"bonusMagicPenPercent"`
	CCReduction          int `json:"ccReduction"`
	CooldownReduction    int `json:"cooldownReduction"`
	Health               int `json:"health"`
	HealthMax            int `json:"healthMax"`
	HealthRegen          int `json:"healthRegen"`
	Lifesteal            int `json:"lifesteal"`
	MagicPen             int `json:"magicPen"`
	MagicPenPercent      int `json:"magicPenPercent"`
	MagicResist          int `json:"magicResist"`
	MovementSpeed        int `json:"movementSpeed"`
	Omnivamp             int `json:"omnivamp"`
	PhysicalVamp         int `json:"physicalVamp"`
	Power                int `json:"power"`
	PowerMax             int `json:"powerMax"`
	PowerRegen           int `json:"powerRegen"`
	SpellVamp            int `json:"spellVamp"`
}

type TimelineDamageStats struct {
	MagicDamageDone               int `json:"magicDamageDone"`
	MagicDamageDoneToChampions    int `json:"magicDamageDoneToChampions"`
	MagicDamageTaken              int `json:"magicDamageTaken"`
	PhysicalDamageDone            int `json:"physicalDamageDone"`
	PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
	PhysicalDamageTaken           int `json:"physicalDamageTaken"`
	TotalDamageDone               int `json:"totalDamageDone"`
	TotalDamageDoneToChampions    int `json:"totalDamageDoneToChampions"`
	TotalDamageTaken              int `json:"totalDamageTaken"`
	TrueDamageDone                int `json:"trueDamageDone"`
	TrueDamageDoneToChampions     int `json:"trueDamageDoneToChampions"`
	TrueDamageTaken               int `json:"trueDamageTaken"`
}

type TimelineParticipant struct {
	ParticipantID int    `json:"participantId"`
	PUUID         string `json:"puuid"`
}

type LeagueList struct {
	LeagueID string        `json:"leagueId"`
	Entries  []LeagueEntry `json:"entries"`
	Tier     string        `json:"tier"`
	Name     string        `json:"name"`
	Queue    string        `json:"queue"`
}

type LeagueEntry struct {
	SummonerID   string     `json:"summonerId"`
	SummonerName string     `json:"summonerName"`
	PUUID        string     `json:"puuid"`
	LeaguePoints int        `json:"leaguePoints"`
	Rank         string     `json:"rank"`
	Wins         int        `json:"wins"`
	Losses       int        `json:"losses"`
	Veteran      bool       `json:"veteran"`
	Inactive     bool       `json:"inactive"`
	FreshBlood   bool       `json:"freshBlood"`
	HotStreak    bool       `json:"hotStreak"`
	MiniSeries   MiniSeries `json:"miniSeries"`
}

type MiniSeries struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}
