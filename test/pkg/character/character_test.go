package character

import (
	"testing"
)

// Test that the character level zero init function has the following behavior:
/**
The function should take in a key.Character, checking if that key exists in the catalog (and thus has a config). This is mostly a formality because
REALISTICALLY a character with a key will likely have an implementation. If the key does not have a config associated, then we should properly
catch the error.
It should also test if the character properly spawns in at level 0, with all 0'd out skils, but with proper element and faction data etc.
If
**/
func testCharacterLevelZero(t *testing.T) {
	// TODO: Testing
}
