package main

import (
	"context"

	"fynegui/ent"
	"fynegui/ent/mdsubsystems"
	//"fynegui/ent/mdtabel"
	"image/color"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"


	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

var left fyne.CanvasObject 

// список форм
var app_values = make(map[string]*FormData)
var myApp = app.New()
var Clientsqllite *ent.Client

func toolMain() *fyne.Container {
	client, _ := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
	tbl, err := client.MDSubSystems.Query().All(context.Background())
	if err != nil {
		println(err)
	}
	ch := container.NewHBox()
	for _, b := range tbl {
		but := widget.NewButton(b.Namerus, func() {
			toolMain21(b.ID )
		})
		ch.AddObject(but)
	}
	return ch
}



func toolMain21(sub string) fyne.CanvasObject {

	client, _ := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
	tbl, err := client.MDSubSystems.Query().Where(mdsubsystems.IDEQ(sub)).QueryMdtables().All(context.Background())
	if err != nil {
		println(err)
	}
	ch := widget.NewAccordion()

cont := container.NewVBox()
	for _, b := range tbl {
		
		but := widget.NewButton(b.Namerus, func() {
			GenForm(b.Nameeng, "0046247f-bd7a-11e7-823e-1c98ec28debf")
		})
	cont.AddObject(but)

	
	}
			ch.Append(&widget.AccordionItem{
			Title:  "таблицы",
			Detail: cont,},
		)
	return ch
}


func main() {
	RegFunc("GetFile", GetFile)
	RegFunc("PutData", PutData)
	go connectServer()

	myWindow := myApp.NewWindow("TabContainer Widget")
	top := toolMain()

	//top := canvas.NewText("top bar", color.White)
	left = toolMain21("316bec67-1cce-43a2-ae92-b841da8bf090")
	middle := canvas.NewText("content", color.White)
	content := container.New(layout.NewBorderLayout(top, nil, left, nil),
		top, left, middle)
	myWindow.Resize(fyne.NewSize(1200, 400))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func makeTable(IDForm, IDTable string) *TableOtoko {
	var t = make([][]string, 3)
	for i := range t {
		t[i] = []string{strconv.Itoa(i)}
	}
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

func Example_Todo() {
	// client, err := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
	// if err != nil {
	// 	WriteLog(fmt.Sprintf("failed opening connection to sqlite: %v", err))
	// }
	// defer client.Close()
	// // Run the auto migration tool.
	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	WriteLog(fmt.Sprintf("db ->failed creating schema resources: %v", err))
	// }
	// Clientsqllite = client
	// ctx := context.Background()
	// ps1, err := Clientsqllite.MDSubSystems.Query().All(ctx)
	// fmt.Printf(ps1[0].Synonym)
	// if err != nil {
	// 	WriteLog(fmt.Sprintf("tbl->Dial error:  (%s)", err))
	// 	return
	// }
	// if err != nil {
	// 	WriteLog(fmt.Sprintf("Connect ent error:  (%s)", err))
	// }
	//GenFormElem("Phone", "eddcc74e-655d-11eb-9325-10f0058e0aed")
	GenForm("Department", "0046247f-bd7a-11e7-823e-1c98ec28debf")
}
