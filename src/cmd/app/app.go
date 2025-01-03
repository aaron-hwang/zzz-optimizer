package optimizer_app

import (
	ui "zzz-optimizer/cmd/ui/menus"
)

// Logic goes here
func Run() {
	ui := ui.NewGui()
	ui.Display()
}
