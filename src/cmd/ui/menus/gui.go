package ui

import (
	"fmt"
	"strings"
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

// i don't wanna talk about it
var (
	Optimizer *optimizer.Optimizer = optimizer.NewOptimizer()
)

type Gui struct {
	pageContent *fyne.Container
	pageTitle   *widget.Label
	window      fyne.Window
}

func NewGui() *Gui {
	new := Gui{}
	return &new
}

func (g *Gui) Display() {
	a := app.New()
	g.window = a.NewWindow("ZZZ Optimizer")
	//top = w

	Optimizer.AddCharacter(character.NewDefaultChar())
	var sb strings.Builder
	for key := range character.CharacterCatalog {
		sb.WriteString(string(key))
	}
	// TODO: USE
	fmt.Println("Characters loaded into string.")

	g.pageTitle = widget.NewLabel("Component")
	g.pageContent = container.NewStack()

	tutorial := container.NewBorder(
		container.NewVBox(g.pageTitle, widget.NewSeparator()), nil, nil, nil, g.pageContent)
	split := container.NewHSplit(g.makeSidebarNav(), tutorial)
	split.Offset = 0.2
	g.window.SetContent(split)
	g.window.Resize(fyne.NewSize(WIDTH, HEIGHT))
	g.window.ShowAndRun()
}

func (g *Gui) SetMenu(menu Menu) {
	g.pageTitle.SetText(menu.Title)
	g.pageContent.Objects = []fyne.CanvasObject{menu.View(g.window)}
	g.pageContent.Refresh()
}

// This should handle creating the structure for the side menu where users will navigate between pages (team management, character management, disk drives, etc)
func (g *Gui) makeSidebarNav() fyne.CanvasObject {
	// TODO: Implement
	//a := fyne.CurrentApp()
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return MenuIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := MenuIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) (o fyne.CanvasObject) {
			return widget.NewLabel("test createnode")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			m, ok := Menus[uid]
			if !ok {
				fyne.LogError("Missing a menu: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(m.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := Menus[uid]; ok {
				g.SetMenu(t)
			}
		},
	}
	return container.NewBorder(nil, nil, nil, tree)
}
