package character

import "zzz-optimizer/pkg/model"

var (
	CharacterCatalog map[string]CharacterConfig
)

func Register(key string, config CharacterConfig) {
	CharacterCatalog[key] = config
}

// Why am i seperating config and instances? Who knows!!!! This might just end up being redundant......
type Character struct {
	Level           int
	Ascension       int
	MindscapeCinema int
	SkillLevels     SkillLevels
}

type SkillLevels struct {
	Core        int
	Basic       int
	Special     int
	Assists     int
	UltAndChain int
}

// Mostly for debugging
func NewDefaultChar() Character {
	char := Character{
		Level:           1,
		Ascension:       0,
		MindscapeCinema: 0,
		SkillLevels:     SkillLevels{Core: 1, Basic: 1, Special: 1, Assists: 1, UltAndChain: 1}}
	return char
}

type CharacterRarity int

const (
	ARank CharacterRarity = 4
	SRank CharacterRarity = 5
)

type CharacterRole int

const (
	Role_INVALID CharacterRole = -1
	Role_ATTACK  CharacterRole = 0
	Role_ANOMALY CharacterRole = 1
	Role_SUPPORT CharacterRole = 2
	Role_DEFENSE CharacterRole = 3
	Role_STUN    CharacterRole = 4
)

type CharacterConfig struct {
	Rarity     CharacterRarity
	Promotions []PromotionData
	// Probably make an enum for this at some point
	Faction string
	// Ditto
	Element   string
	Role      CharacterRole
	SkillInfo SkillInfo
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
	BasicAtk    BasicAttack
	Special     Special
	DodgeAttack DodgeAttack
	Assist      Assist
	Ult         Ult
	ChainAttack ChainAttack
	Core        CoreSkill
	Additional  AdditionalAbility
}

type BasicAttack struct {
	TargetType model.TargetType
}

type DodgeAttack struct {
	TargetType model.TargetType
}

type Special struct {
	Cost       int
	TargetType model.TargetType
}

type Ult struct {
	TargetType model.TargetType
}

type ChainAttack struct {
	TargetType model.TargetType
}

type AssistType int

const (
	AssistType_NULL      AssistType = 0
	ASsistType_DEFENSIVE AssistType = 1
	AssistType_EVASIVE   AssistType = 2
)

type Assist struct {
	AssistType AssistType
	TargetType model.TargetType
}

type CoreSkill struct {
}

type AdditionalAbility struct {
}
