package main

import (
	"bytes"
	"encoding/gob"
	"fyne.io/fyne/v2"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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

//GetDataContainer - получим описание формы 2
func GetDataContainer(idform string, idcontainer string) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	d := GetData{ID: idform, Container: idcontainer}
	mes := MessageGob{
		Action: "GetDataContainer",
		Data:   d,
	}
	enc.Encode(mes)
	k := buff.Bytes()
	println(k)
	Cl.Reci <- k
}

func AccordionTable(c *MessageGob) {
	tbl := c.Data.Data
	contCatalog := container.NewVBox()
	contDocument := container.NewVBox()
	acc := widget.NewAccordion()
	for _, b := range tbl {
		d := widget.NewButton(b[Synonym], nil)
		d.OnTapped = func() {
			_, param := findButton(d)
			mp := strings.Split(param.Parameters, ";")
			w := InitForm(param.Parameters, mp[0])
			w.Show()
		}

		appValues[c.Data.ID].Button[b[ID]] = ButtonData{Fun: b[Name] + "GenForm", Parameters: b[ID], Widget: d}
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
	fd := appValues[c.Data.ID].form
	parent := fd[c.Data.Container][ParentID]
	appValues[c.Data.ID].Container[parent] = d
	SetContent(c.Data.ID)
}

func ToolBarParent(f, c string) fyne.CanvasObject {
	form := appValues[f]
	fd := appValues[f].form

	mp := strings.Split(fd[c][ChildrensID], ";")
	top := container.New(layout.NewGridLayoutWithRows(2))
	for _, rec := range mp {
		if rec != "" {
			if form.Container[rec] != nil {
				top.Add(form.Container[rec])
			}
		}
	}
	return top
}

func SetContent(idform string) {
	var root string

	var top fyne.CanvasObject
	var left fyne.CanvasObject
	var right fyne.CanvasObject
	var bottom fyne.CanvasObject
	var middle fyne.CanvasObject

	win := appValues[idform].W
	f := appValues[idform].form

	for i := range f {
		if f[i][ParentID] == "0" {
			root = f[i][ID]
			break
		}
	}

	mproot := strings.Split(f[root][ChildrensID], ";")

	for _, recroot := range mproot {
		if recroot != "" {
			mpborder := strings.Split(f[recroot][ChildrensID], ";")
			for _, recb := range mpborder {
				if recb != "" {
					if f[recb][TypeContainer] == "top" {
						top = appValues[idform].Container[recb]
					}
					if f[recb][TypeContainer] == "left" {
						left = appValues[idform].Container[recb]
					}
					if f[recb][TypeContainer] == "middle" {
						middle = appValues[idform].Container[recb]
					}
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

	// 	content := container.New(lb,top)
	win.SetContent(content)
	win.Show()
}

func createParent(f, c string) {
	form := appValues[f].form
	if form[form[c][ParentID]][OrderOutput] == "0" {
		switch form[c][TypeContainer] {
		case "Toolbar":
			appValues[f].Container[form[c][ParentID]] = ToolBarParent(f, form[c][ParentID])
		}
	} else {
		if form[form[c][ParentID]][TypeContainer] == "TableList" {
			var top fyne.CanvasObject
			var left fyne.CanvasObject
			var right fyne.CanvasObject
			var bottom fyne.CanvasObject
			var middle fyne.CanvasObject
			switch form[c][TypeContainer] {
			case "Toolbar":
				top = ToolBarParent(f, form[c][ParentID])
			case "Table":
				middle = appValues[f].Container[c]
			}
			mpborder := strings.Split(form[form[c][ParentID]][ChildrensID], ";")
			for _, recb := range mpborder {
				if recb != "" {
					if form[recb][TypeContainer] == "Toolbar" {
						top = appValues[f].Container[recb]
					}
					if form[recb][TypeContainer] == "left" {
						left = appValues[f].Container[recb]
					}
					if form[recb][TypeContainer] == "Table" {
						middle = appValues[f].Container[recb]
					}
				}

			}
			content := container.New(layout.NewBorderLayout(top, bottom, left, right))
			if top != nil {
				content.Add(top)
			}
			if middle != nil {
				content.Add(middle)
			}
			appValues[f].Container[form[c][ParentID]] = content
			p := form[c][ParentID]
			appValues[f].Container[form[p][ParentID]] = content
		}

	}

}

func ToolBar(c *MessageGob) {
	app := c.Data.Data

	tb := ToolBarCreate(c.Data.ID, c.Data.Container, app, color.Gray{Y: 230})
	appValues[c.Data.ID].Container[c.Data.Container] = tb
	createParent(c.Data.ID, c.Data.Container)
	SetContent(c.Data.ID)
}

func InitFormLocal(c *MessageGob) {
	app := c.Data.Data
	fd := appValues[c.Data.ID]
	for i := range app {
		fd.form[app[i][ID]] = app[i]
		if app[i][ParentID] != "0" {
			fd.Container[strconv.Itoa(i)] = nil
		}
	}
}

func InitFormView(c *MessageGob) {
	app := c.Data.Data
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
}

// InitForm - инициализация формы 1
func InitForm(form, parameters string) fyne.Window {
	var br [][]string
	logger.Tracef("Инициализация формы:" + form + " параметры:" + parameters)
	myWindow := myApp.NewWindow("Телефоны")
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
	output[ID] = form         // гуид подсистмы/кнопки
	output[Name] = parameters //параметрыфункции

	br = append(br, output)
	d := GetData{ID: form, Data: br}
	SendMessage("GetDataContainer", d)
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
