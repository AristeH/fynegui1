package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"sort"
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
	Header        *fyne.Container
	we            map[*enterEntry]widget.TableCellID
	wc            map[*widget.Check]widget.TableCellID
	wb            map[*widget.Button]int
}

func (t *TableOtoko) CreateHeader() {
	t.Header = container.New(&ToolButton{IDForm: t.IDForm, IDTable: t.ID})
	for col, value := range t.ColumnsName {
		d := widget.NewButtonWithIcon(value, nil, nil)
		d.OnTapped = func() {
			c := app_values[t.IDForm].Table[t.ID].wb[d]
			sort.Slice(t.Data, func(i, j int) bool { return t.Data[i][c] < t.Data[j][c] })
			t.Table.Refresh()
		}
		t.Header.Add(d)
		app_values[t.IDForm].Table[t.ID].wb[d] = col
	}
}

func (t *TableOtoko) LoadTable(mes []byte) {
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
			con.Add(widget.NewLabel(""))

			check := widget.NewCheck("", nil)
			check.OnChanged = func(b bool) {
				i := app_values[t.IDForm].Table[t.ID].wc[check]
				if check.Checked {
					t.Data[i.Row][i.Col] = "1"
				} else {
					t.Data[i.Row][i.Col] = "0"
				}
				newTableCellID := widget.TableCellID{Col: i.Col, Row: i.Row + 1}
				t.Table.ScrollTo(newTableCellID)
				println(i.Row)
				for key, value := range app_values[t.IDForm].Table[t.ID].wc {
					if value == newTableCellID {
						app_values[t.IDForm].W.Canvas().Focus(key)
						break
					}
				}
			}
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
			var label *widget.Label
			var ic *widget.Check
			var entry *enterEntry
			box := o.(*fyne.Container)
			rect := box.Objects[0].(*canvas.Rectangle)

			if i.Row%2 == 0 {
				rect.FillColor = t.AlterRowColor
			} else {
				rect.FillColor = t.RowColor
			}
			if i.Row == 0 {
				rect.FillColor = t.HeaderColor
			}

			cont := box.Objects[1].(*fyne.Container)

			label = cont.Objects[0].(*widget.Label)

			ic = cont.Objects[1].(*widget.Check)

			entry = cont.Objects[2].(*enterEntry)
			label.Hidden = true
			ic.Hidden = true
			entry.Hidden = true
			switch t.ColumnsType[i.Col] {
			case "bool":
				app_values[t.IDForm].Table[t.ID].wc[ic] = i
				if t.Data[i.Row][i.Col] == "1" {
					ic.Checked = true
				} else {
					ic.Checked = false
				}
				ic.Refresh()
				ic.Hidden = false
			case "string":
				entry.SetText(t.Data[i.Row][i.Col])
				app_values[t.IDForm].Table[t.ID].we[entry] = i
				entry.Hidden = false
			default:
				label.SetText(t.Data[i.Row][i.Col])
				label.Hidden = false
			}
		})
	for ic, v := range t.ColumnsWidth {
		t.Table.SetColumnWidth(ic, v)
	}
	t.Table.OnSelected = func(id widget.TableCellID) {
		fmt.Printf("i.Col: %v\n", id.Col)
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
		},
		),
	)

	

	content := container.NewBorder(
		container.NewVBox(
			t.Tool,
			widget.NewSeparator(),
		),
		nil, nil, nil, t.Table,
	)

	return content
}

//////////////////////////////////////////////
type enterEntry struct {
	IDForm  string
	IDTable string
	widget.Entry
}

func (e *enterEntry) onEnter() {
	fmt.Println(e.Entry.Text)
	i := app_values[e.IDForm].Table[e.IDTable].we[e]
	app_values[e.IDForm].Table[e.IDTable].Data[i.Row][i.Col] = e.Entry.Text
}

func newEnterEntry() *enterEntry {
	entry := &enterEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *enterEntry) KeyDown(key *fyne.KeyEvent) {
	t := app_values[e.IDForm].Table[e.IDTable]
	switch key.Name {
	case fyne.KeyReturn:
		e.onEnter()
	case "KP_Enter":
		e.onEnter()
	case "Down":
		i := t.we[e]
		newTableCellID := widget.TableCellID{Col: i.Col, Row: i.Row + 1}
		t.Table.ScrollTo(newTableCellID)
		for key, value := range app_values[e.IDForm].Table[e.IDTable].we {
			if value == newTableCellID {
				app_values[e.IDForm].W.Canvas().Focus(key)
				break
			}
		}
	case "Up":
		i := t.we[e]
		newTableCellID := widget.TableCellID{Col: i.Col, Row: i.Row - 1}
		t.Table.ScrollTo(newTableCellID)
		for key, value := range t.we {
			if value == newTableCellID {
				app_values[e.IDForm].W.Canvas().Focus(key)
				break
			}
		}
	default:
		e.Entry.KeyDown(key)
		fmt.Printf("Key %v pressed\n", key.Name)
	}
}

func (e *enterEntry) KeyUp(key *fyne.KeyEvent) {
	fmt.Printf("Key %v released\n", key.Name)
}

/////////////////////////////
type ToolButton struct {
	IDForm  string
	IDTable string
}

func (d *ToolButton) MinSize(objects []fyne.CanvasObject) fyne.Size {
	TO := app_values[d.IDForm].Table[d.IDTable]
	w, h := float32(0), float32(0)
	for i, o := range objects {
		childSize := o.MinSize()
		o.Resize(fyne.NewSize(TO.ColumnsWidth[i], childSize.Height))
		w += TO.ColumnsWidth[i]
		h = childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (d *ToolButton) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	TO := app_values[d.IDForm].Table[d.IDTable]
	for i, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(TO.ColumnsWidth[i], 0))
	}
}
