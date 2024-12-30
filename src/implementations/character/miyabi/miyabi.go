package miyabi

import (
	"fmt"
	"zzz-optimizer/pkg/character"
	"zzz-optimizer/pkg/key"
)

// Init a character, by registering them into the character catalog for use (Essentially affirm that 1. they exist, and 2. they have xyz skills)
func init() {
	fmt.Println("Initing miyabi")
	character.Register(key.Miyabi, character.CharacterConfig{})
}
