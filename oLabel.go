package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
	"strings"
)

// / поле ввода
type oLabel struct {
	IDForm  string
	IDTable string
	ID      string
	Ind     *widget.TableCellID
	widget.Label
	parent *TableOtoko
}

func sortS(x [][]string, k int) {
	var temp []string
	n := len(x)
	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			if strings.ToUpper(x[i][k]) > strings.ToUpper(x[j][k]) {
				temp = x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
}

func sortDown(x [][]string, k int) {
	var temp []string
	n := len(x)
	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			if strings.ToUpper(x[i][k]) < strings.ToUpper(x[j][k]) {
				temp = x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
}

func (e *oLabel) Tapped(ev *fyne.PointEvent) {

	id := e.Ind
	activeContainer = e.parent

	if id.Row == 0 {
		sortS(e.parent.Data, id.Col)
		for i := 1; i < len(e.parent.Data); i++ {
			e.parent.Data[i][1] = strconv.Itoa(i)
		}

	}
	activeContainer.Selected = *id
	e.parent.Table.Refresh()

}

func (e *oLabel) DoubleTapped(ev *fyne.PointEvent) {
	ind := e.Ind
	if ind.Row == 0 {
		sortDown(e.parent.Data, ind.Row)
		n := len(e.parent.Data)
		for i := 1; i < n; i++ {
			e.parent.Data[i][1] = strconv.Itoa(i)
		}
		e.parent.Table.Refresh()
	}

	items := make([]*widget.FormItem, 0)
	for col, style := range e.parent.ColumnStyle {
		if style.Width != 0 {
			Entry := widget.NewEntry()
			Entry.Validator = getValidator(style.Type)
			Entry.Text = e.parent.Data[ind.Row][col]
			items = append(items, widget.NewFormItem(style.Name, Entry))

		}
	}
	dialog.ShowForm("введите", "", "cancel", items, func(b bool) {
		if !b {
			return
		}

	}, appValues["main"].W)

}

func (e *oLabel) TappedSecondary(ev *fyne.PointEvent) {
	ind := e.Ind
	Entry := widget.NewEntry()
	Entry.Validator = getValidator(e.parent.ColumnStyle[ind.Col].Type)
	Entry.Text = e.parent.Data[ind.Row][ind.Col]
	items := []*widget.FormItem{
		widget.NewFormItem(e.parent.ColumnStyle[ind.Col].Name, Entry),
	}
	dialog.ShowForm("введите", "", "cancel", items, func(b bool) {
		if !b {
			return
		}
		fmt.Println("KP_Enter", Entry.Text)
		e.parent.Data[ind.Row][ind.Col] = Entry.Text
	}, appValues["main"].W)

}

func (e *oLabel) onEnter() {
	fmt.Println(e.Text)

}

func (e *oLabel) Focusable(key *fyne.KeyEvent) {
	fmt.Printf("Key %v released\n", key.Name)
}

func (e *oLabel) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		println(s)
		return
	}
	log.Println("Shortcut typed:", s)
}

func newOLabel() *oLabel {
	entry := &oLabel{}
	entry.ExtendBaseWidget(entry)
	return entry
}
