package gogsi

import "time"

type (
	// State contains all available information about the current state of a Dota 2 game at a given
	// moment in time. Previous and Added are deltas since the last State (Added contains all new values,
	// and Previous is what they were before.
	State struct {
		*Auth      `json:"auth,omitempty"`
		*Provider  `json:"provider,omitempty"`
		*Map       `json:"map,omitempty"`
		*Player    `json:"player,omitempty"`
		*Hero      `json:"hero,omitempty"`
		*Abilities `json:"abilities,omitempty"`
		*Items     `json:"items,omitempty"`
		Previous   *State `json:"previous,omitempty"`
		Added      *State `json:"added,omitempty"`
	}

	// Map carries information about the state of the Dota 2 map.
	Map struct {
		Name                 string        `json:"name"`
		MatchID              int           `json:"matchid"`
		GameTime             time.Duration `json:"game_time"`
		ClockTime            time.Duration `json:"clock_time"`
		Day                  bool          `json:"daytime"`
		NightstalkerNight    bool          `json:"nightstalker_night"`
		GameState            DotaGameState `json:"game_state"`
		WinTeam              DotaTeam      `json:"win_team"`
		CustomGameName       string        `json:"customgamename"`
		WardPurchaseCooldown time.Duration `json:"ward_purchase_cooldown"`
	}

	// Player provides information about the Steam/Dota user currently logged in to the client.
	Player struct {
		SteamID  string
		Name     string
		Activity PlayerActivity
		*GameStats
	}

	// GameStats holds the information about a player's performance in a given game. Usually nil if
	// the player is not currently in a game.
	GameStats struct {
		Kills          int
		Deaths         int
		Assists        int
		LastHits       int
		Denies         int
		KillStreak     int
		Team           DotaTeam
		ReliableGold   int
		UnreliableGold int
		GPM            int
		XPM            int
	}

	// playerIntermediate matches the incoming JSON structure from the Dota client to allow
	// for unmarshaling into Player.
	playerIntermediate struct {
		SteamID        string         `json:"steamid"`
		Name           string         `json:"name"`
		Activity       PlayerActivity `json:"activity"`
		Kills          int            `json:"kills"`
		Deaths         int            `json:"deaths"`
		Assists        int            `json:"assists"`
		LastHits       int            `json:"last_hits"`
		Denies         int            `json:"denies"`
		KillStreak     int            `json:"kill_streak"`
		Team           DotaTeam       `json:"team_name"`
		ReliableGold   int            `json:"gold_reliable"`
		UnreliableGold int            `json:"gold_unreliable"`
		GPM            int            `json:"gpm"`
		XPM            int            `json:"xpm"`
	}

	// Provider holds Steam-related metadata about the app currently in use (Dota 2).
	Provider struct {
		Name      string    `json:"name"`
		AppID     int       `json:"appid"`
		Version   int       `json:"version"`
		Timestamp time.Time `json:"-"`
	}

	// Auth holds the token set in the integration's configuration file (a shared key is easiest).
	// Mostly useful for remote listeners.
	Auth struct {
		Token string `json:"token"`
	}
)

// GoldAfterDeath returns the amount of gold the player would have if their hero were to die. Will not work
// properly in custom games.
func (s *State) GoldAfterDeath() int {
	if s == nil || s.Hero == nil || s.Player == nil {
		return 0
	}
	return s.ReliableGold + s.UnreliableGoldAfterDeath()
}

// UnreliableGoldAfterDeath computes the amount of unreliable gold the player would have if their hero were to die.
// Will not work properly for custom games.
func (s *State) UnreliableGoldAfterDeath() int {
	if s == nil || s.Hero == nil || s.Player == nil {
		return 0
	}

	deathGold := 30 * s.Level

	if s.UnreliableGold < deathGold {
		return 0
	}
	return s.UnreliableGold - deathGold
}

// CanAffordBuyback returns whether or not the player has enough gold to buy back. Ignores buyback timer.
func (s *State) CanAffordBuyback() bool {
	if s == nil || s.Hero == nil || s.Player == nil {
		return false
	}

	return s.BuybackCost < (s.ReliableGold + s.UnreliableGoldAfterDeath())
}
