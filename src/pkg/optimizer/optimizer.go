package optimizer

import (
	"zzz-optimizer/pkg/character"
	"zzz-optimizer/pkg/wengine"
)

// Can basically think of this as the C in MVC
// Stitch together the model to do stuff, pass stuff to view(?)

// Essentially a container for the current state of a user locally
type Optimizer struct {
	// Need a list of characters the current user has
	Characters []character.Character
	// list of all the wengines they own
	Wengines []wengine.Wengine
	// A list of all the disks they own

	// A list of all the bangboos they own
}

func NewOptimizer() *Optimizer {
	opti := Optimizer{Characters: make([]character.Character, 1), Wengines: make([]wengine.Wengine, 1)}
	return &opti
}

func (o *Optimizer) AddCharacter(c character.Character) {
	o.Characters = append(o.Characters, c)
}

type UserConfig struct {
}

func (o *Optimizer) Load(config UserConfig) error {

	return nil
}
