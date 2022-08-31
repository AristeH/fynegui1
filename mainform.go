package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
	"strings"
)

// константы метаданные для вывода интерфейса пользователя, номера столбца для вывода
const (
	ID            int = 0  //гуид метаданного(подсистема, таблица, реквизит, форма)
	Name          int = 1  //имя англ
	Synonym       int = 2  //синоним
	ParentID      int = 3  //гуид родителя или владельца метаданного
	OrderOutput   int = 4  //порядок вывода элементов интерфейса
	Icon          int = 5  //область вывода{top,bottom,left,right,middle,""}
	TypeContainer int = 6  //имя контейнера, виджета
	NameContainer int = 7  //имя заголовка контейнера
	Style         int = 8  //ширина
	ChildrensID   int = 9  //иерархия дети( подсистемы, табличные части, реквизиты)
	Parameters    int = 10 // тип метаданного
	Fun           int = 11 // тип метаданного
)

func UpdateForm(mes *MessageGob) {
	f := mes.Data.ID
	c := mes.Data.Container
	if c == "form" {
		InitFormLocal(mes)

		return
	}
	form := appValues[f].form // map[string][]string
	for i := range form {
		if c == form[i][ID] {
			if fnc, bExist := mfu[form[i][TypeContainer]]; bExist {
				fnc(mes)
			} else {
				logger.Errorf("Не удалось найти функцию: " + form[i][TypeContainer])
			}
			return
		}
	}

}

func nextContainer(mes *MessageGob) {

	fd := appValues[mes.Data.ID]
	fun := true

	for i := range fd.form {
		if fun && fd.form[i][ID] > mes.Data.Container && fd.form[i][Parameters] == "true" {
			var br [][]string
			output := make([]string, 12)

			output[ID] = fd.form[i][ID] // гуид подсистмы/кнопки
			//
			output[Fun] = fd.form[i][Fun] // функция для кнопки

			br = append(br, output)
			fun = false
			d := GetData{ID: mes.Data.ID, Data: br, Container: output[ID]}
			UpdateContainer(d)
			break
		}
	}
	if fun {
		SetContent(mes.Data.ID)
	}

}

// InitFormLocal - инициализация структуры формы.
func InitFormLocal(mes *MessageGob) {
	app := mes.Data.Data
	fd := appValues[mes.Data.ID]
	fun := true
	for i := range app {
		fd.form[app[i][ID]] = app[i]
		if fun && app[i][Parameters] == "true" {
			var br [][]string
			output := make([]string, 12)

			output[ID] = app[i][ID] // гуид подсистмы/кнопки

			output[Fun] = app[i][Fun] // функция для кнопки
			fun = false
			br = append(br, output)

			d := GetData{ID: mes.Data.ID, Data: br, Container: output[ID]}
			UpdateContainer(d)
		}
		if app[i][ParentID] != "0" {
			fd.Container[strconv.Itoa(i)] = nil
		}
	}

}

// initform - заголовок, стиль формы
func initform(mes *MessageGob) {
	//f := mes.Data.ID
	//c := mes.Data.Container
	app := mes.Data.Data

	//attr := getid(c.Data.ID)
	fd := appValues[app[0][ID]]
	fd.W.SetTitle(app[0][Synonym])
	w := strings.Split(app[0][Style], ";")
	if len(w) == 2 {
		wf, _ := strconv.Atoi(w[0])
		hf, _ := strconv.Atoi(w[1])
		fd.W.Resize(fyne.NewSize(float32(wf), float32(hf)))
	} else {
		fd.W.Resize(fyne.NewSize(1000, 100))
	}

	fd.W.SetCloseIntercept(func() {
		fd.W.Hide()
	})

}

func ToolBar(mes *MessageGob) {
	f := mes.Data.ID
	c := mes.Data.Container
	app := mes.Data.Data
	tb := ToolBarCreate(f, c, app, color.Gray{Y: 230})
	appValues[f].Container[c] = tb
	createParent(f, appValues[f].form[c][ParentID])
	nextContainer(mes)
}

