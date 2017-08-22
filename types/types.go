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
	Strength            StatPair     `json:"strength"`
	Dexterity           StatPair     `json:"dexterity"`
	Constitution        StatPair     `json:"constitution"`
	Intelligence        StatPair     `json:"intelligence"`
	Wisdom              StatPair     `json:"wisdom"`
	Charisma            StatPair     `json:"charisma"`
	Challenge           int          `json:"challenge"`
	ArmorClass          int          `json:"armor_class"`
	HitPoints           int          `json:"hit_points"`
	Speed               int          `json:"speed"`
	DamageResistances   []resistance `json:"damage_resistances"`
	DamageImmunities    []immunity   `json:"damage_immunities"`
	ConditionImmunities []immunity   `json:"condition_immunities"`
	Sense               []sense      `json:"sense"`
	Languages           []string     `json:"languages"`
	Habilities          []hability   `json:"habilities"`
	Actions             []action     `json:"actions"`
	Description         string       `json:"description"`
	Picture             string       `json:"picture"`
}

// StatPair describes a base attribute (such as strength) and its modifier
type StatPair struct {
	Attribute int `json:"attribute"`
	Modifier  int `json:"modifier"`
}

// TODO: further specify habilities, resistances, immunities and senses

type resistance struct{}
type immunity struct{}
type sense struct{}
type hability struct{}
type action struct{}

// Result describes a an HTTP call response when fetching monsters
type Result struct {
	Items       []MonsterEntity
	TotalCount  int64
	CurrentPage int
	TotalPages  int
}
