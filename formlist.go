package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func NewTableList1(mess *GetData) *widget.Table {
	data := mess.Data
	t := &TableOtoko{}
	TextStyle := fyne.TextStyle{
		Bold:      false,
		Italic:    false,
		Monospace: false,
		TabWidth:  0,
	}

	si := fyne.MeasureText("ф", 12, TextStyle).Width

	for i := 0; i < len(mess.DataDescription[2]); i++ {
		cs := ColumnStyle{}
		cs.Name = mess.DataDescription[2][i]
		cs.Type = mess.DataDescription[1][i]

		b := mess.DataDescription[0][i] == "id" || strings.HasPrefix(mess.DataDescription[0][i], "id_")
		if b {
			cs.Width = 0
		} else {
			l := 0

			for j := 0; j < len(mess.Data); j++ {
				if l < len([]rune(mess.Data[j][i])) {
					l = len([]rune(mess.Data[j][i]))
				}
			}
			cs.Width = float32(l) * si
		}

		t.ColumnStyle = append(t.ColumnStyle, cs)
	}

	t.TabStyle.RowAlterColor = "RowAlterColor"
	t.TabStyle.HeaderColor = "HeaderColor"
	t.TabStyle.RowColor = "RowColor"
	t.Data = data
	t.ID = mess.Container
	t.IDForm = mess.Form
	//t.wol = make(map[*oLabel]widget.TableCellID)
	t.MakeTableLabel()
	activeContainer = t
	return t.Table
}

func Table(mess *GetData) {
	f := mess.Form
	c := mess.Container
	t := NewTableList1(mess)

	appValues[f].Container[c] = t

	createParent(f, appValues[f].form[c][ParentID])
	nextContainer(mess)
	t.Refresh()
}
