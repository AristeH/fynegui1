package main

import (
	"context"
	"fynegui/ent"
	"fynegui/ent/mdsubsystems"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/widget"
	_ "github.com/mattn/go-sqlite3"
)

var left fyne.CanvasObject

// список форм
var app_values = make(map[string]*FormData)
var myApp = app.New()
var Clientsqllite *ent.Client

// toolMain функция отображающая подсистемы
func toolMain() *fyne.Container {
	client, _ := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
	defer client.Close()
	tbl, err := client.MDSubSystems.Query().All(context.Background())
	if err != nil {
		println(err)
	}
	ch := container.NewHBox()
	for _, b := range tbl {
		but := widget.NewButton(b.Namerus, func() {
			toolMain21(b.ID)
		})
		ch.Add(but)
	}
	return ch
}

// toolMain функция отображающая таблицы подсистемы
func toolMain21(sub string) fyne.CanvasObject {

	client, _ := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
	defer client.Close()
	tbl, err := client.MDSubSystems.Query().Where(mdsubsystems.IDEQ(sub)).QueryMdtables().All(context.Background())
	if err != nil {
		println(err)
	}
	ch := widget.NewAccordion()
	contCatalog := container.NewVBox()
	contDocument := container.NewVBox()

	for _, b := range tbl {
		d := widget.NewButton(b.Namerus, nil)
		d.OnTapped = func() {
			param, _ := findButton(d)
			mp := strings.Split(param, ";")
			GenForm(mp[0], mp[1])
		}
		p := b.Nameeng + ";" + "0005bfbd-e65c-11e8-8828-3440b5b05858"
		app_values["main"].Button[b.Namerus] = ButtonData{Fun: b.Nameeng + "GenForm", Parameters: p, Widget: d}
		switch b.Type {
		case "Справочник":
		contCatalog.Add(d)

		case "Документ":
			contDocument.Add(d)

		}
	}

	ch.Append(&widget.AccordionItem{Title:  "Документы",  Detail: contDocument},)
	ch.Append(&widget.AccordionItem{Title:  "Справочники",Detail: contCatalog},)
	

	return ch
}

func main() {
	RegFunc("GetFile", GetFile)
	RegFunc("PutData", PutData)
	RegFunc("GetMetaData", GetMetaData)
	go connectServer()
	myWindow := mainform()
	myWindow.ShowAndRun()
}


