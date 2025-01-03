package ui

import (
	"zzz-optimizer/pkg/key"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// This page should handle the display of the user's current state of their characters.
func charView(char key.Character) fyne.CanvasObject {
	content := container.NewVBox()
	character := Optimizer.Characters[char]
	content.Add(widget.NewLabel(string(character.Key)))
	return content
}