func Accordion(mes *MessageGob) {
	f := mes.Data.ID
	c := mes.Data.Container
	tbl := mes.Data.Data
	contCatalog := container.NewVBox()
	contDocument := container.NewVBox()
	acc := widget.NewAccordion()
	for _, b := range tbl {
		d := widget.NewButton(b[Synonym], nil)
		d.OnTapped = func() {
			_, param := findButton(d)
			//mp := strings.Split(param.Parameters, ";")
			w := InitForm(param.Parameters, "")
			w.Show()
			w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {

				i := activeContainer.Selected
				switch k.Name {
				case "Down":
					if len(activeContainer.Data)-1 > activeContainer.Selected.Row {
						activeContainer.Selected = widget.TableCellID{Col: i.Col, Row: i.Row + 1}
					}
				case "Up":
					if i.Row > 1 {
						activeContainer.Selected = widget.TableCellID{Col: i.Col, Row: i.Row - 1}
					}
				case "Left":
					if i.Col >= 1 {
						activeContainer.Selected = widget.TableCellID{Col: i.Col - 1, Row: i.Row}
					}
				case "Right":
					if len(activeContainer.Data[0])-1 > i.Col {
						activeContainer.Selected = widget.TableCellID{Col: i.Col + 1, Row: i.Row}
					}
				case "KP_Enter", "Return":
					Entry := widget.NewEntry()
					Entry.Validator = getValidator(activeContainer.ColumnStyle[i.Row].Type)
					items := []*widget.FormItem{
						widget.NewFormItem("Username", Entry),
					}
					dialog.ShowForm("введите", "", "cancel", items, func(b bool) {
						if !b {
							return
						}
						fmt.Println("KP_Enter", Entry.Text)
						activeContainer.Data[i.Row][i.Col] = Entry.Text
					}, w)

				}
				activeContainer.Table.ScrollTo(activeContainer.Selected)
				activeContainer.Table.Refresh()
			})
		}

		appValues[f].Button[b[ID]] = ButtonData{Fun: b[Name] + "GenForm", Parameters: b[ID], Widget: d}
		switch b[TypeContainer] {
		case "Справочник":
			contCatalog.Add(d)
		case "Документ":
			contDocument.Add(d)
		}
	}
	acc.Append(&widget.AccordionItem{Title: "Документы", Detail: contDocument})
	acc.Append(&widget.AccordionItem{Title: "Справочники", Detail: contCatalog})
	d := container.NewVBox(container.NewVScroll(acc))
	d.Layout = layout.NewMaxLayout()
	appValues[f].Container[c] = d
	createParent(f, appValues[f].form[c][ParentID])
	nextContainer(mes)

}

func top(f, c string) {
	form := appValues[f]
	fd := appValues[f].form

	mp := strings.Split(fd[c][ChildrensID], ";")
	top := container.New(layout.NewGridLayoutWithRows(len(mp) - 1))
	for _, rec := range mp {
		if rec != "" {
			if form.Container[rec] != nil {
				top.Add(form.Container[rec])
			}
		}
	}

	appValues[f].Container[c] = top
	createParent(f, fd[c][ParentID])
}

func TableList(f, c string) {
	var t fyne.CanvasObject
	var m fyne.CanvasObject
	form := appValues[f]
	fd := appValues[f].form

	mp := strings.Split(fd[c][ChildrensID], ";")

	if val, ok := form.Container[mp[0]]; ok {
		t = val
	}
	if val, ok := form.Container[mp[1]]; ok {
		m = val
		content := container.NewBorder(
			t,
			nil,
			nil,
			nil,
			m,
		)
		appValues[f].Container[c] = content
	}

	createParent(f, fd[c][ParentID])
}

func border(f, c string) {
	var top fyne.CanvasObject
	var left fyne.CanvasObject
	var right fyne.CanvasObject
	var bottom fyne.CanvasObject
	var middle fyne.CanvasObject

	form := appValues[f]
	fd := appValues[f].form

	mp := strings.Split(fd[c][ChildrensID], ";")

	for _, rec := range mp {
		if rec != "" {
			if form.Container[rec] != nil {
				switch fd[rec][TypeContainer] {
				case "top":
					top = form.Container[rec]

				case "bottom":
					bottom = form.Container[rec]

				case "left":
					left = form.Container[rec]

				case "right":
					right = form.Container[rec]

				case "middle":
					middle = form.Container[rec]

				}
			}
		}
	}
	content := container.New(layout.NewBorderLayout(top, bottom, left, right))
	if left != nil {
		content.Add(left)
	}
	if top != nil {
		content.Add(top)
	}
	if middle != nil {
		content.Add(middle)
	}
	form.Container[c] = content

	//appValues[f].Container[c] = top
	//createParent(f, fd[c][ParentID])
}

func SetContent(f string) {
	var content fyne.CanvasObject
	form := appValues[f].form // map[string][]string
	for i := range form {
		if "init" == form[i][TypeContainer] {
			root := strings.Split(form[i][ChildrensID], ";")
			for _, rec := range root {
				if rec != "" {
					content = appValues[f].Container[rec]
				}
			}
			break
		}
	}
	appValues[f].W.SetContent(content)
	appValues[f].W.Show()
}

