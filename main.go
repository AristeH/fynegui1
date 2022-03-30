package main

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fynegui/ent"
	_ "github.com/mattn/go-sqlite3"
	"image/color"
	"strconv"
)


// список форм
var app_values = make(map[string]*FormData)
var myApp = app.New()
var Clientsqllite *ent.Client

func main() {
	RegFunc("GetFile", GetFile)
	RegFunc("PutData", PutData)
	go connectServer()
	Example_Todo()
	myWindow := myApp.NewWindow("TabContainer Widget")
	horizontalSplitter := makeTable("tovar", "test")
	t := make(map[string]*TableOtoko)
	tree := make(map[string]*TreeOtoko)
	t["tovar"] = horizontalSplitter
	app_values["test"] = &FormData{Table: t,Tree: tree, W: myWindow}
	//content := widget.NewLabel("text")
	c := horizontalSplitter.makeTable()
	myWindow.Resize(fyne.NewSize(1200, 400))
	myWindow.SetContent(c)
	myWindow.ShowAndRun()
}

func makeTable(IDForm, IDTable string) *TableOtoko {
	var t = make([][]string, 3)
	for i := range t {
		t[i] = []string{strconv.Itoa(i)}
	}
	var TO = TableOtoko{}
	TO.ColumnsName   = []string{"node_0"}
	TO.ColumnsType   = []string{"label"}
	TO.ColumnsWidth  = []float32{40}
	TO.AlterRowColor = color.Gray{250}
	TO.HeaderColor   = color.Gray{80}
	TO.RowColor      = color.Gray{200}
	TO.Data          = t
	TO.Edit          = true
	TO.ID            = IDTable
	TO.IDForm        = IDForm
	TO.wb = make(map[*widget.Button]int)
	TO.wc = make(map[*widget.Check]widget.TableCellID)
	TO.we = make(map[*enterEntry]widget.TableCellID)
	return &TO
}

func Example_Todo() {
	client, err := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
	if err != nil {
		WriteLog(fmt.Sprintf("failed opening connection to sqlite: %v", err))
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		WriteLog(fmt.Sprintf("db ->failed creating schema resources: %v", err))
	}
	Clientsqllite = client
	ctx := context.Background()
	ps1, err := Clientsqllite.MDSubSystems.Query().All(ctx)
	fmt.Printf(ps1[0].Synonym)
	if err != nil {
		WriteLog(fmt.Sprintf("tbl->Dial error:  (%s)", err))
		return
	}
	if err != nil {
		WriteLog(fmt.Sprintf("Connect ent error:  (%s)", err))
	}
	//GenFormElem("Phone", "eddcc74e-655d-11eb-9325-10f0058e0aed")
	GenForm("Department", "0046247f-bd7a-11e7-823e-1c98ec28debf")
}
