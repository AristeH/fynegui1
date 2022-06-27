package main

//import "github.com/ilyakaznacheev/cleanenv"
import (
	"fyne.io/fyne/v2/app"
	"fynegui/pkg/logging"
)

// список форм
var appValues = make(map[string]*FormData)
var myApp = app.New()
var logger logging.Logger

func main() {
	logger = logging.GetLogger()
	logger.Infof("InitFormLocal")
	go connectServer()

	RegFunc("InitFormLocal", InitFormLocal) //Получим структуру создаваемой формы
	RegFunc("InitFormView", InitFormView)   //Получим описание формы
	RegFunc("ToolBar", ToolBar)
	RegFunc("AccordionTable", AccordionTable)
	RegFunc("TableList", TableList)
	RegFunc("GetFile", GetFile)

	myWindow := InitForm("main", "")
	myWindow.ShowAndRun()

}

func TableList(gob *MessageGob) {
	println(gob)
}
