package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	welcomeMessage = "This is the ZZZ optimizer. Similar to GO for Genshin Impact, utilize this program to optimize drive disk and gear loadouts for character in various different teams. Currently a work in progress."
)

func homeMenu(_ fyne.Window) fyne.CanvasObject {
	content := container.NewVBox()
	content.Add(widget.NewLabel(welcomeMessage))
	return content
}
