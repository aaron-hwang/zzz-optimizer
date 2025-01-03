package ui

import (
	"strconv"
	"zzz-optimizer/pkg/character"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// This page should handle the display of the user's current state of their given character.
func charView(char character.Character) fyne.CanvasObject {
	content := container.NewVBox()
	content.Add(widget.NewLabel(string(char.Key)))
	content.Add(widget.NewLabel("Core Level: " + strconv.Itoa(char.SkillLevels.Core)))
	return content
}