func createParent(f, c string) {
	form := appValues[f].form

	switch form[c][TypeContainer] {
	case "top":
		top(f, c)

	case "Border":
		border(f, c)
	case "TableList":
		TableList(f, c)
	case "left":
		top(f, c)
	case "middle":
		top(f, c)
	}
}

// InitForm - инициализация формы 1
func InitForm(form, parameters string) fyne.Window {
	var br [][]string
	logger.Tracef("Инициализация формы f:" + form + " c:" + parameters)
	if val, ok := appValues[form]; ok {
		return val.W
	}
	myWindow := myApp.NewWindow("")
	myWindow.Resize(fyne.NewSize(1200, 400))
	appValues[form] = &FormData{
		ID:        form,
		W:         myWindow,
		Button:    make(map[string]ButtonData),
		Container: make(map[string]fyne.CanvasObject),
		Entry:     make(map[string]entryForm),
		Table:     make(map[string]*TableOtoko),
		Tree:      make(map[string]*TreeOtoko),
		form:      make(map[string][]string),
	}

	output := make([]string, 12)
	output[ID] = form               // гуид подсистмы/кнопки
	output[Parameters] = parameters //параметрыфункции
	appValues[form].W = myWindow
	br = append(br, output)
	d := GetData{ID: form, Container: "", Data: br}
	logger.Tracef("Получим данные формы с параметрами:" + form + " параметры:" + parameters)
	UpdateContainer(d)
	//UpdateFormContent(GetData{ID: form})
	myWindow.SetCloseIntercept(func() {
		myWindow.Hide()
	})
	return myWindow
}

// SendMessage -  отправить сообщение серверу
func SendMessage(Action string, d GetData) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	mes := MessageGob{
		Action: Action,
		Data:   d,
	}
	enc.Encode(mes)
	k := buff.Bytes()
	logger.Infof("ОТПР- Форма:" + d.ID + "Контейнер:" + d.Container + "Функция:" + Action + " параметры:" + d.Data[0][Parameters])
	Cl.Reci <- k
}

func UpdateContainer(param GetData) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	enc.Encode(MessageGob{
		Action: "GetDataContainer", //update container
		Data:   param,              // данные контейнера, формы
	})
	logger.Tracef("Отправка сообщения форма:" + param.ID + " контейнер:" + param.Container)
	Cl.Reci <- buff.Bytes()
}

func UpdateFormContent(param GetData) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	enc.Encode(MessageGob{
		Action: "UpdateFormContent", //update container
		Data:   param,               // данные контейнера, формы
	})
	Cl.Reci <- buff.Bytes()
}

func FieldsCreate(mes *MessageGob) {

	var vb []*fyne.Container
	f := mes.Data.ID
	c := mes.Data.Container
	fd := mes.Data.Data
	fEntry := make(map[string]entryForm)
	v := container.New(layout.NewHBoxLayout())

	for i := 0; i < len(fd); i++ {
		for k := len(vb); k < len(fd[i])*2; k++ {
			vb = append(vb, container.New(layout.NewVBoxLayout()))
			vb = append(vb, container.New(layout.NewVBoxLayout()))
		}
		for j := 0; j < len(fd[i]); j++ {
			vb[j*2].Add(widget.NewLabel(fd[i][j])) //.Name
			entry := widget.NewEntry()
			if fd[i][j] == "" {
				entry.PlaceHolder = fd[i][j] //.Title
			} else {
				entry.Text = fd[i][j] //.Value
			}
			entry.Wrapping = fyne.TextWrapOff
			vb[j*2+1].Add(entry)
			fEntry[fd[i][j]] = entryForm{Value: fd[i][j], Widget: entry} //.Value
		}
	}

	for i := 0; i < len(vb); i++ {
		v.Add(vb[i])
	}
	beans := appValues[f]
	beans.Entry = fEntry
	appValues[f] = beans
	for j := 0; j < len(fd[0]); j++ {
		h := container.New(layout.NewVBoxLayout())
		for i := 0; i < len(fd); i++ {
			h.Add(widget.NewLabel(fd[i][j]))
			// entry := widget.NewEntry()
			// entry.PlaceHolder = f[i][j].Title
			// h.Add(entry)
		}
		v.Add(h)
	}
	createParent(f, appValues[f].form[c][ParentID])
	nextContainer(mes)
}
