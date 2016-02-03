package gogsi

type (
	// Hero represents the hero being controlled by the connected client.
	Hero struct {
		ID              int      `json:"id"`
		Name            string   `json:"name"`
		Level           int      `json:"level"`
		Alive           bool     `json:"alive"`
		Respawn         Duration `json:"respawn_seconds"`
		BuybackCost     int      `json:"buyback_cost"`
		BuybackCooldown Duration `json:"buyback_cooldown"`
		Health          int      `json:"health"`
		MaxHealth       int      `json:"max_health"`
		HealthPercent   int      `json:"health_percent"`
		Mana            int      `json:"mana"`
		MaxMana         int      `json:"max_mana"`
		ManaPercent     int      `json:"mana_percent"`
		Silenced        bool     `json:"silenced"`
		Stunned         bool     `json:"stunned"`
		Disarmed        bool     `json:"disarmed"`
		MagicImmune     bool     `json:"magicimmune"`
		Hexed           bool     `json:"hexed"`
		Muted           bool     `json:"muted"`
		Break           bool     `json:"break"`
		Debuffed        bool     `json:"has_debuff"`
	}

	// Ability represents a Hero's ability.
	Ability struct {
		Name     string   `json:"name"`
		Level    int      `json:"level"`
		Cooldown Duration `json:"cooldown"`
		CanCast  bool     `json:"can_cast"`
		Active   bool     `json:"ability_active"`
		Passive  bool     `json:"passive"`
		Ultimate bool     `json:"ultimate"`
	}

	// Attributes is the current level to which a Hero's attributes (stats) are upgraded.
	Attributes struct {
		Level int `json:"level"`
	}

	// Item represents a Dota 2 item.
	Item struct {
		Name         string   `json:"name"`
		ContainsRune string   `json:"contains_rune"`
		Charges      int      `json:"charges"`
		Cooldown     Duration `json:"cooldown"`
		CanCast      bool     `json:"can_cast"`
		Passive      bool     `json:"passive"`
	}

	// Abilities represents the set of abilities belonging to a hero. A nil slot
	// means that the hero does not have an ability for that slot (unlearned abilities
	// are present but level 0.
	Abilities struct {
		AbilitySlice [6]*Ability
		*Attributes
	}

	// Items represents the entire set of item slots belonging to a hero. A
	// nil entry in either array represents that the slot does not hold an item.
	Items struct {
		Inventory [6]*Item
		Stash     [6]*Item
	}

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

// Q returns the first skill.
func (a Abilities) Q() *Ability {
	return a.AbilitySlice[0]
}

// W returns the second skill.
func (a Abilities) W() *Ability {
	return a.AbilitySlice[1]
}

// E returns the third skill.
func (a Abilities) E() *Ability {
	return a.AbilitySlice[2]
}

// R returns the fourth skill (not necessarily ultimate).
func (a Abilities) R() *Ability {
	return a.AbilitySlice[3]
}

// D returns the fifth skill.
func (a Abilities) D() *Ability {
	return a.AbilitySlice[4]
}

// F returns the sixth skill.
func (a Abilities) F() *Ability {
	return a.AbilitySlice[5]
}

// Ultimate returns the ultimate skill.
func (a Abilities) Ultimate() *Ability {
	for _, v := range a.AbilitySlice {
		if v != nil && v.Ultimate {
			return v
		}
	}
	return nil
}

// WithName checks whether an ability with the given name exists in the list. If so, it returns it.
// Otherwise, it returns nil.
func (a Abilities) WithName(name string) *Ability {
	for _, v := range a.AbilitySlice {
		if v != nil && v.Name == name {
			return v
		}
	}
	return nil
}

// WithName returns all items in the player's inventory and stash that match the given name.
func (i Items) WithName(name string) []*Item {
	var items []*Item
	for _, v := range i.Inventory {
		if v != nil && v.Name == name {
			items = append(items, v)
		}
	}
	for _, v := range i.Stash {
		if v != nil && v.Name == name {
			items = append(items, v)
		}
	}
	return items
}
