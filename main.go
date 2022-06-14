package main

import (
	"fyne.io/fyne/v2/app"
)

// список форм
var app_values = make(map[string]*FormData)
var myApp = app.New()

func main() {
	go connectServer()

	RegFunc("InitFormLocal", InitFormLocal) //Получим структуру создаваемой формы
	RegFunc("InitFormView", InitFormView)   //Получим описание формы
	RegFunc("ToolBar", ToolBar)
	RegFunc("ListTable", ListTable)
	RegFunc("GetFile", GetFile)

	myWindow := InitForm("main", "")
	myWindow.ShowAndRun()

}
