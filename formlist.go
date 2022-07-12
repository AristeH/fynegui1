package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func NewTableList1(IDForm, IDTable string, data [][]string) *widget.Table {
	t := &TableOtoko{}
	t.ColumnsName = data[2]
	t.ColumnsType = data[1]
	t.ColumnsWidth = []float32{40}
	t.AlterRowColor = color.Gray{Y: 250}
	t.HeaderColor = color.Gray{Y: 80}
	t.RowColor = color.Gray{Y: 200}
	t.Data = data
	t.Edit = true
	t.ID = IDTable
	t.IDForm = IDForm
	//TO.wb = make(map[*widget.Button]int)
	t.wc = make(map[widget.TableCellID]*enterCheck)
	t.we = make(map[widget.TableCellID]*enterEntry)

	t.Table = widget.NewTable(
		func() (int, int) {
			return len(t.Data), len(t.Data[0])
		},
		func() fyne.CanvasObject {

			return container.New(layout.NewMaxLayout(),
				canvas.NewRectangle(color.Gray{Y: 250}),
				widget.NewLabel("wide content"),
			)

		},
		func(i widget.TableCellID, o fyne.CanvasObject) {

			box := o.(*fyne.Container)
			rect := box.Objects[0].(*canvas.Rectangle)
			if i.Row == 0 {
				rect.FillColor = t.HeaderColor
			} else if i.Row%2 == 0 {
				rect.FillColor = t.AlterRowColor
			} else {
				rect.FillColor = t.RowColor
			}
			cont := box.Objects[1].(*widget.Label)
			if i.Row == 0 {
				cont.SetText(t.ColumnsName[i.Col])
			} else {
				cont.SetText(data[i.Row][i.Col])
			}

		})

	//table
	t.ColumnsWidth = make([]float32, len(t.ColumnsName))
	for j := 0; j < len(t.ColumnsName); j++ {
		t.ColumnsWidth[j] = float32(len(data[2][j])) * 9
	}

	// 		fd.Table[app.Table].ColumnsType = ColumnsType
	// 		fd.Table[app.Table].ColumnsWidth = ColumnsWidth
	// 		fd.Table[app.Table].ColumnsName = ColumnsName

	for ic, v := range t.ColumnsWidth {
		t.Table.SetColumnWidth(ic, v)
	}
	t.ColumnsWidth[0] = 0
	appValues[IDForm].Table[IDTable] = t
	return t.Table

}

func Table(mess *MessageGob) {
	f := mess.Data.ID
	c := mess.Data.Container
	list := NewTableList1(f, f, mess.Data.Data)
	appValues[f].Container[c] = list

	createParent(f, appValues[f].form[c][ParentID])
	SetContent(f)
}
