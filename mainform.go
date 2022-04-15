package main

import (
	"bytes"
	"encoding/gob"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetTableList() {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	d := GetData{Table: "md_tabels"}
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
	app := c.Data
	ch := container.NewHBox()
	for _, b := range app.Data {
		but := widget.NewButton(b[1], func() {
			toolMain21(b[0])
		})
		ch.Add(but)
	}
}

func mainform() fyne.Window {
	RegFunc("SetMenu", SetMenu)
	RegFunc("SetTableList", SetTableList)
	var top *fyne.Container
	var left *fyne.Container
	myWindow := myApp.NewWindow("Телефоны")
	myWindow.Resize(fyne.NewSize(1200, 400))
	app_values["main"] = &FormData{}
	app_values["main"].W = myWindow
	app_values["main"].Button = make(map[string]ButtonData)
	middle := canvas.NewText("content", color.White)
	content := container.New(layout.NewBorderLayout(top, nil, left, nil),
		top, left, middle)
	myWindow.SetContent(content)
	GetMenu()
	return myWindow
}
