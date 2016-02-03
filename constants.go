package gogsi

type (
	// DotaGameState is an enumeration of all possible dota game states.
	DotaGameState string

	// DotaTeam is the team the player is on. Unknown how this plays with custom games.
	DotaTeam string

	// PlayerActivity enumerates the different states a player might be in. This
	// is not an exhaustive list.
	PlayerActivity string

	// ItemName enumerates possible item names. As of now, this type simply captures the "empty" value.
	ItemName string

	// RuneType enumerates the runes that a bottle could contain. At the moment, these values haven't been
	// checked or enumerated.
	RuneType string
)

// Known state constants
const (
	StateUndefined      DotaGameState = ""
	StateDisconnect                   = "DOTA_GAMERULES_STATE_DISCONNECT"
	StateGameInProgress               = "DOTA_GAMERULES_STATE_GAME_IN_PROGRESS"
	StateHeroSelection                = "DOTA_GAMERULES_STATE_HERO_SELECTION"
	StateInit                         = "DOTA_GAMERULES_STATE_INIT"
	StateLast                         = "DOTA_GAMERULES_STATE_LAST"
	StatePostGame                     = "DOTA_GAMERULES_STATE_POST_GAME"
	StatePreGame                      = "DOTA_GAMERULES_STATE_PRE_GAME"
	StateStrategyTime                 = "DOTA_GAMERULES_STATE_STRATEGY_TIME"
	StateWaitForPlayers               = "DOTA_GAMERULES_STATE_WAIT_FOR_PLAYERS"

	TeamUndefined DotaTeam = ""
	TeamNone               = "none"
	TeamDire               = "dire"
	TeamRadiant            = "radiant"

	ActivityUndefined PlayerActivity = ""
	ActivityMenu                     = "menu"
	ActivityPlaying                  = "playing"

	RuneNone RuneType = ""
)
