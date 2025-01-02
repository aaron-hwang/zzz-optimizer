package main

import (
	"fmt"
	"strings"
	"zzz-optimizer/cmd/menus"
	"zzz-optimizer/pkg/character"
	"zzz-optimizer/pkg/optimizer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

// The window that is currently on top to be rendered
//var top fyne.Window

var (
	pageContent *fyne.Container
	pageTitle   *widget.Label
)

// Main entry point for the actual program currently
// TODO: Refactor code so that this logic lives in a seperate package (call it app?) to isolate view/controller logic, and so other parts of view
// can communicate with it when necessary
func main() {
	fmt.Println("Hello world")

	// Structure should look like:
	/*
		Run optimizer app
		Optimizer app should accept user input
		(CLI mode basically act like shell? Might not be suited to cli for a prototype, maybe jump straight to desktop app w/gui?)
		Optimizer app should keep track of state of player account, keep track of current characters, weps etc
		Mimic GO design i guess? Initialize
	*/

	// First test should be to register a bunch of characters, have a character management menu where user can press button to open up menu that allows them to add character
	a := app.New()
	w := a.NewWindow("ZZZ Optimizer")
	//top = w

	Optimizer := optimizer.NewOptimizer()
	Optimizer.AddCharacter(character.NewDefaultChar())
	var sb strings.Builder
	for key := range character.CharacterCatalog {
		sb.WriteString(string(key))
	}
	// TODO: USE
	fmt.Println("Characters loaded into string.")

	pageTitle = widget.NewLabel("Component")
	pageContent = container.NewStack()

	setMenu := func(m menus.Menu) {
		pageTitle.SetText(m.Title)
		pageContent.Objects = []fyne.CanvasObject{m.View(w)}
		pageContent.Refresh()
	}

	tutorial := container.NewBorder(
		container.NewVBox(pageTitle, widget.NewSeparator()), nil, nil, nil, pageContent)
	split := container.NewHSplit(makeSidebarNav(setMenu), tutorial)
	split.Offset = 0.2
	w.SetContent(split)
	w.Resize(fyne.NewSize(WIDTH, HEIGHT))
	w.ShowAndRun()
}

// This should handle creating the structure for the side menu where users will navigate between pages (team management, character management, disk drives, etc)
func makeSidebarNav(setMenu func(menu menus.Menu)) fyne.CanvasObject {
	// TODO: Implement
	//a := fyne.CurrentApp()
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return menus.MenuIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := menus.MenuIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) (o fyne.CanvasObject) {
			return widget.NewLabel("test createnode")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			m, ok := menus.Menus[uid]
			if !ok {
				fyne.LogError("Missing a menu: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(m.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := menus.Menus[uid]; ok {
				setMenu(t)
			}
		},
	}
	return container.NewBorder(nil, nil, nil, tree)
}
