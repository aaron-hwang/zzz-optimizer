package dummy

import (
	"zzz-optimizer/pkg/character"
	"zzz-optimizer/pkg/key"
)

func init() {
	character.Register(key.Dummy, character.CharacterConfig{})
}
