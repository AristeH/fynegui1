package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

var (
	floatValidator = validation.NewRegexp("[+-]?([0-9]*[.])?[0-9]+", "Не правильное число")
	emailValidator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "Не правильный email")
	intValidator   = validation.NewRegexp(`[+-]?[0-9]*$`, "Не целое число")
	dateValidator  = validation.NewRegexp(`^\d{4}-\d{2}-\d{2}$`, "Не правильная дата")
	//emptyValidator = validation.NewRegexp(`^.+$`, "Поле не может быть пустым")
	//yearValidator  = validation.NewRegexp(`^[0-9]{4}$`, "Год содержит только 4 цифры.")
)

type ColumnStyle struct {
	ID      string
	Name    string  //Заголовок столбца
	Format  string  //Форматированный вывод
	Width   float32 //Ширина столбца
	BGColor string  // Цвет фона
	Color   string  // Цвет текста
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
	W           fyne.Window
	Selected    widget.TableCellID
	//	wol         map[*oLabel]widget.TableCellID
	wb map[*widget.Button]int
}

func (t *TableOtoko) MakeTableLabel() {
	t.Table = widget.NewTable(
		func() (int, int) {
			rows := len(t.Data)
			columns := len(t.Data[0])
			return rows, columns
		},
		func() fyne.CanvasObject {
			entry := newOLabel()
			entry.IDForm = t.IDForm
			entry.IDTable = t.ID
			entry.parent = t

			return container.New(layout.NewMaxLayout(),
				canvas.NewRectangle(color.Gray{Y: 250}),
				entry,
			)

		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			var entry *oLabel

			box := o.(*fyne.Container)
			rect := box.Objects[0].(*canvas.Rectangle)
			entry = box.Objects[1].(*oLabel)
			entry.Ind = &i
			entry.SetText(t.Data[i.Row][i.Col])
			if t.ColumnStyle[i.Col].Width == 0 {
				entry.Hidden = true
			} else {
				entry.Hidden = false
			}

			if t.ColumnStyle[i.Col].Type == "float" {
				entry.Label.Alignment = fyne.TextAlignTrailing
			} else {
				entry.Label.Alignment = fyne.TextAlignLeading
			}
			entry.TextStyle = fyne.TextStyle{
				Bold: false,
			}
			if i.Row == 0 {
				rect.FillColor = MapColor[t.TabStyle.HeaderColor]
				entry.Alignment = fyne.TextAlignCenter
				entry.TextStyle = fyne.TextStyle{
					Bold: true,
				}
			} else if i.Row%2 == 0 {
				rect.FillColor = MapColor[t.TabStyle.RowAlterColor]

			} else {
				rect.FillColor = MapColor[t.TabStyle.RowColor]
			}
			if val, ok := MapColor[t.ColumnStyle[i.Col].BGColor]; ok {
				rect.FillColor = mix(val, rect.FillColor)
			}
			if i == t.Selected {
				rect.FillColor = MapColor["Selected"]
			}
		})
	for ic, v := range t.ColumnStyle {
		t.Table.SetColumnWidth(ic, v.Width)
	}
	t.Table.OnSelected = func(id widget.TableCellID) {
		t.Selected = id
		fmt.Printf("i.Col: %v\n", id.Col)
	}
	//	t.Table.Refresh()
	//t.Tool = widget.NewToolbar(
	//	widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
	//		log.Println("New document")
	//	}),
	//	widget.NewToolbarSeparator(),
	//	widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
	//	widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	//	widget.NewToolbarSpacer(),
	//	widget.NewToolbarAction(theme.SettingsIcon(), func() {}))
}

func getValidator(t string) fyne.StringValidator {
	switch t {
	case "String":
		return dateValidator
	case "email":
		return emailValidator
	case "Time":
		return floatValidator
	case "DOUBLE":
		return intValidator
	case "Перечисление":
		return intValidator
	case "bool":
		return intValidator
	default:
		return nil
	}
}

func getValue(s string, t string) string {
	switch t {
	case "float":
		if s, err := strconv.ParseFloat(s, 32); err == nil {
			return fmt.Sprintf("%15f", s)
		}
	case "int":
		if s, err := strconv.Atoi(s); err == nil {
			return fmt.Sprintf("%3d", s)
		}
	default:
		return s
	}
	return ""
}
