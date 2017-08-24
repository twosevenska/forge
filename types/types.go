package types

import (
	"time"

	"github.com/satori/go.uuid"
)

// BaseEntity describes a generic MongoDB row
type BaseEntity struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id"`
	Name      string    `bson:"name" json:"name"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// MonsterEntity describes the general stats of a Monster
type MonsterEntity struct {
	BaseEntity
	Strength            StatPair     `bson:"strength" json:"strength"`
	Dexterity           StatPair     `bson:"dexterity" json:"dexterity"`
	Constitution        StatPair     `bson:"constitution" json:"constitution"`
	Intelligence        StatPair     `bson:"intelligence" json:"intelligence"`
	Wisdom              StatPair     `bson:"wisdom" json:"wisdom"`
	Charisma            StatPair     `bson:"charisma" json:"charisma"`
	Challenge           int          `bson:"challenge" json:"challenge"`
	ArmorClass          int          `bson:"armor_class" json:"armor_class"`
	HitPoints           int          `bson:"hit_points" json:"hit_points"`
	Speed               int          `bson:"speed" json:"speed"`
	DamageResistances   []resistance `bson:"damage_resistances" json:"damage_resistances"`
	DamageImmunities    []immunity   `bson:"damage_immunities" json:"damage_immunities"`
	ConditionImmunities []immunity   `bson:"condition_immunities" json:"condition_immunities"`
	Sense               []sense      `bson:"sense" json:"sense"`
	Languages           []string     `bson:"languages" json:"languages"`
	Habilities          []hability   `bson:"habilities" json:"habilities"`
	Actions             []action     `bson:"actions" json:"actions"`
	Description         string       `bson:"description" json:"description"`
	Picture             string       `bson:"picture" json:"picture"`
}

// StatPair describes a base attribute (such as strength) and its modifier
type StatPair struct {
	Attribute int `bson:"attribute" json:"attribute"`
	Modifier  int `bson:"modifier" json:"modifier"`
}

// TODO: further specify habilities, resistances, immunities and senses

type resistance struct{}
type immunity struct{}
type sense struct{}
type hability struct{}
type action struct{}

// BaseResult a generic HTTP call response when fetching data
type BaseResult struct {
	Items       []interface{}
	TotalCount  int64
	CurrentPage int
	TotalPages  int
}

// MonsterResult describes a an HTTP call response when fetching monsters
type MonsterResult struct {
	BaseResult
	Items []MonsterEntity
}
