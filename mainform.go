package main

import (
	"bytes"
	"encoding/gob"
	//"image/color"
	"fyne.io/fyne/v2"
	//	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/recoilme/pudge"
	//	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// константы метаданные для вывода интерфейса пользователя, номера столбца для вывода 
const ( 
	ID          int = 0   //гуид данных
	Nameeng     int = 1   //имя англ
	Namerus     int = 2   //имя рус 
	Synonym     int = 3   //синоним
	ParentID    int = 4   //гуид родителя или владельца
	ParentName  int = 5   //имя англ родителя
	OrderOutput int = 6   //порядок вывода элементов интерфейса
	MestoOutput int = 7   //область вывода{top,bottom,left,right,middle}
	TypeContainer int = 8 //имя контейнера, виджета
	NameHeader int = 9    //имя заголовка
	width       int = 10  //ширина
	ChildrensID int = 11  //иерархия дети
)

func GetTableList(pod string) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	d := GetData{Table: "md_tabels",Type: pod}
	mes := MessageGob{
		Action: "GetTableList",
		Data:   d,
	}
	enc.Encode(mes)
	k := buff.Bytes()
	println(k)
	Cl.Reci <- k
}

func SetTableList(c *MessageGob) []byte {
	defer pudge.CloseAll()
	app := c.Data
	for _, b := range app.Data {
		if len(b) > 0 {
			gui := []string{b[ID], b[Nameeng], b[Namerus], b[Synonym], b[ParentID], b[ParentName],
				b[OrderOutput], b[MestoOutput], b[TypeContainer], b[NameHeader], b[width],b[ChildrensID]}
			pudge.Set("C:/проект/fynegui/gui", "tab"+b[ID], gui)
			but := widget.NewButton( b[Namerus], func() {
				GenForm(b[Nameeng], "")
			})
			app_values["main"].Button[b[Nameeng]] = ButtonData{Fun: b[Nameeng], Parameters: b[Nameeng], Widget: but}
		}
	}
	SetContent("main")
	return nil
}

func GetMenu() {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	d := GetData{Table: "md_sub_systems"}
	mes := MessageGob{
		Action: "GetMainMenu",
		Data:   d,
	}
	enc.Encode(mes)
	k := buff.Bytes()
	println(k)
	Cl.Reci <- k
}

func SetMenu(c *MessageGob) []byte {
	defer pudge.CloseAll()
	app := c.Data
	for _, b := range app.Data {
		if len(b) > 0 {
			gui := []string{b[ID], b[Nameeng], b[Namerus], b[Synonym], b[ParentID], b[ParentName],
				b[OrderOutput], b[MestoOutput], b[TypeContainer], b[NameHeader], b[width],b[ChildrensID]}
			pudge.Set("C:/проект/fynegui/gui", "pod"+b[ID], gui)
			but := widget.NewButton( b[Namerus], func() {
				GetTableList(b[ID])
			})
			app_values["main"].Button[b[Nameeng]] = ButtonData{Fun: b[Nameeng], Parameters: b[Nameeng], Widget: but}
		}
	}
	SetContent("main")
	return nil
}

func SetContent(w string) {
	var top *fyne.Container
	win := app_values["main"].W
	cfg := &pudge.Config{SyncInterval: 0}
	db, err := pudge.Open("C:/проект/fynegui/gui", cfg)
	if err != nil {
		println(err)
	}
		var b [][]string
		keys, _ := db.KeysByPrefix([]byte("pod"), 0, 0, true)
		for _, key := range keys {
			var output []string
			db.Get(key, &output)
			b = append(b, output)
		}
		top = ToolBarCreate("main",b)

	content := container.New(layout.NewBorderLayout(top, nil, nil, nil), top)
	win.SetContent(content)
}

func mainform() fyne.Window {
	RegFunc("SetMenu", SetMenu)
	RegFunc("SetTableList", SetTableList)
	//	var top *fyne.Container
	//	var left *fyne.Container
	myWindow := myApp.NewWindow("Телефоны")
	myWindow.Resize(fyne.NewSize(1200, 400))
	app_values["main"] = &FormData{}
	app_values["main"].W = myWindow
	app_values["main"].Button = make(map[string]ButtonData)
	//middle := canvas.NewText("content", color.White)
	//SetContent("main")
	GetMenu()
	return myWindow
}
