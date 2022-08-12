package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

func NewTableList1(mess *MessageGob) *widget.Table {
	data := mess.Data.Data
	t := &TableOtoko{}
	for i := 0; i < len(mess.Data.DataDescription[2]); i++ {
		t.ColumnStyle[i].Name = mess.Data.DataDescription[2][i]
		t.ColumnStyle[i].Type = mess.Data.DataDescription[1][i]
	}

	t.TabStyle.RowAlterColor = "lightslategrey"
	t.TabStyle.HeaderColor = "darkslategrey"
	t.TabStyle.RowColor = "lightgrey"
	t.Data = data
	t.ID = mess.Data.Container
	t.IDForm = mess.Data.ID

	//TO.wb = make(map[*widget.Button]int)
	//t.wc = make(map[widget.TableCellID]*enterCheck)
	//	t.we = make(map[widget.TableCellID]*enterEntry)

	t.Table = widget.NewTable(
		func() (int, int) {
			return len(t.Data), len(t.Data[0])
		},
		func() fyne.CanvasObject {

			return container.New(layout.NewMaxLayout(),
				canvas.NewRectangle(color.Gray{Y: 250}),
				canvas.NewText("Hello world", color.Opaque), //widget.NewLabel("wide content")
			)

		},
		func(i widget.TableCellID, o fyne.CanvasObject) {

			box := o.(*fyne.Container)

			rect := box.Objects[0].(*canvas.Rectangle)
			if i.Row == 0 {
				box.Objects[1].(*canvas.Text).Text = t.ColumnStyle[i.Col].Name
				rect.FillColor = MapColor[t.TabStyle.HeaderColor]
			} else if i.Row%2 == 0 {
				box.Objects[1].(*canvas.Text).Text = data[i.Row-1][i.Col]
				box.Objects[1].(*canvas.Text).Color = &color.NRGBA{R: 128, G: 0, B: 128, A: 255}
				//box.Objects[1].(*widget.Label).SetText(data[i.Row-1][i.Col])
				rect.FillColor = MapColor[t.TabStyle.RowAlterColor]
			} else {
				box.Objects[1].(*canvas.Text).Text = data[i.Row-1][i.Col]
				rect.FillColor = MapColor[t.TabStyle.RowColor]
			}
		})
	for i := 0; i < len(mess.Data.DataDescription[1]); i++ {
		s, _ := strconv.ParseFloat(mess.Data.DataDescription[3][i], 32)
		t.Table.SetColumnWidth(i, float32(s)*32)
	}

	return t.Table
}

func Table(mess *MessageGob) {
	f := mess.Data.ID
	c := mess.Data.Container
	t := NewTableList1(mess)

	appValues[f].Container[c] = t

	createParent(f, appValues[f].form[c][ParentID])
	nextContainer(mess)
}
