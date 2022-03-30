package main

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

//UserForm структура описывающая форму
type UserForm struct {
	Name  string
	ID    string
	Title string
	Icon  string
	Size  string
}

//Button - описание кнопки
type Button struct {
	Name  string //имя функции на сервере
	Title string //текст кнопки
	Param string //возвращаемые параметры
	Fun   string //имя функции на клиенте
	Image string //рисунок для кнопки
}

type ButtonData struct {
	Parameters string
	Fun        string
	Widget     *widget.Button
}

//FormData - данные формы
type FormData struct {
	Entry  map[string]entryForm   // Entry  - список полей ввода формы
	Table  map[string]*TableOtoko // Table  - список таблиц формы
	Tree   map[string]*TreeOtoko  // Table  - список таблиц формы
	Button map[string]ButtonData  // Button - список кнопок формы
	W      fyne.Window
}

//FieldSection - описание поля со значением и ролью
type FieldSection struct {
	Name    string // Уникальное имя в местоположении
	Title   string // Текст метки
	Value   string // Значение поля
	Tip     string // Тип поля
	Buttons string // кнопки 0-чтение, 1 - просмотр, 2 - редактирование
}

type Form struct {
	Title   []byte
	Toolbar []byte
	Fields  []byte
	Table   []byte
}

// обновим форму
func update_data(fd *FormData, res Message) {
	var form_data Form
	json.Unmarshal([]byte(res.Parameters), &form_data)
	// обновим сведения о форме
	// for key, value := range form_data {
	// 	i, ok := app_values[key]
	// 	if ok {
	// 		fmt.Println("Key:", value, "Value:", i)
	// 	}
	// }
}

func create_form(res Message) {
	var form_data Form
	var Title UserForm
	var w fyne.Window
	var ID string
	var fd *FormData = &FormData{}
	var toolbar *fyne.Container
	var fields *fyne.Container
	var tab *fyne.Container
	var content1 *fyne.Container
	var FieldS []FieldSection

	json.Unmarshal([]byte(res.Parameters), &form_data)
	json.Unmarshal(form_data.Title, &Title)
	json.Unmarshal(form_data.Fields, &FieldS)
	// обновим сведения о форме
	ID = Title.ID
	app_values[ID] = fd
	w = myApp.NewWindow(Title.Title)
	fd.W = w
	words := strings.Split(Title.Size, ",")
	wf, err := strconv.ParseFloat(words[0], 32)
	if err != nil {
		wf = 300
	}
	hf, err := strconv.ParseFloat(words[0], 32)
	if err != nil {
		hf = 400
	}
	w.Resize(fyne.NewSize(float32(wf), float32(hf)))

	fd.Table = make(map[string]*TableOtoko)
	//var fa [][]entryForm
	// for _, value := range res.Child {
	// 	//fmt.Println("Key:", key, "Value:", value)
	// 	switch value.Name {
	// 	case "fields":
	// 		var bt []entryForm
	// 		json.Unmarshal([]byte(value.Body), &bt)
	// 		fa = append(fa, bt)
	// 	case "toolbar":
	// 		toolbar = ToolBarCreate(ID, value.Body)
	// 	case "table":
	// 		var bt = TableOtoko{}
	// 		json.Unmarshal([]byte(value.Body), &bt)
	// 		bt.wb = make(map[*widget.Button]int)
	// 		bt.wc = make(map[*widget.Check]widget.TableCellID)
	// 		bt.we = make(map[*enterEntry]widget.TableCellID)
	// 		fd.Table[bt.ID] = &bt
	// 		tab = bt.makeTable()
	// 		//ntent.Add(fd.Table[bt.ID].Table)
	// 	}
	// }
	fields = FieldsCreate(ID, FieldS)
	vb := container.NewVBox()

	if toolbar != nil {
		vb.Add(toolbar)
		vb.Add(widget.NewSeparator())
	}

	if fields != nil {
		vb.Add(fields)
		vb.Add(widget.NewSeparator())
	}

	content1 = container.NewBorder(
		vb,
		nil,
		nil,
		nil,
	)

	if tab != nil {
		content1.Add(tab)

	}

	w.SetContent(content1)

	w.Show()

}

