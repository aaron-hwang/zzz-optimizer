package ui

import "fyne.io/fyne/v2"

type Menu struct {
	Title string
	View  func(w fyne.Window) fyne.CanvasObject
}

var (
	Menus = map[string]Menu{
		"home":       {"Home", homeMenu},
		"characters": {"Characters", charSelect},
		"teams":      {"Teams", teams},
	}

	MenuIndex = map[string][]string{
		"": {"characters", "teams", "home"},
	}
)
