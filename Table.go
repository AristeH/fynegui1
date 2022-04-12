package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"
	"strings"

	//"sort"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TableOtoko struct {
	ID            string
	IDForm        string
	ColumnsName   []string
	ColumnsFormat []string
	ColumnsType   []string
	ColumnsWidth  []float32
	HeaderColor   color.Color
	AlterRowColor color.Color
	RowColor      color.Color
	Data          [][]string
	Edit          bool
	Tool          *widget.Toolbar
	Table         *widget.Table
	we            map[widget.TableCellID]*enterEntry
	wc            map[widget.TableCellID]*enterCheck
}

func sortS(x [][]string, k int) {
	var temp []string
	n := len(x)
	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			if strings.ToUpper( x[i][k]) > strings.ToUpper(x[j][k]) {
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

func (t *TableOtoko) makeTable() *fyne.Container {
	t.Table = widget.NewTable(
		func() (int, int) {
			rows := len(t.Data)
			columns := len(t.Data[0])
			return rows, columns
		},
		func() fyne.CanvasObject {
			con := container.NewHBox()
			con.Layout = layout.NewMaxLayout()
			check := newenterCheck()
			// обработка нажатия на чек бокс
			check.OnChanged = func(b bool) {
				t := app_values[t.IDForm].Table[t.ID]
				n := len(t.Data)
				row := 0
				for i := 1; i < n; i++ {
					if t.Data[i][0] == check.ID {
						row = i
						break
					}
				}
				if check.Checked {
					t.Data[row][check.col] = "true"
				} else {
					t.Data[row][check.col] = "false"
				}
				// направление движения по столбцу вниз в дальнейшем условие
				newTableCellID := widget.TableCellID{Col: check.col, Row: row + 1}
				t.Table.ScrollTo(newTableCellID)

			}
			check.IDForm = t.IDForm
			check.IDTable = t.ID
			con.Add(check)

			entry := newEnterEntry()
			entry.IDForm = t.IDForm
			entry.IDTable = t.ID
			con.Add(entry)

			return container.New(layout.NewMaxLayout(),
				canvas.NewRectangle(color.Gray{250}),
				con,
			)

		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			var ic *enterCheck
			var entry *enterEntry
			box := o.(*fyne.Container)
			rect := box.Objects[0].(*canvas.Rectangle)
			if i.Row == 0 {
				rect.FillColor = t.HeaderColor
			} else if i.Row%2 == 0 {
				rect.FillColor = t.AlterRowColor
			} else {
				rect.FillColor = t.RowColor
			}
			cont := box.Objects[1].(*fyne.Container)
			ic = cont.Objects[0].(*enterCheck)
			entry = cont.Objects[1].(*enterEntry)

			ic.Hidden = true

			entry.Hidden = true

			if i.Row == 0 {
				entry.SetText(t.Data[i.Row][i.Col])
				entry.col = i.Col
				entry.ID = t.Data[i.Row][0]
				entry.Hidden = false

			} else {

				switch t.ColumnsType[i.Col] {
				case "bool":
					if t.Data[i.Row][i.Col] == "true" {
						ic.Checked = true
					} else {
						ic.Checked = false
					}
					ic.ID = t.Data[i.Row][0]
					ic.col = i.Col
					ic.Refresh()
					ic.Hidden = false
					app_values[t.IDForm].Table[t.ID].wc[i] = ic
				default:
					entry.SetText(t.Data[i.Row][i.Col])
					app_values[t.IDForm].Table[t.ID].we[i] = entry
					entry.col = i.Col
					entry.ID = t.Data[i.Row][0]
					entry.Hidden = false
				}
			}
		})
	for ic, v := range t.ColumnsWidth {
		t.Table.SetColumnWidth(ic, v)
	}

	t.Tool = widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {

			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			log.Println("Display help")
		}))

	content := container.NewBorder(
		container.NewVBox(
			t.Tool,
			widget.NewSeparator(),
		),
		nil,
		nil,
		nil,

		t.Table,
	)
	return content

}

/// поле ввода
type enterEntry struct {
	IDForm  string
	IDTable string
	ID      string
	col     int
	widget.Entry
}

