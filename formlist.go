package main

import (
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func NewTableList1(mess *MessageGob) *widget.Table {
	data := mess.Data.Data
	t := &TableOtoko{}

	for i := 0; i < len(mess.Data.DataDescription[2]); i++ {
		cs := ColumnStyle{}
		cs.Name = mess.Data.DataDescription[2][i]
		cs.Type = mess.Data.DataDescription[1][i]
		ch, _ := strconv.Atoi(mess.Data.DataDescription[3][i])
		cs.Width = float32(ch) * 32
		t.ColumnStyle = append(t.ColumnStyle, cs)
	}

	t.TabStyle.RowAlterColor = "lightslategrey"
	t.TabStyle.HeaderColor = "darkslategrey"
	t.TabStyle.RowColor = "lightgrey"
	t.Data = data
	t.ID = mess.Data.Container
	t.IDForm = mess.Data.ID
	t.wol = make(map[*oLabel]widget.TableCellID)
	t.MakeTableLabel()
	activeContainer = t
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
