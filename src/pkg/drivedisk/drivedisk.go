package drivedisk

var (
	diskCatalog map[string]DiskConfig
)

type DiskRarity int

const (
	BRank DiskRarity = 3
	ARank DiskRarity = 4
	SRank DiskRarity = 5
)

type DiskConfig struct {
	Rarity DiskRarity
}

func Register(key string, config DiskConfig) {
	diskCatalog[key] = config
}
