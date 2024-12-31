package menus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func teams(_ fyne.Window) fyne.CanvasObject {
	content := container.NewVBox()
	return container.NewCenter(content)
}
