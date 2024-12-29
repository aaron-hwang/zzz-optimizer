package miyabi

import (
	"zzz-optimizer/pkg/character"
	"zzz-optimizer/pkg/key"
)

// Init a character, by registering them into the character catalog for use (Essentially affirm that 1. they exist, and 2. they have xyz skills)
func init() {
	character.Register(key.Miyabi, character.CharacterConfig{})
}
