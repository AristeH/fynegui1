package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// / поле ввода
type oEntry struct {
	Ind     *widget.TableCellID
	parent  *TableOtoko
	IDForm  string
	IDTable string
	ID      string
	widget.Entry
}

func (e *oEntry) Tapped(ev *fyne.PointEvent) {
	//t := appValues[e.IDForm].Table[e.IDTable]
	//n := len(t.Data)
	//row := 0
	//for i := 1; i < n; i++ {
	//	if t.Data[i][0] == e.ID {
	//		row = i
	//		break
	//	}
	//}
	//
	//if row == 0 {
	//	sortS(t.Data, e.col)
	//	for i := 1; i < n; i++ {
	//		t.Data[i][1] = strconv.Itoa(i)
	//	}
	//	t.Table.Refresh()
	//}
}
func (e *oEntry) DoubleTapped(ev *fyne.PointEvent) {
	//t := appValues[e.IDForm].Table[e.IDTable]
	//n := len(t.Data)
	//row := 0
	//for i := 1; i < n; i++ {
	//	if t.Data[i][0] == e.ID {
	//		row = i
	//		break
	//	}
	//}
	//
	//if row == 0 {
	//	sortDown(t.Data, e.col)
	//	n := len(t.Data)
	//	for i := 1; i < n; i++ {
	//		t.Data[i][1] = strconv.Itoa(i)
	//	}
	//	t.Table.Refresh()
	//}
}

func (e *oEntry) TappedSecondary(ev *fyne.PointEvent) {

	menuItems := make([]*fyne.MenuItem, 0)
	menuItem := fyne.NewMenuItem(
		"Отбор",
		func() {
			fmt.Println(e.Entry.Text)
		},
	)
	menuItems = append(menuItems, menuItem)
	menuItem = fyne.NewMenuItem(
		"Сортировка",
		func() {
			fmt.Println(e.Entry.Text)
		},
	)
	menuItems = append(menuItems, menuItem)
	widget.ShowPopUpMenuAtPosition(
		fyne.NewMenu("", menuItems...),
		fyne.CurrentApp().Driver().CanvasForObject(e),
		ev.AbsolutePosition,
	)
	//	widget.ShowPopUpMenuAtPosition(fyne.NewMenu("", menuItems...), fyne.CurrentApp().Driver().CanvasForObject(&e.Entry), e.Position())
	fmt.Println(e.Entry.Text)
}

func (e *oEntry) onEnter() {

	id := activeContainer.Selected
	if id.Row == e.Ind.Row {
		activeContainer.Data[id.Row][id.Col] = e.Text
	}
	if len(activeContainer.Data)-1 > activeContainer.Selected.Row {
		activeContainer.Selected = widget.TableCellID{Col: id.Col, Row: id.Row + 1}
		activeContainer.Table.ScrollTo(activeContainer.Selected)
		id := activeContainer.Selected
		e.Entry.Text = activeContainer.Data[id.Row][id.Col]
	}
	fmt.Println(e.Entry.Text)

}

func (e *oEntry) OnChanged(t string) {
	id := activeContainer.Selected
	activeContainer.Data[id.Row][id.Col] = t
	fmt.Println(e.Entry.Text)

}

func newoEntry() *oEntry {
	entry := &oEntry{}

	entry.ExtendBaseWidget(entry)
	entry.Entry.OnChanged = func(sText string) {
		fmt.Println(sText)
	}
	return entry
}

func (e *oEntry) KeyDown(key *fyne.KeyEvent) {
	//t := appValues[e.IDForm].Table[e.IDTable]

	id := activeContainer.Selected
	switch key.Name {
	case fyne.KeyReturn:
		e.onEnter()
	case "KP_Enter":
		e.onEnter()
	case "Down":
		if len(activeContainer.Data)-1 > activeContainer.Selected.Row {
			activeContainer.Selected = widget.TableCellID{Col: id.Col, Row: id.Row + 1}
			activeContainer.Table.ScrollTo(activeContainer.Selected)
		}
	case "Up":
		if id.Row > 1 {
			activeContainer.Selected = widget.TableCellID{Col: id.Col, Row: id.Row - 1}
			activeContainer.Table.ScrollTo(activeContainer.Selected)
		}
	case "Left":
		fmt.Printf("Key %v pressed\n", key.Name)
	case "Right":

		fmt.Printf("Key %v pressed\n", key.Name)

	default:
		//e.Entry.KeyDown(key)
		fmt.Printf("Key %v pressed\n", key.Name)
	}
	//activeContainer.Table.ScrollTo(activeContainer.Selected)
}

func (e *oEntry) KeyUp(key *fyne.KeyEvent) {

	fmt.Printf("Key %v released\n", key.Name)
}
