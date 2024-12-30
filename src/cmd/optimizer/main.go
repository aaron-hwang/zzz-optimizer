package main

import (
	"fmt"
	"image/color"
	"strings"
	"zzz-optimizer/pkg/character"
	"zzz-optimizer/pkg/optimizer"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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

	// First test should be to register a bunch of characters, have a character management menu where user can press button to open up menu that allows them to add character
	a := app.New()
	w := a.NewWindow("ZZZ Optimizer")
	Optimizer := optimizer.NewOptimizer()
	Optimizer.AddCharacter(character.NewDefaultChar())
	var sb strings.Builder
	for key := range character.CharacterCatalog {
		sb.WriteString(string(key))
	}
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText(sb.String(), green)
	content := container.NewWithoutLayout(text1)
	//str := fmt.Sprintf("%#v", Optimizer)
	w.SetContent(content)
	w.ShowAndRun()
}
