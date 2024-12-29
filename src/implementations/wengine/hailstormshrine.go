package hailstormshrine

import (
	"zzz-optimizer/pkg/key"
	"zzz-optimizer/pkg/wengine"
)

func init() {
	wengine.Register(key.HailstormShrine, wengine.WengineConfig{})
}
