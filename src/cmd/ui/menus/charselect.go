package ui

import (
	"zzz-optimizer/pkg/character"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func charSelect(w fyne.Window) fyne.CanvasObject {
	content := container.NewVBox()
	for key := range character.CharacterCatalog {
		content.Add(widget.NewButton(string(key), func() {
			dialog.ShowCustom(string(key), "x", charView(key), w)
		}))
	}
	return content
}
