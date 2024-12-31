package menus

import (
	"fmt"
	"zzz-optimizer/pkg/character"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func charSelect(_ fyne.Window) fyne.CanvasObject {
	content := container.NewVBox()
	for key := range character.CharacterCatalog {
		content.Add(widget.NewLabel(string(key)))
	}
	content.Add(widget.NewButton("Button", func() {fmt.Println("pressed!")}))
	return container.NewCenter(content)
}
