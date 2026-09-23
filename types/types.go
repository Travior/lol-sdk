package types

type Account struct {
	PUUID    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type Summoner struct {
	// Deprecated: Riot is removing summoner IDs. Use PUUID instead.
	ID            string `json:"id"`
	PUUID         string `json:"puuid"`
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
	EndOfGameResult    string        `json:"endOfGameResult"`
	GameCreation       int64         `json:"gameCreation"`
	GameDuration       int           `json:"gameDuration"`
	GameEndTimestamp   int64         `json:"gameEndTimestamp"`
	GameID             int64         `json:"gameId"`
	GameMode           string        `json:"gameMode"`
	GameModeMutators   []string      `json:"gameModeMutators"`
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
	AllInPings                     int                       `json:"allInPings"`
	AssistMePings                  int                       `json:"assistMePings"`
	Assists                        int                       `json:"assists"`
	BaitPings                      int                       `json:"baitPings"`
	BaronKills                     int                       `json:"baronKills"`
	BasicPings                     int                       `json:"basicPings"`
	BountyLevel                    int                       `json:"bountyLevel"`
	CausedGameEndFromIGNBSurrender bool                      `json:"causedGameEndFromIGNBSurrender"`
	Challenges                     ParticipantChallenges     `json:"challenges"`
	ChampExperience                int                       `json:"champExperience"`
	ChampionID                     int                       `json:"championId"`
	ChampionName                   string                    `json:"championName"`
	ChampionSkinID                 int                       `json:"championSkinId"`
	ChampionTransform              int                       `json:"championTransform"`
	ChampLevel                     int                       `json:"champLevel"`
	CommandPings                   int                       `json:"commandPings"`
	ConsumablesPurchased           int                       `json:"consumablesPurchased"`
	DamageDealtToBuildings         int                       `json:"damageDealtToBuildings"`
	DamageDealtToEpicMonsters      int                       `json:"damageDealtToEpicMonsters"`
	DamageDealtToObjectives        int                       `json:"damageDealtToObjectives"`
	DamageDealtToTurrets           int                       `json:"damageDealtToTurrets"`
	DamageSelfMitigated            int                       `json:"damageSelfMitigated"`
	DangerPings                    int                       `json:"dangerPings"`
	Deaths                         int                       `json:"deaths"`
	DetectorWardsPlaced            int                       `json:"detectorWardsPlaced"`
	DoubleKills                    int                       `json:"doubleKills"`
	DragonKills                    int                       `json:"dragonKills"`
	EligibleForProgression         bool                      `json:"eligibleForProgression"`
	EnemyMissingPings              int                       `json:"enemyMissingPings"`
	EnemyVisionPings               int                       `json:"enemyVisionPings"`
	FirstBloodAssist               bool                      `json:"firstBloodAssist"`
	FirstBloodKill                 bool                      `json:"firstBloodKill"`
	FirstTowerAssist               bool                      `json:"firstTowerAssist"`
	FirstTowerKill                 bool                      `json:"firstTowerKill"`
	GameEndedInEarlySurrender      bool                      `json:"gameEndedInEarlySurrender"`
	GameEndedInIGNBSurrender       bool                      `json:"gameEndedInIGNBSurrender"`
	GameEndedInSurrender           bool                      `json:"gameEndedInSurrender"`
	GetBackPings                   int                       `json:"getBackPings"`
	GoldEarned                     int                       `json:"goldEarned"`
	GoldSpent                      int                       `json:"goldSpent"`
	HoldPings                      int                       `json:"holdPings"`
	IndividualPosition             string                    `json:"individualPosition"`
	InhibitorKills                 int                       `json:"inhibitorKills"`
	InhibitorsLost                 int                       `json:"inhibitorsLost"`
	InhibitorTakedowns             int                       `json:"inhibitorTakedowns"`
	Item0                          int                       `json:"item0"`
	Item1                          int                       `json:"item1"`
	Item2                          int                       `json:"item2"`
	Item3                          int                       `json:"item3"`
	Item4                          int                       `json:"item4"`
	Item5                          int                       `json:"item5"`
	Item6                          int                       `json:"item6"`
	ItemsPurchased                 int                       `json:"itemsPurchased"`
	KillingSprees                  int                       `json:"killingSprees"`
	Kills                          int                       `json:"kills"`
	Lane                           string                    `json:"lane"`
	LargestCriticalStrike          int                       `json:"largestCriticalStrike"`
	LargestKillingSpree            int                       `json:"largestKillingSpree"`
	LargestMultiKill               int                       `json:"largestMultiKill"`
	LongestTimeSpentLiving         int                       `json:"longestTimeSpentLiving"`
	MagicDamageDealt               int                       `json:"magicDamageDealt"`
	MagicDamageDealtToChampions    int                       `json:"magicDamageDealtToChampions"`
	MagicDamageTaken               int                       `json:"magicDamageTaken"`
	Missions                       ParticipantMissions       `json:"missions"`
	NeedVisionPings                int                       `json:"needVisionPings"`
	NeutralMinionsKilled           int                       `json:"neutralMinionsKilled"`
	NexusKills                     int                       `json:"nexusKills"`
	NexusLost                      int                       `json:"nexusLost"`
	NexusTakedowns                 int                       `json:"nexusTakedowns"`
	ObjectivesStolen               int                       `json:"objectivesStolen"`
	ObjectivesStolenAssists        int                       `json:"objectivesStolenAssists"`
	OnMyWayPings                   int                       `json:"onMyWayPings"`
	ParticipantID                  int                       `json:"participantId"`
	PentaKills                     int                       `json:"pentaKills"`
	Perks                          ParticipantPerks          `json:"perks"`
	PhysicalDamageDealt            int                       `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions int                       `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken            int                       `json:"physicalDamageTaken"`
	Placement                      int                       `json:"placement"`
	PlayerAugment1                 int                       `json:"playerAugment1"`
	PlayerAugment2                 int                       `json:"playerAugment2"`
	PlayerAugment3                 int                       `json:"playerAugment3"`
	PlayerAugment4                 int                       `json:"playerAugment4"`
	PlayerAugment5                 int                       `json:"playerAugment5"`
	PlayerAugment6                 int                       `json:"playerAugment6"`
	PlayerBehavior                 ParticipantPlayerBehavior `json:"PlayerBehavior"`
	PlayerScore0                   float64                   `json:"playerScore0"`
	PlayerScore1                   float64                   `json:"playerScore1"`
	PlayerScore10                  float64                   `json:"playerScore10"`
	PlayerScore11                  float64                   `json:"playerScore11"`
	PlayerScore2                   float64                   `json:"playerScore2"`
	PlayerScore3                   float64                   `json:"playerScore3"`
	PlayerScore4                   float64                   `json:"playerScore4"`
	PlayerScore5                   float64                   `json:"playerScore5"`
	PlayerScore6                   float64                   `json:"playerScore6"`
	PlayerScore7                   float64                   `json:"playerScore7"`
	PlayerScore8                   float64                   `json:"playerScore8"`
	PlayerScore9                   float64                   `json:"playerScore9"`
	PlayerSubteamID                int                       `json:"playerSubteamId"`
	PositionAssignedByMatchmaking  string                    `json:"positionAssignedByMatchmaking"`
	ProfileIcon                    int                       `json:"profileIcon"`
	PushPings                      int                       `json:"pushPings"`
	PUUID                          string                    `json:"puuid"`
	QuadraKills                    int                       `json:"quadraKills"`
	RetreatPings                   int                       `json:"retreatPings"`
	RiotIDGameName                 string                    `json:"riotIdGameName"`
	// Deprecated: use RiotIDGameName instead.
	RiotIDName              string `json:"riotIdName"`
	RiotIDTagline           string `json:"riotIdTagline"`
	Role                    string `json:"role"`
	RoleBoundItem           int    `json:"roleBoundItem"`
	SelectedRolePreferences string `json:"selectedRolePreferences"`
	SightWardsBoughtInGame  int    `json:"sightWardsBoughtInGame"`
	Spell1Casts             int    `json:"spell1Casts"`
	Spell2Casts             int    `json:"spell2Casts"`
	Spell3Casts             int    `json:"spell3Casts"`
	Spell4Casts             int    `json:"spell4Casts"`
	SubteamPlacement        int    `json:"subteamPlacement"`
	Summoner1Casts          int    `json:"summoner1Casts"`
	Summoner1ID             int    `json:"summoner1Id"`
	Summoner2Casts          int    `json:"summoner2Casts"`
	Summoner2ID             int    `json:"summoner2Id"`
	// Deprecated: Riot is removing summoner IDs. Use PUUID instead.
	SummonerID    string `json:"summonerId"`
	SummonerLevel int    `json:"summonerLevel"`
	// Deprecated: summoner names were replaced by Riot IDs. Use RiotIDGameName and RiotIDTagline instead.
	SummonerName                     string `json:"summonerName"`
	TeamEarlySurrendered             bool   `json:"teamEarlySurrendered"`
	TeamID                           int    `json:"teamId"`
	TeamIGNBSurrendered              bool   `json:"teamIGNBSurrendered"`
	TeamPosition                     string `json:"teamPosition"`
	TimeCCingOthers                  int    `json:"timeCCingOthers"`
	TimePlayed                       int    `json:"timePlayed"`
	TotalAllyJungleMinionsKilled     int    `json:"totalAllyJungleMinionsKilled"`
	TotalDamageDealt                 int    `json:"totalDamageDealt"`
	TotalDamageDealtToChampions      int    `json:"totalDamageDealtToChampions"`
	TotalDamageShieldedOnTeammates   int    `json:"totalDamageShieldedOnTeammates"`
	TotalDamageTaken                 int    `json:"totalDamageTaken"`
	TotalEnemyJungleMinionsKilled    int    `json:"totalEnemyJungleMinionsKilled"`
	TotalHeal                        int    `json:"totalHeal"`
	TotalHealsOnTeammates            int    `json:"totalHealsOnTeammates"`
	TotalMinionsKilled               int    `json:"totalMinionsKilled"`
	TotalTimeCCDealt                 int    `json:"totalTimeCCDealt"`
	TotalTimeSpentDead               int    `json:"totalTimeSpentDead"`
	TotalUnitsHealed                 int    `json:"totalUnitsHealed"`
	TripleKills                      int    `json:"tripleKills"`
	TrueDamageDealt                  int    `json:"trueDamageDealt"`
	TrueDamageDealtToChampions       int    `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                  int    `json:"trueDamageTaken"`
	TurretKills                      int    `json:"turretKills"`
	TurretsLost                      int    `json:"turretsLost"`
	TurretTakedowns                  int    `json:"turretTakedowns"`
	UnrealKills                      int    `json:"unrealKills"`
	VisionClearedPings               int    `json:"visionClearedPings"`
	VisionScore                      int    `json:"visionScore"`
	VisionWardsBoughtInGame          int    `json:"visionWardsBoughtInGame"`
	WardsKilled                      int    `json:"wardsKilled"`
	WardsPlaced                      int    `json:"wardsPlaced"`
	WasPremadeWithIGNBGameEndCauser  bool   `json:"wasPremadeWithIGNBGameEndCauser"`
	WasPremadeWithSevereTransgressor bool   `json:"wasPremadeWithSevereTransgressor"`
	WasSevereTransgressor            bool   `json:"wasSevereTransgressor"`
	Win                              bool   `json:"win"`
}

type ParticipantPlayerBehavior struct {
	IsHeroInCombat int `json:"PlayerBehavior_IsHeroInCombat"`
}

type ParticipantMissions struct {
	PlayerScore0  float64 `json:"playerScore0"`
	PlayerScore1  float64 `json:"playerScore1"`
	PlayerScore2  float64 `json:"playerScore2"`
	PlayerScore3  float64 `json:"playerScore3"`
	PlayerScore4  float64 `json:"playerScore4"`
	PlayerScore5  float64 `json:"playerScore5"`
	PlayerScore6  float64 `json:"playerScore6"`
	PlayerScore7  float64 `json:"playerScore7"`
	PlayerScore8  float64 `json:"playerScore8"`
	PlayerScore9  float64 `json:"playerScore9"`
	PlayerScore10 float64 `json:"playerScore10"`
	PlayerScore11 float64 `json:"playerScore11"`
}

// Challenge values are decoded as float64 because Riot does not consistently
// return whole numbers for fields documented as integers.
type ParticipantChallenges struct {
	AssistStreakCount12                       float64 `json:"12AssistStreakCount,omitempty"`
	AbilityUses                               float64 `json:"abilityUses,omitempty"`
	AcesBefore15Minutes                       float64 `json:"acesBefore15Minutes,omitempty"`
	AlliedJungleMonsterKills                  float64 `json:"alliedJungleMonsterKills,omitempty"`
	BaronBuffGoldAdvantageOverThreshold       float64 `json:"baronBuffGoldAdvantageOverThreshold,omitempty"`
	BaronTakedowns                            float64 `json:"baronTakedowns,omitempty"`
	BlastConeOppositeOpponentCount            float64 `json:"blastConeOppositeOpponentCount,omitempty"`
	BountyGold                                float64 `json:"bountyGold,omitempty"`
	BuffsStolen                               float64 `json:"buffsStolen,omitempty"`
	CompleteSupportQuestInTime                float64 `json:"completeSupportQuestInTime,omitempty"`
	ControlWardsPlaced                        float64 `json:"controlWardsPlaced,omitempty"`
	ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf,omitempty"`
	DamagePerMinute                           float64 `json:"damagePerMinute,omitempty"`
	DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage,omitempty"`
	DancedWithRiftHerald                      float64 `json:"dancedWithRiftHerald,omitempty"`
	DeathsByEnemyChamps                       float64 `json:"deathsByEnemyChamps,omitempty"`
	DodgeSkillShotsSmallWindow                float64 `json:"dodgeSkillShotsSmallWindow,omitempty"`
	DoubleAces                                float64 `json:"doubleAces,omitempty"`
	DragonTakedowns                           float64 `json:"dragonTakedowns,omitempty"`
	EarliestBaron                             float64 `json:"earliestBaron,omitempty"`
	EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown,omitempty"`
	EarliestElderDragon                       float64 `json:"earliestElderDragon,omitempty"`
	EarlyLaningPhaseGoldExpAdvantage          float64 `json:"earlyLaningPhaseGoldExpAdvantage,omitempty"`
	EffectiveHealAndShielding                 float64 `json:"effectiveHealAndShielding,omitempty"`
	ElderDragonKillsWithOpposingSoul          float64 `json:"elderDragonKillsWithOpposingSoul,omitempty"`
	ElderDragonMultikills                     float64 `json:"elderDragonMultikills,omitempty"`
	EnemyChampionImmobilizations              float64 `json:"enemyChampionImmobilizations,omitempty"`
	EnemyJungleMonsterKills                   float64 `json:"enemyJungleMonsterKills,omitempty"`
	EpicMonsterKillsNearEnemyJungler          float64 `json:"epicMonsterKillsNearEnemyJungler,omitempty"`
	EpicMonsterKillsWithin30SecondsOfSpawn    float64 `json:"epicMonsterKillsWithin30SecondsOfSpawn,omitempty"`
	EpicMonsterSteals                         float64 `json:"epicMonsterSteals,omitempty"`
	EpicMonsterStolenWithoutSmite             float64 `json:"epicMonsterStolenWithoutSmite,omitempty"`
	FasterSupportQuestCompletion              float64 `json:"fasterSupportQuestCompletion,omitempty"`
	FastestLegendary                          float64 `json:"fastestLegendary,omitempty"`
	FirstTurretKilled                         float64 `json:"firstTurretKilled,omitempty"`
	FirstTurretKilledTime                     float64 `json:"firstTurretKilledTime,omitempty"`
	FistBumpParticipation                     float64 `json:"fistBumpParticipation,omitempty"`
	FlawlessAces                              float64 `json:"flawlessAces,omitempty"`
	FullTeamTakedown                          float64 `json:"fullTeamTakedown,omitempty"`
	GameLength                                float64 `json:"gameLength,omitempty"`
	GetTakedownsInAllLanesEarlyJungleAsLaner  float64 `json:"getTakedownsInAllLanesEarlyJungleAsLaner,omitempty"`
	GoldPerMinute                             float64 `json:"goldPerMinute,omitempty"`
	HadAfkTeammate                            float64 `json:"hadAfkTeammate,omitempty"`
	HadOpenNexus                              float64 `json:"hadOpenNexus,omitempty"`
	HealFromMapSources                        float64 `json:"HealFromMapSources,omitempty"`
	HighestChampionDamage                     float64 `json:"highestChampionDamage,omitempty"`
	HighestCrowdControlScore                  float64 `json:"highestCrowdControlScore,omitempty"`
	HighestWardKills                          float64 `json:"highestWardKills,omitempty"`
	ImmobilizeAndKillWithAlly                 float64 `json:"immobilizeAndKillWithAlly,omitempty"`
	InfernalScalePickup                       float64 `json:"InfernalScalePickup,omitempty"`
	InitialBuffCount                          float64 `json:"initialBuffCount,omitempty"`
	InitialCrabCount                          float64 `json:"initialCrabCount,omitempty"`
	JungleCSBefore10Minutes                   float64 `json:"jungleCsBefore10Minutes,omitempty"`
	JunglerKillsEarlyJungle                   float64 `json:"junglerKillsEarlyJungle,omitempty"`
	JunglerTakedownsNearDamagedEpicMonster    float64 `json:"junglerTakedownsNearDamagedEpicMonster,omitempty"`
	KDA                                       float64 `json:"kda,omitempty"`
	KillAfterHiddenWithAlly                   float64 `json:"killAfterHiddenWithAlly,omitempty"`
	KilledChampTookFullTeamDamageSurvived     float64 `json:"killedChampTookFullTeamDamageSurvived,omitempty"`
	KillingSprees                             float64 `json:"killingSprees,omitempty"`
	KillParticipation                         float64 `json:"killParticipation,omitempty"`
	KillsNearEnemyTurret                      float64 `json:"killsNearEnemyTurret,omitempty"`
	KillsOnLanersEarlyJungleAsJungler         float64 `json:"killsOnLanersEarlyJungleAsJungler,omitempty"`
	KillsOnOtherLanesEarlyJungleAsLaner       float64 `json:"killsOnOtherLanesEarlyJungleAsLaner,omitempty"`
	KillsOnRecentlyHealedByARAMPack           float64 `json:"killsOnRecentlyHealedByAramPack,omitempty"`
	KillsUnderOwnTurret                       float64 `json:"killsUnderOwnTurret,omitempty"`
	KillsWithHelpFromEpicMonster              float64 `json:"killsWithHelpFromEpicMonster,omitempty"`
	KnockEnemyIntoTeamAndKill                 float64 `json:"knockEnemyIntoTeamAndKill,omitempty"`
	KTurretsDestroyedBeforePlatesFall         float64 `json:"kTurretsDestroyedBeforePlatesFall,omitempty"`
	LandSkillShotsEarlyGame                   float64 `json:"landSkillShotsEarlyGame,omitempty"`
	LaneMinionsFirst10Minutes                 float64 `json:"laneMinionsFirst10Minutes,omitempty"`
	LaningPhaseGoldExpAdvantage               float64 `json:"laningPhaseGoldExpAdvantage,omitempty"`
	LegendaryCount                            float64 `json:"legendaryCount,omitempty"`
	LegendaryItemUsed                         []int   `json:"legendaryItemUsed,omitempty"`
	LostAnInhibitor                           float64 `json:"lostAnInhibitor,omitempty"`
	MaxCSAdvantageOnLaneOpponent              float64 `json:"maxCsAdvantageOnLaneOpponent,omitempty"`
	MaxKillDeficit                            float64 `json:"maxKillDeficit,omitempty"`
	MaxLevelLeadLaneOpponent                  float64 `json:"maxLevelLeadLaneOpponent,omitempty"`
	MejaisFullStackInTime                     float64 `json:"mejaisFullStackInTime,omitempty"`
	MoreEnemyJungleThanOpponent               float64 `json:"moreEnemyJungleThanOpponent,omitempty"`
	MostWardsDestroyedOneSweeper              float64 `json:"mostWardsDestroyedOneSweeper,omitempty"`
	MultiKillOneSpell                         float64 `json:"multiKillOneSpell,omitempty"`
	Multikills                                float64 `json:"multikills,omitempty"`
	MultikillsAfterAggressiveFlash            float64 `json:"multikillsAfterAggressiveFlash,omitempty"`
	MultiTurretRiftHeraldCount                float64 `json:"multiTurretRiftHeraldCount,omitempty"`
	MythicItemUsed                            float64 `json:"mythicItemUsed,omitempty"`
	OuterTurretExecutesBefore10Minutes        float64 `json:"outerTurretExecutesBefore10Minutes,omitempty"`
	OutnumberedKills                          float64 `json:"outnumberedKills,omitempty"`
	OutnumberedNexusKill                      float64 `json:"outnumberedNexusKill,omitempty"`
	PerfectDragonSoulsTaken                   float64 `json:"perfectDragonSoulsTaken,omitempty"`
	PerfectGame                               float64 `json:"perfectGame,omitempty"`
	PickKillWithAlly                          float64 `json:"pickKillWithAlly,omitempty"`
	PlayedChampSelectPosition                 float64 `json:"playedChampSelectPosition,omitempty"`
	PoroExplosions                            float64 `json:"poroExplosions,omitempty"`
	QuickCleanse                              float64 `json:"quickCleanse,omitempty"`
	QuickFirstTurret                          float64 `json:"quickFirstTurret,omitempty"`
	QuickSoloKills                            float64 `json:"quickSoloKills,omitempty"`
	RiftHeraldTakedowns                       float64 `json:"riftHeraldTakedowns,omitempty"`
	SaveAllyFromDeath                         float64 `json:"saveAllyFromDeath,omitempty"`
	ScuttleCrabKills                          float64 `json:"scuttleCrabKills,omitempty"`
	ShortestTimeToAceFromFirstTakedown        float64 `json:"shortestTimeToAceFromFirstTakedown,omitempty"`
	SkillshotsDodged                          float64 `json:"skillshotsDodged,omitempty"`
	SkillshotsHit                             float64 `json:"skillshotsHit,omitempty"`
	SnowballsHit                              float64 `json:"snowballsHit,omitempty"`
	SoloBaronKills                            float64 `json:"soloBaronKills,omitempty"`
	SoloKills                                 float64 `json:"soloKills,omitempty"`
	SoloTurretsLategame                       float64 `json:"soloTurretsLategame,omitempty"`
	StealthWardsPlaced                        float64 `json:"stealthWardsPlaced,omitempty"`
	SurvivedSingleDigitHpCount                float64 `json:"survivedSingleDigitHpCount,omitempty"`
	SurvivedThreeImmobilizesInFight           float64 `json:"survivedThreeImmobilizesInFight,omitempty"`
	SWARMDefeatAatrox                         float64 `json:"SWARM_DefeatAatrox,omitempty"`
	SWARMDefeatBriar                          float64 `json:"SWARM_DefeatBriar,omitempty"`
	SWARMDefeatMiniBosses                     float64 `json:"SWARM_DefeatMiniBosses,omitempty"`
	SWARMEvolveWeapon                         float64 `json:"SWARM_EvolveWeapon,omitempty"`
	SWARMHave3Passives                        float64 `json:"SWARM_Have3Passives,omitempty"`
	SWARMKillEnemy                            float64 `json:"SWARM_KillEnemy,omitempty"`
	SWARMPickupGold                           float64 `json:"SWARM_PickupGold,omitempty"`
	SWARMReachLevel50                         float64 `json:"SWARM_ReachLevel50,omitempty"`
	SWARMSurvive15Min                         float64 `json:"SWARM_Survive15Min,omitempty"`
	SWARMWinWith5EvolvedWeapons               float64 `json:"SWARM_WinWith5EvolvedWeapons,omitempty"`
	TakedownOnFirstTurret                     float64 `json:"takedownOnFirstTurret,omitempty"`
	Takedowns                                 float64 `json:"takedowns,omitempty"`
	TakedownsAfterGainingLevelAdvantage       float64 `json:"takedownsAfterGainingLevelAdvantage,omitempty"`
	TakedownsBeforeJungleMinionSpawn          float64 `json:"takedownsBeforeJungleMinionSpawn,omitempty"`
	TakedownsFirst25Minutes                   float64 `json:"takedownsFirst25Minutes,omitempty"`
	TakedownsFirstXMinutes                    float64 `json:"takedownsFirstXMinutes,omitempty"`
	TakedownsInAlcove                         float64 `json:"takedownsInAlcove,omitempty"`
	TakedownsInEnemyFountain                  float64 `json:"takedownsInEnemyFountain,omitempty"`
	TeamBaronKills                            float64 `json:"teamBaronKills,omitempty"`
	TeamDamagePercentage                      float64 `json:"teamDamagePercentage,omitempty"`
	TeamElderDragonKills                      float64 `json:"teamElderDragonKills,omitempty"`
	TeamRiftHeraldKills                       float64 `json:"teamRiftHeraldKills,omitempty"`
	TeleportTakedowns                         float64 `json:"teleportTakedowns,omitempty"`
	ThirdInhibitorDestroyedTime               float64 `json:"thirdInhibitorDestroyedTime,omitempty"`
	ThreeWardsOneSweeperCount                 float64 `json:"threeWardsOneSweeperCount,omitempty"`
	TookLargeDamageSurvived                   float64 `json:"tookLargeDamageSurvived,omitempty"`
	TurretPlatesTaken                         float64 `json:"turretPlatesTaken,omitempty"`
	TurretsTakenWithRiftHerald                float64 `json:"turretsTakenWithRiftHerald,omitempty"`
	TurretTakedowns                           float64 `json:"turretTakedowns,omitempty"`
	TwentyMinionsIn3SecondsCount              float64 `json:"twentyMinionsIn3SecondsCount,omitempty"`
	TwoWardsOneSweeperCount                   float64 `json:"twoWardsOneSweeperCount,omitempty"`
	UnseenRecalls                             float64 `json:"unseenRecalls,omitempty"`
	VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent,omitempty"`
	VisionScorePerMinute                      float64 `json:"visionScorePerMinute,omitempty"`
	VoidMonsterKill                           float64 `json:"voidMonsterKill,omitempty"`
	WardsGuarded                              float64 `json:"wardsGuarded,omitempty"`
	WardTakedowns                             float64 `json:"wardTakedowns,omitempty"`
	WardTakedownsBefore20M                    float64 `json:"wardTakedownsBefore20M,omitempty"`
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
	Feats      TeamFeats      `json:"feats"`
	Objectives TeamObjectives `json:"objectives"`
	TeamID     int            `json:"teamId"`
	Win        bool           `json:"win"`
}

type TeamFeats struct {
	EpicMonsterKill TeamFeat `json:"EPIC_MONSTER_KILL"`
	FirstBlood      TeamFeat `json:"FIRST_BLOOD"`
	FirstTurret     TeamFeat `json:"FIRST_TURRET"`
}

type TeamFeat struct {
	FeatState int `json:"featState"`
}

type TeamBan struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type TeamObjectives struct {
	Atakhan    TeamObjective `json:"atakhan"`
	Baron      TeamObjective `json:"baron"`
	Champion   TeamObjective `json:"champion"`
	Dragon     TeamObjective `json:"dragon"`
	Horde      TeamObjective `json:"horde"`
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
	EndOfGameResult string                `json:"endOfGameResult"`
	FrameInterval   int                   `json:"frameInterval"`
	Frames          []TimelineFrame       `json:"frames"`
	GameID          int64                 `json:"gameId"`
	Participants    []TimelineParticipant `json:"participants"`
}

type TimelineFrame struct {
	Events            []TimelineEvent                     `json:"events"`
	ParticipantFrames map[string]TimelineParticipantFrame `json:"participantFrames"`
	Timestamp         int                                 `json:"timestamp"`
}

type TimelineEvent struct {
	RealTimestamp                 int64                  `json:"realTimestamp"`
	Timestamp                     int                    `json:"timestamp"`
	Type                          string                 `json:"type"`
	ItemID                        int                    `json:"itemId,omitempty"`
	ParticipantID                 int                    `json:"participantId,omitempty"`
	LevelUpType                   string                 `json:"levelUpType,omitempty"`
	SkillSlot                     int                    `json:"skillSlot,omitempty"`
	CreatorID                     int                    `json:"creatorId,omitempty"`
	WardType                      string                 `json:"wardType,omitempty"`
	Level                         int                    `json:"level,omitempty"`
	AssistingParticipantIDs       []int                  `json:"assistingParticipantIds,omitempty"`
	BountyLevel                   int                    `json:"bountyLevel,omitempty"`
	KillStreakLength              int                    `json:"killStreakLength,omitempty"`
	KillerID                      int                    `json:"killerId,omitempty"`
	Position                      TimelinePosition       `json:"position"`
	VictimDamageDealt             []TimelineVictimDamage `json:"victimDamageDealt,omitempty"`
	VictimDamageReceived          []TimelineVictimDamage `json:"victimDamageReceived,omitempty"`
	VictimID                      int                    `json:"victimId,omitempty"`
	KillType                      string                 `json:"killType,omitempty"`
	LaneType                      string                 `json:"laneType,omitempty"`
	TeamID                        int                    `json:"teamId,omitempty"`
	MonsterType                   string                 `json:"monsterType,omitempty"`
	MonsterSubType                string                 `json:"monsterSubType,omitempty"`
	BuildingType                  string                 `json:"buildingType,omitempty"`
	TowerType                     string                 `json:"towerType,omitempty"`
	AfterID                       int                    `json:"afterId,omitempty"`
	BeforeID                      int                    `json:"beforeId,omitempty"`
	GoldGain                      int                    `json:"goldGain,omitempty"`
	GameID                        int64                  `json:"gameId,omitempty"`
	WinningTeam                   int                    `json:"winningTeam,omitempty"`
	Bounty                        int                    `json:"bounty,omitempty"`
	ShutdownBounty                int                    `json:"shutdownBounty,omitempty"`
	MultiKillLength               int                    `json:"multiKillLength,omitempty"`
	KillerTeamID                  int                    `json:"killerTeamId,omitempty"`
	TransformType                 string                 `json:"transformType,omitempty"`
	ActualStartTime               int64                  `json:"actualStartTime,omitempty"`
	FeatType                      int                    `json:"featType,omitempty"`
	FeatValue                     int                    `json:"featValue,omitempty"`
	VictimTeamfightDamageDealt    []TimelineVictimDamage `json:"victimTeamfightDamageDealt,omitempty"`
	VictimTeamfightDamageReceived []TimelineVictimDamage `json:"victimTeamfightDamageReceived,omitempty"`
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
	Tier     Tier          `json:"tier"`
	Name     string        `json:"name"`
	Queue    Queue         `json:"queue"`
}

// LeagueEntry is used both for entries of a LeagueList and for the entries
// endpoints. LeagueID, QueueType and Tier are only set by the entries endpoints.
type LeagueEntry struct {
	// Deprecated: Riot is removing summoner IDs. Use PUUID instead.
	SummonerID   string     `json:"summonerId"`
	PUUID        string     `json:"puuid"`
	LeagueID     string     `json:"leagueId"`
	QueueType    Queue      `json:"queueType"`
	Tier         Tier       `json:"tier"`
	LeaguePoints int        `json:"leaguePoints"`
	Rank         Division   `json:"rank"`
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