func (e *enterEntry) Tapped(ev *fyne.PointEvent) {
	t := app_values[e.IDForm].Table[e.IDTable]
	n := len(t.Data)
	row := 0
	for i := 1; i < n; i++ {
		if t.Data[i][0] == e.ID {
			row = i
			break
		}
	}

	if row == 0 {
		sortS(t.Data, e.col)
		for i := 1; i < n; i++ {
			t.Data[i][1] = strconv.Itoa(i)
		}
		t.Table.Refresh()
	}
}
func (e *enterEntry) DoubleTapped(ev *fyne.PointEvent) {
	t := app_values[e.IDForm].Table[e.IDTable]
	n := len(t.Data)
	row := 0
	for i := 1; i < n; i++ {
		if t.Data[i][0] == e.ID {
			row = i
			break
		}
	}

	if row == 0 {
		sortDown(t.Data, e.col)
		n := len(t.Data)
		for i := 1; i < n; i++ {
			t.Data[i][1] = strconv.Itoa(i)
		}
		t.Table.Refresh()
	}
}

func (e *enterEntry) TappedSecondary(ev *fyne.PointEvent) {
	fmt.Println(e.Entry.Text)
}

func (e *enterEntry) onEnter() {
	fmt.Println(e.Entry.Text)

}

func newEnterEntry() *enterEntry {
	entry := &enterEntry{}

	entry.ExtendBaseWidget(entry)
	return entry
}

func scrolltable(row int, col int, t *TableOtoko) {
	switch t.ColumnsType[col] {
	case "bool":
		newTableCellID := widget.TableCellID{Col: col, Row: row}
		t.Table.ScrollTo(newTableCellID)
		key := t.wc[newTableCellID]
		app_values[t.IDForm].W.Canvas().Focus(key)

	default:
		newTableCellID := widget.TableCellID{Col: col, Row: row}
		t.Table.ScrollTo(newTableCellID)
		key := t.we[newTableCellID]
		app_values[t.IDForm].W.Canvas().Focus(key)

	}

}

func (e *enterEntry) KeyDown(key *fyne.KeyEvent) {
	t := app_values[e.IDForm].Table[e.IDTable]
	n := len(t.Data)
	row := 0
	for i := 1; i < n; i++ {
		if t.Data[i][0] == e.ID {
			row = i
			break
		}
	}
	switch key.Name {
	case fyne.KeyReturn:
		e.onEnter()
	case "KP_Enter":
		e.onEnter()
	case "Down":
		if n == row+1 {
			scrolltable(row, e.col, t)
		} else {
			scrolltable(row+1, e.col, t)
		}
	case "Left":
		if e.col == 0 {
			scrolltable(row, e.col, t)
		} else {
			scrolltable(row, e.col-1, t)
		}
	case "Right":

		if len(t.ColumnsType) == e.col+1 {
			scrolltable(row, e.col, t)
		} else {
			scrolltable(row, e.col+1, t)
		}

	case "Up":
		if row == 0 {
			scrolltable(row, e.col, t)
		} else {
			scrolltable(row-1, e.col, t)
		}
	default:
		//e.Entry.KeyDown(key)
		fmt.Printf("Key %v pressed\n", key.Name)
	}
}

func (e *enterEntry) KeyUp(key *fyne.KeyEvent) {
	fmt.Printf("Key %v released\n", key.Name)
}

/// чек бокс
type enterCheck struct {
	IDForm  string
	IDTable string
	ID      string
	col     int
	widget.Check
}

func (e *enterCheck) onEnter() {
	fmt.Println(e.Check.Text)

}

func newenterCheck() *enterCheck {
	entry := &enterCheck{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *enterCheck) KeyDown(key *fyne.KeyEvent) {
	t := app_values[e.IDForm].Table[e.IDTable]
	n := len(t.Data)
	row := 0
	for i := 1; i < n; i++ {
		if t.Data[i][0] == e.ID {
			row = i
			break
		}
	}
	switch key.Name {
	case fyne.KeyReturn:
		e.onEnter()
	case "KP_Enter":
		e.onEnter()
case "Down":
		if n == row+1 {
			scrolltable(row, e.col, t)
		} else {
			scrolltable(row+1, e.col, t)
		}
	case "Left":
		if e.col == 0 {
			scrolltable(row, e.col, t)
		} else {
			scrolltable(row, e.col-1, t)
		}
	case "Right":

		if len(t.ColumnsType) == e.col+1 {
			scrolltable(row, e.col, t)
		} else {
			scrolltable(row, e.col+1, t)
		}

	case "Up":
		if row == 0 {
			scrolltable(row, e.col, t)
		} else {
			scrolltable(row-1, e.col, t)
		}
	default:

		fmt.Printf("Key %v pressed\n", key.Name)
	}
}

func (e *enterCheck) KeyUp(key *fyne.KeyEvent) {
	fmt.Printf("Key %v released\n", key.Name)
}
