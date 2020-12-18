package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var content *widget.Entry

// var currentButton *widget.Button

func setNote(n *note) {
	content.SetText(n.content)
}

func loadUI(l *notelist) fyne.CanvasObject {

	list := widget.NewVBox()

	for _, n := range l.notes {
		note := n
		item := widget.NewButton(n.title(), func() {
			setNote(note)
		})
		list.Append(item)
	}

	content = widget.NewMultiLineEntry()
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			note := l.add()
			note.content = content.Text
			setNote(note)
			newItem := widget.NewButton(note.title(), func() {
				setNote(note)
			})
			list.Append(newItem)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			list.Children[len(list.Children)-1].Hide()
			list.Children = list.Children[:len(list.Children)-1]
		}),
	)

	if len(l.notes) > 0 {
		content.SetText(l.notes[0].content)
	}

	left := fyne.NewContainerWithLayout(layout.NewBorderLayout(bar, nil, nil, nil), bar, list)

	ui := widget.NewHSplitContainer(left, content)
	ui.Offset = 0.25
	return ui
}

func main() {
	a := app.New()
	w := a.NewWindow("Notes")

	notes := &notelist{[]*note{
		{content: "My first note"},
		{content: "This note has\n many lines"},
	}}

	w.SetContent(loadUI(notes))
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
