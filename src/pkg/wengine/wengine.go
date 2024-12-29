package wengine

import "zzz-optimizer/pkg/character"

var (
	wengineCatalog map[string]WengineConfig
)

type Wengine struct {
	Level      int
	Refinement int
	Ascension  int
	// a lil sus but should be fine
	Type character.CharacterRole
}

type WengineRarity int

const (
	BRank WengineRarity = 3
	ARank WengineRarity = 4
	SRank WengineRarity = 5
)

type WengineConfig struct {
	Rarity WengineRarity
}

// Register a wengine config (assert that a wengine of this type can exist essentially)
func Register(key string, config WengineConfig) {
	wengineCatalog[key] = config
}