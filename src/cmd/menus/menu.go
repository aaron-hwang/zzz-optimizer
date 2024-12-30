package menus

import "fyne.io/fyne/v2"

type Menu struct {
	View func(w fyne.Window) fyne.CanvasObject
}
