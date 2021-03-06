package main

import (
	"bytes"
	"encoding/gob"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
		}
	}
	if fun {
		SetContent(mes.Data.ID)
	}

}

//InitFormLocal - инициализация структуры формы.
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
	SetContent(f)
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
	SetContent(f)

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

	case "left":
		top(f, c)
	case "middle":
		top(f, c)
	}
}

// InitForm - инициализация формы 1
func InitForm(form, parameters string) fyne.Window {
	var br [][]string
	logger.Tracef("Инициализация формы:" + form + " параметры:" + parameters)
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

	br = append(br, output)
	d := GetData{ID: form, Container: "", Data: br}
	logger.Tracef("Получим данные формы с параметрами:" + form + " параметры:" + parameters)
	UpdateContainer(d)
	//UpdateFormContent(GetData{ID: form})
	return myWindow
}

//SendMessage -  отправить сообщение серверу
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
