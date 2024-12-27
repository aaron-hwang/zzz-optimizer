package character

import "zzz-optimizer/pkg/model"

var ()

type CharacterRarity int

const (
	ARank CharacterRarity = 4
	SRank CharacterRarity = 5
)

type CharacterConfig struct {
	Rarity     CharacterRarity
	Promotions []PromotionData
	// Probably make an enum for this at some point
	Faction string
	// Ditto
	Element string
}

// Characters all have the following (Promotion related) data
type PromotionData struct {
	MaxLevel int
	BaseHp   float64
	BaseDef  float64
	BaseAtk  float64
	// Anomaly Mastery
	BaseAM float64
	// Anomaly Proficiency
	BaseAP float64
}

type SkillInfo struct {
	BasicAtk BasicAttack
}

type BasicAttack struct {
	TargetType model.TargetType
}

type Skill struct {
	Cost       int
	TargetType model.TargetType
}
