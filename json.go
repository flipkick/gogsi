package gogsi

import (
	"encoding/json"
	"time"

	"github.com/fatih/structs"
)

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

// UnmarshalJSON implements the json.Unmarshaler interface for Ability in order to handle custom time.Duration
// deserialization.
func (a *Ability) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, a); err != nil {
		return err
	}

	a.Cooldown *= time.Second
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Item in order to handle custom time.Duration
// deserialization.
func (i *Item) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, i); err != nil {
		return err
	}
	i.Cooldown *= time.Second
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Hero in order to handle custom time.Duration
// deserialization.
func (h *Hero) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, h); err != nil {
		return err
	}
	h.BuybackCooldown *= time.Second
	h.Respawn *= time.Second
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Provider in order to handle custom time.Duration
// deserialization.
func (p *Provider) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, p); err != nil {
		return err
	}

	var t struct {
		time.Time `json:"timestamp"`
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	p.Timestamp = t.Time

	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Map in order to handle custom time.Duration
// deserialization.
func (m *Map) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, m); err != nil {
		return err
	}
	m.ClockTime *= time.Second
	m.GameTime *= time.Second
	m.WardPurchaseCooldown *= time.Second
	return nil
}

type (
	// abilityIntermediate matches the structure we receive from Dota. It is decoded
	// and transformed to an Abilities struct.
	abilityIntermediate struct {
		Ability1 *Ability    `json:"ability0,omitempty"`
		Ability2 *Ability    `json:"ability1,omitempty"`
		Ability3 *Ability    `json:"ability2,omitempty"`
		Ability4 *Ability    `json:"ability3,omtiempty"`
		Ability5 *Ability    `json:"ability4,omitempty"`
		Ability6 *Ability    `json:"ability5,omitempty"`
		Attr     *Attributes `json:"attributes,omitempty"`
	}

	// itemIntermediate matches the structure we receive from Dota. It is decoded
	// and transformed to an Items struct.
	itemIntermediate struct {
		Item1  *Item `json:"slot0,omitempty"`
		Item2  *Item `json:"slot1,omitempty"`
		Item3  *Item `json:"slot2,omitempty"`
		Item4  *Item `json:"slot3,omtiempty"`
		Item5  *Item `json:"slot4,omitempty"`
		Item6  *Item `json:"slot5,omitempty"`
		Stash1 *Item `json:"stash0,omitempty"`
		Stash2 *Item `json:"stash1,omitempty"`
		Stash3 *Item `json:"stash2,omitempty"`
		Stash4 *Item `json:"stash3,omitempty"`
		Stash5 *Item `json:"stash4,omitempty"`
		Stash6 *Item `json:"stash5,omitempty"`
	}
)

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
