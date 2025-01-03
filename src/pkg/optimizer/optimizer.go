package optimizer

import (
	"errors"
	"zzz-optimizer/pkg/character"
	"zzz-optimizer/pkg/key"
	"zzz-optimizer/pkg/wengine"
)

// Can basically think of this as the C in MVC
// Stitch together the model to do stuff, pass stuff to view(?)

// Essentially a container for the current state of a user locally
type Optimizer struct {
	// Need a list of characters the current user has
	Characters map[key.Character]character.Character
	// list of all the wengines they own
	Wengines []wengine.Wengine
	// A list of all the disks they own

	// A list of all the bangboos they own
}

func NewOptimizer() *Optimizer {
	opti := Optimizer{Characters: make(map[key.Character]character.Character), Wengines: make([]wengine.Wengine, 1)}
	return &opti
}

func (o *Optimizer) AddCharacter(c character.Character) error {
	// Check if character with same key currently exists
	// If so, return an error
	_, exists := o.Characters[c.Key]
	if exists {
		// Should realistically never happen, edge case
		return errors.New("cannot add a character that already exists")
	}

	o.Characters[c.Key] = c
	return nil
}

type UserConfig struct {
}

func (o *Optimizer) Load(config UserConfig) error {

	return nil
}
