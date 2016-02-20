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
	StateDisconnect     DotaGameState = "DOTA_GAMERULES_STATE_DISCONNECT"
	StateGameInProgress DotaGameState = "DOTA_GAMERULES_STATE_GAME_IN_PROGRESS"
	StateHeroSelection  DotaGameState = "DOTA_GAMERULES_STATE_HERO_SELECTION"
	StateInit           DotaGameState = "DOTA_GAMERULES_STATE_INIT"
	StateLast           DotaGameState = "DOTA_GAMERULES_STATE_LAST"
	StatePostGame       DotaGameState = "DOTA_GAMERULES_STATE_POST_GAME"
	StatePreGame        DotaGameState = "DOTA_GAMERULES_STATE_PRE_GAME"
	StateStrategyTime   DotaGameState = "DOTA_GAMERULES_STATE_STRATEGY_TIME"
	StateWaitForPlayers DotaGameState = "DOTA_GAMERULES_STATE_WAIT_FOR_PLAYERS"

	TeamUndefined DotaTeam = ""
	TeamNone      DotaTeam = "none"
	TeamDire      DotaTeam = "dire"
	TeamRadiant   DotaTeam = "radiant"

	ActivityUndefined PlayerActivity = ""
	ActivityMenu      PlayerActivity = "menu"
	ActivityPlaying   PlayerActivity = "playing"

	RuneNone RuneType = ""
)
