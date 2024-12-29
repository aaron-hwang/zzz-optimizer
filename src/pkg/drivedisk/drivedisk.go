package drivedisk

import "zzz-optimizer/pkg/key"

var (

	// TODO: make thread safe (add a mutex)
	diskCatalog map[key.DriveDisk]DiskConfig
)

type DriveDisk struct {
	Set key.DriveDisk
}

type DiskRarity int

const (
	BRank DiskRarity = 3
	ARank DiskRarity = 4
	SRank DiskRarity = 5
)

type DiskConfig struct {
	Rarity DiskRarity
}

func Register(key key.DriveDisk, config DiskConfig) {
	diskCatalog[key] = config
}
