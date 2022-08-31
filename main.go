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
	logger.Infof("InitFormLocal")
	go connectServer()
	myApp.Settings().SetTheme(theme.DefaultTheme())
	RegFunc("uc", UpdateForm) // обновим форму

	//RegFunc("InitFormLocal", GetDataContainer) //Получим структуру создаваемой формы
	RegFunc("init", initform) //Получим описание формы
	RegFunc("Toolbar", ToolBar)
	RegFunc("Accordion", Accordion)
	RegFunc("Table", Table)
	//RegFunc("GetFile", GetFile)

	myWindow := InitForm("main", "")
	myWindow.ShowAndRun()

}
