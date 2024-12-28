package main

import (
	"fmt"
	"zzz-optimizer/pkg/character"
	"zzz-optimizer/pkg/optimizer"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// Main entry point for the actual program currently
func main() {
	fmt.Println("Hello world")

	// Structure should look like:
	/*
		Run optimizer app
		Optimizer app should accept user input
		(CLI mode basically act like shell? Might not be suited to cli for a prototype, maybe jump straight to desktop app w/gui?)
		Optimizer app should keep track of state of player account, keep track of current characters, weps etc
		Mimic GO design i guess? Initialize
	*/
	a := app.New()
	w := a.NewWindow("ZZZ Optimizer")
	Optimizer := optimizer.NewOptimizer()
	Optimizer.AddCharacter(character.NewDefaultChar())
	str := fmt.Sprintf("%#v", Optimizer)
	w.SetContent(widget.NewLabel(str))
	w.ShowAndRun()
}
