package gogsi

import "time"

type (
	// Hero represents the hero being controlled by the connected client.
	Hero struct {
		ID              int           `json:"id"`
		Name            string        `json:"name"`
		Level           int           `json:"level"`
		Alive           bool          `json:"alive"`
		Respawn         time.Duration `json:"respawn_seconds"`
		BuybackCost     int           `json:"buyback_cost"`
		BuybackCooldown time.Duration `json:"buyback_cooldown"`
		Health          int           `json:"health"`
		MaxHealth       int           `json:"max_health"`
		HealthPercent   int           `json:"health_percent"`
		Mana            int           `json:"mana"`
		MaxMana         int           `json:"max_mana"`
		ManaPercent     int           `json:"mana_percent"`
		Silenced        bool          `json:"silenced"`
		Stunned         bool          `json:"stunned"`
		Disarmed        bool          `json:"disarmed"`
		MagicImmune     bool          `json:"magicimmune"`
		Hexed           bool          `json:"hexed"`
		Muted           bool          `json:"muted"`
		Break           bool          `json:"break"`
		Debuffed        bool          `json:"has_debuff"`
	}

	// Ability represents a Hero's ability.
	Ability struct {
		Name     string        `json:"name"`
		Level    int           `json:"level"`
		Cooldown time.Duration `json:"cooldown"`
		CanCast  bool          `json:"can_cast"`
		Active   bool          `json:"ability_active"`
		Passive  bool          `json:"passive"`
		Ultimate bool          `json:"ultimate"`
	}

	// Attributes is the current level to which a Hero's attributes (stats) are upgraded.
	Attributes struct {
		Level int `json:"level"`
	}

	// Item represents a Dota 2 item.
	Item struct {
		Name         string        `json:"name"`
		ContainsRune string        `json:"contains_rune"`
		Charges      int           `json:"charges"`
		Cooldown     time.Duration `json:"cooldown"`
		CanCast      bool          `json:"can_cast"`
		Passive      bool          `json:"passive"`
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
)

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
