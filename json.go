package gogsi

import (
	"encoding/json"
	"time"

	"github.com/fatih/structs"
)

type (
	// Time wraps the time.Time struct, providing JSON deserialization from a JSON int as epoch seconds.
	Time struct {
		time.Time
	}

	// Duration wraps time.Duration, providing JSON deserialization on the assumption that a JSON int
	// is a number of seconds.
	Duration struct {
		time.Duration
	}
)

// UnmarshalJSON implements the json.Unmarshaler interface for Time, allowing it to deserialize from an
// epoch time int.
func (t *Time) UnmarshalJSON(b []byte) error {
	var i int64
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}

	t.Time = time.Unix(i, 0)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Duration, allowing it to be deserialized
// from a JSON int interpreted as a number of seconds.
func (d *Duration) UnmarshalJSON(b []byte) error {
	var t time.Duration
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	d.Duration = t * time.Second
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for DotaGameState, ensuring that all types resolve
// to the 'enum' values.
func (d *DotaGameState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	case StateDisconnect, StateGameInProgress, StateHeroSelection, StateInit, StateLast, StatePostGame,
		StatePreGame, StateStrategyTime, StateWaitForPlayers:
	default:
		*d = StateUndefined
	}

	*d = DotaGameState(s)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for DotaTeam, ensuring that all types resolve to the
// 'enum' values.
func (d *DotaTeam) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	case TeamDire, TeamNone, TeamRadiant:
	default:
		*d = TeamUndefined
	}

	*d = DotaTeam(s)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RuneType, ensuring that all types resolve to the
// 'enum' values.
func (d *RuneType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	default:
		*d = RuneNone
	}

	*d = RuneType(s)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PlayerActivity, ensuring that all types resolve to
// the 'enum' values.
func (d *PlayerActivity) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	case ActivityMenu, ActivityPlaying:
	default:
		*d = ActivityUndefined
	}

	*d = PlayerActivity(s)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Player, allowing us
// to segregate GameStats from metadata (SteamID, Name, Activity).
func (p *Player) UnmarshalJSON(b []byte) error {
	var intermed playerIntermediate

	err := json.Unmarshal(b, &intermed)

	if err != nil {
		return err
	}

	gameStats := &GameStats{
		Kills:          intermed.Kills,
		Deaths:         intermed.Kills,
		Assists:        intermed.Assists,
		LastHits:       intermed.LastHits,
		Denies:         intermed.Denies,
		KillStreak:     intermed.KillStreak,
		Team:           intermed.Team,
		ReliableGold:   intermed.ReliableGold,
		UnreliableGold: intermed.UnreliableGold,
		GPM:            intermed.GPM,
		XPM:            intermed.XPM,
	}

	if structs.IsZero(*gameStats) {
		gameStats = nil
	}

	*p = Player{
		SteamID:   intermed.SteamID,
		Name:      intermed.Name,
		Activity:  intermed.Activity,
		GameStats: gameStats,
	}

	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Abilities. This is implemented
// custom because the incoming format from Valve is different from the gogsi representation.
func (a *Abilities) UnmarshalJSON(b []byte) error {
	var intermed abilityIntermediate

	err := json.Unmarshal(b, &intermed)
	if err != nil {
		return err
	}

	*a = Abilities{
		[6]*Ability{
			intermed.Ability1, intermed.Ability2,
			intermed.Ability3, intermed.Ability4,
			intermed.Ability5, intermed.Ability6,
		},
		intermed.Attr,
	}

	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Items. This is implemented
// custom because the incoming format from Valve is different from the gogsi representation.
func (i *Items) UnmarshalJSON(b []byte) error {
	var intermed itemIntermediate

	err := json.Unmarshal(b, &intermed)
	if err != nil {
		return err
	}

	*i = Items{
		[6]*Item{
			nilIfEmpty(intermed.Item1), nilIfEmpty(intermed.Item2), nilIfEmpty(intermed.Item3),
			nilIfEmpty(intermed.Item4), nilIfEmpty(intermed.Item5), nilIfEmpty(intermed.Item6),
		},
		[6]*Item{
			nilIfEmpty(intermed.Stash1), nilIfEmpty(intermed.Stash2), nilIfEmpty(intermed.Stash3),
			nilIfEmpty(intermed.Stash4), nilIfEmpty(intermed.Stash5), nilIfEmpty(intermed.Stash6),
		},
	}

	return nil
}

func nilIfEmpty(a *Item) *Item {
	if a != nil && a.Name != "empty" {
		return a
	}
	return nil
}
