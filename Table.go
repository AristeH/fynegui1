package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

type ColumnStyle struct {
	ID      string
	Name    string  //Заголовок столбца
	Format  string  //Форматированный вывод
	Width   float32 //Ширина столбца
	BGColor string  // Цвет фона
	Type    string
}

type TabStyle struct {
	ID            string
	Name          string
	BGColor       string // Цвет фона
	RowAlterColor string // Цвет строки четной
	HeaderColor   string // Цвет текста
	RowColor      string // Цвет строки нечетной
	Font          string // Шрифт

}
type TableOtoko struct {
	ID          string
	IDForm      string
	ColumnStyle []ColumnStyle
	TabStyle    TabStyle
	Data        [][]string
	Tool        *widget.Toolbar
	Table       *widget.Table
	Header      *fyne.Container
	Properties  *TableOtoko
	wl          map[*widget.Label]widget.TableCellID
	we          map[*enterEntry]widget.TableCellID
	wc          map[*enterCheck]widget.TableCellID
	wb          map[*widget.Button]int
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

// / поле ввода
type enterEntry struct {
	IDForm  string
	IDTable string
	ID      string
	col     int
	widget.Entry
}

func (e *enterEntry) Tapped(ev *fyne.PointEvent) {
	t := appValues[e.IDForm].Table[e.IDTable]
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
	t := appValues[e.IDForm].Table[e.IDTable]
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
	i := appValues[e.IDForm].Table[e.IDTable].we[e]
	appValues[e.IDForm].Table[e.IDTable].Data[i.Row][i.Col] = e.Entry.Text
}

func newEnterEntry() *enterEntry {
	entry := &enterEntry{}

	entry.ExtendBaseWidget(entry)
	return entry
}

func scrolltable(row int, col int, t *TableOtoko) {
	switch t.ColumnStyle[col].Type {
	case "bool":
		newTableCellID := widget.TableCellID{Col: col, Row: row}
		t.Table.ScrollTo(newTableCellID)
		for key, value := range t.we {
			if value == newTableCellID {
				appValues[t.IDForm].W.Canvas().Focus(key)
				break
			}
		}

	default:
		newTableCellID := widget.TableCellID{Col: col, Row: row}
		t.Table.ScrollTo(newTableCellID)
	}

}

func (e *enterEntry) KeyDown(key *fyne.KeyEvent) {
	t := appValues[e.IDForm].Table[e.IDTable]
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

		if len(t.ColumnStyle[e.col].Type) == e.col+1 {
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

// / чек бокс
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
	t := appValues[e.IDForm].Table[e.IDTable]
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

		if len(t.ColumnStyle[e.col].Type) == e.col+1 {
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
