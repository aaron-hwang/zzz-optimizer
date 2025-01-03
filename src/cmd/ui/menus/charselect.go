package ui

import (
	"errors"
	"zzz-optimizer/pkg/character"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func charSelect(w fyne.Window) fyne.CanvasObject {
	content := container.NewVBox()
	content.Refresh()
	content.Add(widget.NewButton("Add character", func() {
		charCatalog(w)
	}))
	for key, char := range Optimizer.Characters {
		content.Add(widget.NewButton(string(key), func() {
			dialog.ShowCustom(string(key), "x", charView(char), w)
		}))
	}
	return content
}

// Handles the logic of showing the character catalog.
func charCatalog(w fyne.Window) {
	content := container.NewVBox()
	dlg := dialog.NewCustom("Characters", "exit", content, w)
	for key := range character.CharacterCatalog {
		content.Add(widget.NewButton(string(key), func() {
			if _, ok := Optimizer.Characters[key]; ok {
				d := dialog.NewError(errors.New("character already owned"), w)
				d.Show()
			} else {
				Optimizer.Characters[key] = character.NewLevelZeroChar(key)
				dlg.Hide()
			}
		}))
	}
	dlg.Show()
}
