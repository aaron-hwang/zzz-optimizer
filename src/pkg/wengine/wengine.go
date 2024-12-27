package wengine

var (
	wengineCatalog map[string]WengineConfig
)

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
