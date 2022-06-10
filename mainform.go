package main

import (
	"bytes"
	"encoding/gob"
	"image/color"
	"strconv"
	"strings"

	//"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
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
	MestoOutput   int = 5  //область вывода{top,bottom,left,right,middle,""}
	TypeContainer int = 6  //имя контейнера, виджета
	NameContainer int = 7  //имя заголовка контейнера
	width         int = 8  //ширина
	ChildrensID   int = 9  //иерархия дети( подсистемы, табличные части, реквизиты)
	TypeMetaData  int = 10 // тип метаданного
	Fun           int = 11 // тип метаданного
)

//GetDescription - получим описание формы 2
func GetDataContainer(idform string, idcontainer string) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	d := GetData{Table: "md_forms", ID: idform, Type: idcontainer}
	mes := MessageGob{
		Action: "GetDataContainer",
		Data:   d,
	}
	enc.Encode(mes)
	k := buff.Bytes()
	println(k)
	Cl.Reci <- k
}


func GetListTable(fd *FormData, p *ButtonData) {
	for i := range fd.form {
		if fd.form[i][Fun] == "ListTable" {
			GetDataContainer("idform:"+fd.ID+";idcontainer:"+fd.form[i][ID], p.Parameters)
		}
	}
}

func GetToolBarPodsystem2(fd *FormData, p *ButtonData) {
	for i := range fd.form {
		if fd.form[i][Fun] == p.Fun {
			d := GetData{ID: "idform:"+fd.ID+";idcontainer:"+fd.form[i][ID]}
			SendMessage("GetDataContainer", d)
		}
	}
}

func ListTable(c *MessageGob) {

   attr:= getid(c.Data.ID) 
	
	tbl := c.Data.Data
	contCatalog := container.NewVBox()
	contDocument := container.NewVBox()
	acc := widget.NewAccordion()
	for _, b := range tbl {
		d := widget.NewButton(b[Synonym], nil)
		d.OnTapped = func() {
			_, param := findButton(d)
			mp := strings.Split(param.Parameters, ";")
			w:=InitForm("TableList", mp[0], "")
			w.Show()
		}
		p := "idform:" + b[ID]
		app_values["main"].Button[b[Synonym]] = ButtonData{Fun: b[Name] + "GenForm", Parameters: p, Widget: d}
		switch b[TypeMetaData] {
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
	app_values[attr["idform"]].Container[app_values[attr["idform"]].form[attr["idcontainer"]][ParentID]] = d
	SetContent(attr["idform"])
}

func SetContent(idform string) {
	var root string

	var top fyne.CanvasObject
	var left fyne.CanvasObject
	var right fyne.CanvasObject
	var bottom fyne.CanvasObject
	top = nil
	bottom = nil
	left = nil
	right = nil
	win := app_values[idform].W
	f := app_values[idform].form
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

						top = app_values[idform].Container[recb]

					}
					if f[recb][TypeContainer] == "left" {
						left = app_values[idform].Container[recb]

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
	//middle := canvas.NewText("content", color.Black)
	// container := container.NewBorderLayout(top, left, nil, nil)

	// 	container.Objects = append(container.Objects, top)

	// 	container.Objects = append(container.Objects, canvas.NewText("content", color.Black))

	// 	container.Objects = append(container.Objects, canvas.NewText("fgdfgdfg", color.Black))

	// 	container.Objects = append(container.Objects, canvas.NewText("eretreg", color.Black))

	//	if top != left {
	//		container.Objects = append(container.Objects, left)
	//	}

	//container.AddObject(topc.)

	// 	content := container.New(lb,top)
	win.SetContent(content)
}

func ToolBarPodsystem2(c *MessageGob) {
	//ToolBarPodsystem1(c)
}
func getid(id string) map[string]string {
	rm := make(map[string]string)
	mp := strings.Split(id, ";")
	for _, rec := range mp {
		if rec != "" {
			mp1 := strings.Split(rec, ":")
			if len(mp1) == 2 {
				switch mp1[0] {

				case "idform":
					rm["idform"] = mp1[1]
				case "idtable":
					rm["idtable"] = mp1[1]
				case "idelem":
					rm["idelem"] = mp1[1]
				case "idcontainer":
					rm["idcontainer"] = mp1[1]
				}
			}

		}
	}
	return rm
}

func ToolBar(c *MessageGob) {
	app := c.Data.Data
	attr := getid(c.Data.ID)

	form := app_values[attr["idform"]]
	fd := app_values[attr["idform"]].form

	tb := ToolBarCreate(attr["idform"], app, color.Gray{230})
	app_values[attr["idform"]].Container[attr["idcontainer"]] = tb
	parent := fd[attr["idcontainer"]][ParentID]
	// update form top fegin
	mp := strings.Split(fd[parent][ChildrensID], ";")
	top := container.New(layout.NewGridLayoutWithRows(2))
	//top.Layout=layout.NewMaxLayout()
	for _, rec := range mp {
		if rec != "" {
			if form.Container[rec] != nil {
				top.Add(form.Container[rec])
			}
		}
	}
	form.Container[parent] = top
	SetContent(attr["idform"])
}





func InitFormLocal(c *MessageGob) {
	app := c.Data.Data
	attr := getid(c.Data.ID)
	fd := app_values[attr["idform"]]
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
	fd := app_values[app[0][ID]]
	fd.W.SetTitle(app[0][Synonym])
	w:=strings.Split(app[0][width], ";")
	if len(w) == 2 {
		wf,_ := strconv.Atoi(w[0])
		hf,_ := strconv.Atoi(w[1])
		fd.W.Resize(fyne.NewSize(float32(wf), float32(hf)))
	}else{
		fd.W.Resize(fyne.NewSize(1000, 100))
	}
}


// InitForm - инициализация формы 1
func InitForm(idform,idtable,idelem string) fyne.Window {
	myWindow := myApp.NewWindow("Телефоны")
	myWindow.Resize(fyne.NewSize(1200, 400))
	app_values[idform] = &FormData{}
	app_values[idform].ID = idform
	app_values[idform].W = myWindow
	app_values[idform].Button = make(map[string]ButtonData)
	app_values[idform].Container = make(map[string]fyne.CanvasObject)
	app_values[idform].Entry = make(map[string]entryForm)
	app_values[idform].Table = make(map[string]*TableOtoko)
	app_values[idform].Tree = make(map[string]*TreeOtoko)
	app_values[idform].form = make(map[string][]string)
	d := GetData{ID: "idform:"+idform}
	SendMessage("GetDescription", d)
	return myWindow
}

//GetDescription - получим описание формы 2
func SendMessage(Action string, d GetData ) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	mes := MessageGob{
		Action: Action,
		Data:   d,
	}
	enc.Encode(mes)
	k := buff.Bytes()
	println(Action)
	Cl.Reci <- k
}



