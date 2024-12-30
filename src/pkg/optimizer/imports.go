package optimizer

// This file exists to import all character implementation files. This ensures that all "init()" functions in a character's package are loaded, which
// ensures that they are registered in the character catalog.

import (
	_ "zzz-optimizer/implementations/character/miyabi"
)
