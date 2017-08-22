package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// BaseEntity describes a generic MongoDB row
type BaseEntity struct {
	ID        bson.Binary `bson:"_id,omitempty" json:"id"`
	Name      string      `bson:"name" json:"name"`
	CreatedAt time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at" json:"updated_at"`
}

// MonsterEntity describes the general stats of a Monster
type MonsterEntity struct {
	BaseEntity
	Strength            StatPair
	Dexterity           StatPair
	Constitution        StatPair
	Intelligence        StatPair
	Wisdom              StatPair
	Charisma            StatPair
	Challenge           int
	ArmorClass          int
	HitPoints           int
	Speed               int
	DamageResistances   []resistance
	DamageImmunities    []immunity
	ConditionImmunities []immunity
	Sense               []sense
	Languages           []string
	Habilities          []hability
	Actions             []action
	Description         string
	Picture             string
}

// StatPair describes a base attribute (such as strength) and its modifier
type StatPair struct {
	Attribute int
	Modifier  int
}

// TODO: further specify habilities, resistances, immunities and senses

type resistance struct{}
type immunity struct{}
type sense struct{}
type hability struct{}
type action struct{}

// MonsterResult describes a an HTTP call response when fetching monsters
type MonsterResult struct {
	Items       []MonsterEntity
	TotalCount  int64
	CurrentPage int
	TotalPages  int
}
