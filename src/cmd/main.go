package main

import (
	optimizer_app "zzz-optimizer/cmd/app"
)

// Main entry point for the actual program currently
// TODO: Refactor code so that this logic lives in a seperate package (call it app?) to isolate view/controller logic, and so other parts of view
// can communicate with it when necessary
func main() {
	// Can intercept cli args if we want to run in debug mode to test certain things
	optimizer_app.Run()
}
