package main

import (
	"context"
	"strings"

	"fynegui/ent"
	"fynegui/ent/mdsubsystems"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/widget"
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
		p := b.Nameeng + ";" + "0046247f-bd7a-11e7-823e-1c98ec28debf"
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
	go connectServer()

	myWindow := myApp.NewWindow("TabContainer Widget")
	myWindow.Resize(fyne.NewSize(1200, 400))
	app_values["main"] = &FormData{}
	app_values["main"].W = myWindow
	app_values["main"].Button = make(map[string]ButtonData)
	top := toolMain()

	//top := canvas.NewText("top bar", color.White)
	left = toolMain21("316bec67-1cce-43a2-ae92-b841da8bf090")
	middle := canvas.NewText("content", color.White)
	content := container.New(layout.NewBorderLayout(top, nil, left, nil),
		top, left, middle)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func makeTable(IDForm, IDTable string) *TableOtoko {
	var t = make([][]string, 3)
	var TO = TableOtoko{}
	TO.ColumnsName = []string{"node_0"}
	TO.ColumnsType = []string{"label"}
	TO.ColumnsWidth = []float32{40}
	TO.AlterRowColor = color.Gray{250}
	TO.HeaderColor = color.Gray{80}
	TO.RowColor = color.Gray{200}
	TO.Data = t
	TO.Edit = true
	TO.ID = IDTable
	TO.IDForm = IDForm
	TO.wb = make(map[*widget.Button]int)
	TO.wc = make(map[*widget.Check]widget.TableCellID)
	TO.we = make(map[*enterEntry]widget.TableCellID)
	return &TO
}
