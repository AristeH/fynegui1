package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fynegui/pkg/logging"
)

var (
	// список форм
	appValues       = make(map[string]*FormData)
	myApp           = app.New()
	logger          logging.Logger
	activeContainer *TableOtoko
)

func main() {
	logger = logging.GetLogger()
	go connectServer()
	myApp.Settings().SetTheme(theme.DefaultTheme())
	RegFunc("uc", UpdateForm) // обновим форму

	//RegFunc("InitFormLocal", GetDataContainer) //Получим структуру создаваемой формы
	RegFunc("FormStyle", FormStyle) //Получим описание формы
	RegFunc("Toolbar", ToolBar)
	RegFunc("Accordion", Accordion)
	RegFunc("Table", Table)
	RegFunc("FormDescription", FormDescription)

	// создадим форму

	myWindow := InitForm("main")
	d := GetData{Form: "main", Action: "FormDescription"}
	UpdateContainer(d)
	myWindow.ShowAndRun()

}
