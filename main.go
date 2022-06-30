package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fynegui/pkg/logging"
)

var (
	// список форм
	appValues = make(map[string]*FormData)
	myApp     = app.New()
	logger    logging.Logger
)

func main() {
	logger = logging.GetLogger()
	logger.Infof("InitFormLocal")
	go connectServer()

	RegFunc("InitFormLocal", InitFormLocal) //Получим структуру создаваемой формы
	RegFunc("InitFormView", InitFormView)   //Получим описание формы
	RegFunc("ToolBar", ToolBar)
	RegFunc("AccordionTable", AccordionTable)
	RegFunc("Table", Table)
	RegFunc("GetFile", GetFile)

	myWindow := InitForm("main", "")
	myWindow.ShowAndRun()

}

func Table(c *MessageGob) {
	println(c)
	list := widget.NewTable(
		func() (int, int) {
			return len(c.Data.Data), len(c.Data.Data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(c.Data.Data[i.Row][i.Col])
		})
	appValues[c.Data.ID].Container[c.Data.Container] = list
	createParent(c.Data.ID, c.Data.Container)
	SetContent(c.Data.ID)
}