func FieldsCreate(id string, f []FieldSection) *fyne.Container {
	// var vb []*fyne.Container
	// fEntry := make(map[string]entryForm)
	// v := container.New(layout.NewHBoxLayout())

	// for i := 0; i < len(f); i++ {
	// 	for k := len(vb); k < len(f[i])*2; k++ {
	// 		vb = append(vb, container.New(layout.NewVBoxLayout()))
	// 		vb = append(vb, container.New(layout.NewVBoxLayout()))
	// 	}
	// 	for j := 0; j < len(f[i]); j++ {
	// 		vb[j*2].Add(widget.NewLabel(f[i][j].Name))

	// 		entry := widget.NewEntry()
	// 		if f[i][j].Value == "" {
	// 			entry.PlaceHolder = f[i][j].Title
	// 		} else {
	// 			entry.Text = f[i][j].Value
	// 		}
	// 		entry.Wrapping = fyne.TextWrapOff
	// 		vb[j*2+1].Add(entry)
	// 		fEntry[f[i][j].Name] = entryForm{Value: f[i][j].Value, Widget: entry}
	// 	}
	// }

	// for i := 0; i < len(vb); i++ {

	// 	v.Add(vb[i])
	// }
	// beans := app_values[id]
	// beans.Entry = fEntry
	// app_values[id] = beans
	// for j := 0; j < len(f[0]); j++ {
	// 	h := container.New(layout.NewVBoxLayout())

	// 	for i := 0; i < len(f); i++ {
	// 		h.Add(widget.NewLabel(f[i][j].Name))
	// 		// entry := widget.NewEntry()
	// 		// entry.PlaceHolder = f[i][j].Title
	// 		// h.Add(entry)
	// 	}

	// 	v.Add(h)
	// }

	return nil
}

func findButton(d *widget.Button) (string, *FormData) {
	for _, f := range app_values {
		if f.Button[d.Text].Widget == d {
			return f.Button[d.Text].Parameters, f
		}
	}
	return "", &FormData{}
}

// func findEntry(f *FormData, name string)string{

// 	return f.Entry[name].Value

// }

func ToolBarCreate(id string, but []byte) *fyne.Container {
	var bt []Button
	fd := app_values[id]
	fd.Button = make(map[string]ButtonData)
	toolbar := container.New(layout.NewHBoxLayout())
	json.Unmarshal([]byte(but), &bt)

	for _, value := range bt {
		//	fmt.Println("Key:", key, "Value:", value)
		d := widget.NewButtonWithIcon(value.Name, GetIcon(value.Image), nil)
		d.OnTapped = func() {
			param, f := findButton(d)
			mp := strings.Split(param, ",")
			p := ""
			for _, r := range mp {
				if _, ok := f.Entry[r]; ok {
					p = p + r + ":" + f.Entry[r].Widget.Text + ","
				} else {
					p = p + r + ";"
				}

			}
			mes := map[string]string{
				"Action":     d.Text,
				"Parameters": p,
			}
			send(mes)

		}
		fd.Button[value.Name] = ButtonData{Fun: value.Name, Parameters: value.Param, Widget: d}
		toolbar.Add(d)
	}
	beans := app_values[id]
	beans.Button = fd.Button
	app_values[id] = beans

	return toolbar
}

func GetIcon(n string) fyne.Resource {
	switch n {
	case "DocumentCreateIcon":
		return theme.DocumentCreateIcon()
	case "AccountIcon":
		return theme.AccountIcon()
	}
	return nil
}

func send(mes map[string]string) {
	jsonMessage, _ := json.Marshal(&mes)
	Cl.Reci <- jsonMessage
}
