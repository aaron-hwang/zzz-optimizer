package menus

import (
	"zzz-optimizer/pkg/character"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// This page should handle the display of the user's current state of their characters.
func charView(_ character.Character) fyne.CanvasObject {
	content := container.NewVBox()
	return content
}
