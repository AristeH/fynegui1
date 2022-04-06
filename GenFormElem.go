package main

import (
	"context"
	"encoding/json"
	"fmt"
	"fynegui/ent"
	"fynegui/ent/mdrekvizit"
	"fynegui/ent/mdtabel"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Message struct {
	Action     []byte // имя  функции
	Parameters []byte // параметры
}

type Getdate struct {
	Table string
	ID    string
	Data  [][]string
}

func GenFormLayout(fd map[string]entryForm, rek []*ent.MDRekvizit) *fyne.Container {
	grid := container.New(layout.NewFormLayout())
	for _, v := range rek {
		if v.Nameeng != "id" && v.Type != "String,0" {
			label := widget.NewLabel(v.Synonym)
			grid.Add(label)
			if strings.HasPrefix(v.Type, "bool") {
				input := widget.NewCheck("", nil)
				grid.Add(input)
			} else {
				input := widget.NewEntry()
				input.SetPlaceHolder(v.Namerus)
				fd[v.Nameeng] = entryForm{Value: "", Widget: input}
				contentb := widget.NewButton("...", func() {
					WriteLog("tapped")
				})
				contentb1 := widget.NewButton("?", func() {
					WriteLog("tapped")
				})
				con := container.NewBorder(nil, nil,
					nil, container.NewHBox(contentb, contentb1),
					// Middle
					input,
				)
				grid.Add(con)
			}
		}
	}

	return container.NewVBox(grid,canvas.NewLine(color.Black))
}

func GenData(elemname string, id string) {
	req := Getdate{Table: elemname, ID: id}
	jsonMessage, _ := json.Marshal(&req)
	Action, _ := json.Marshal(elemname + "GetData")
	mes := Message{
		Action:     Action,
		Parameters: jsonMessage,
	}
	jsonMessage, _ = json.Marshal(&mes)
	Cl.Reci <- jsonMessage
}

func PutData(param []byte) []byte {
	app := Getdate{}
	json.Unmarshal(param, &app)
	if app.ID == "" {
		client, _ := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
		tbl, _ := client.MDTabel.Query().Where(mdtabel.NameengEQ(app.Table)).All(context.Background())

		rec, err := client.MDRekvizit.Query().Order(ent.Asc(mdrekvizit.FieldPor)).Where(mdrekvizit.OwnerID(tbl[0].ID)).All(context.Background())
		if err != nil {
			WriteLog(fmt.Sprintf("tbl->Dial error:  (%s)", err))
			return nil
		}

		kolstolb := 0
		for _, field := range rec {
			if field.Nameeng == "id"{
				kolstolb++
				continue
			}
			if field.WidthSpisok > 0 {
				kolstolb++
			}
		}
		kolstolb++

		ColumnsName := make([]string, kolstolb)
		ColumnsType := make([]string, kolstolb)
		ColumnsWidth := make([]float32, kolstolb)
		ColumnsNameRecive := make([]int, kolstolb)
		ColumnsWidth[0] = 40
		ColumnsType[0] = "string"
		ColumnsName[0] = "N"
		kolstolb = 1
		for _, field := range rec {

			// if field.Nameeng == "id" {
			// 	continue
			// }
			if field.WidthSpisok == 0 {
				continue
			}
			for i := 0; i < len(app.Data[0]); i++ {
				if strings.EqualFold(app.Data[0][i], field.Nameeng) || strings.EqualFold(app.Data[0][i], field.Nameeng+"Name"){
					ColumnsNameRecive[kolstolb] = i
					break
				}

			}
			if field.Nameeng == "id" {
				ColumnsWidth[kolstolb] = 0
				ColumnsType[kolstolb] = "string"
				ColumnsName[kolstolb] = field.Synonym

			} else if strings.HasPrefix( strings.ToUpper(field.Type), "STRING") {
				ColumnsWidth[kolstolb] = float32(field.WidthSpisok)
				ColumnsType[kolstolb] = "string"
				ColumnsName[kolstolb] = field.Synonym
			} else if strings.HasPrefix(field.Type, "bool") {
				ColumnsWidth[kolstolb] = float32(len(field.Synonym))
				ColumnsType[kolstolb] = "bool"
				ColumnsName[kolstolb] = field.Synonym
			} else {
				ColumnsWidth[kolstolb] = float32(field.WidthSpisok)
				ColumnsType[kolstolb] = field.Type
				ColumnsName[kolstolb] = field.Synonym
			}
			kolstolb++

		}
		fd := app_values[app.Table]
		ParentID := 0
		Name := 0
		ID := 0
		// name stolb
		tabl := make([][]string, len(app.Data))
		tabl[0] = make([]string, kolstolb)
		copy(tabl[0],ColumnsName)
		for i := 0; i < len(app.Data[0]); i++ {
			if app.Data[0][i] == "ParentID" {
				ParentID = i
			} else if app.Data[0][i] == "Name" {
				Name = i
			} else if app.Data[0][i] == "ID" {
				ID = i
			}
		}
		//tree
		for i := 1; i < len(app.Data); i++ {
			fd.Tree[app.Table].TextForTreeUID[app.Data[i][ID]] = app.Data[i][Name]
			k := fd.Tree[app.Table].TreeUIDMapping[app.Data[i][ParentID]]
			k = append(k, app.Data[i][ID])
			fd.Tree[app.Table].TreeUIDMapping[app.Data[i][0]] = k
		}
		fd.Tree[app.Table].Tree.Refresh()
		//table
		for i := 1; i < len(app.Data); i++ {
			tabl[i] = make([]string, kolstolb)
			tabl[i][0] = strconv.Itoa(i)

			for j := 1; j < len(ColumnsName); j++ {
				 
				if ColumnsType[j] == "Time"{
					tabl[i][j] = app.Data[i][ColumnsNameRecive[j]]
				}else{
					tabl[i][j] = app.Data[i][ColumnsNameRecive[j]]
				}
				ColumnsWidth[j] = float32( len(app.Data[i][ColumnsNameRecive[j]]))*9
			}
		}
		ColumnsWidth[0] = float32( len(strconv.Itoa(len(app.Data))))*12
		fd.Table[app.Table].Data = tabl

		fd.Table[app.Table].ColumnsType = ColumnsType
		fd.Table[app.Table].ColumnsWidth = ColumnsWidth
		fd.Table[app.Table].ColumnsName = ColumnsName
		fd.Table[app.Table].Table.Refresh()
		for ic, v := range ColumnsWidth {
			fd.Table[app.Table].Table.SetColumnWidth(ic, v)
			if ColumnsName[ic] == "Ссылка"{
				fd.Table[app.Table].Table.SetColumnWidth(ic, 0)
			}
			
		}
		println(ColumnsType)
	} else {
		// обхода двумерного массива
		for i, f := range app.Data {
			for j := range f {
				if i > 0 {
					for _, v3 := range app_values {
						for k, v4 := range v3.Entry {
							if k == app.Data[0][j] || k+"Name" == app.Data[0][j] {
								v4.Value = app.Data[i][j]
								v4.Widget.SetText(app.Data[i][j])
							}
						}
					}
				}

			}
		}
	}
	return nil
}

func GenFormElem(elemname, id string) (f *fyne.Container, fEntry map[string]entryForm) {
	Clientsqllite, _ := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
	ctx := context.Background()
	ps, err := Clientsqllite.MDTabel.Query().WithMdrekvizits().Where(mdtabel.NameengEQ(elemname)).All(ctx)
	if err != nil {
		WriteLog(fmt.Sprintf("tbl->Dial error:  (%s)", err))
	}
	fEntry = make(map[string]entryForm)
	f = GenFormLayout(fEntry, ps[0].Edges.Mdrekvizits)
	return f, fEntry
}

func GenFormTable(NameTable, IDForm string) (f *fyne.Container, t map[string]*TableOtoko) {
	Clientsqllite, _ := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
	ctx := context.Background()
	_, err := Clientsqllite.MDTabel.Query().WithMdrekvizits().Where(mdtabel.NameengEQ(NameTable)).All(ctx)
	if err != nil {
		WriteLog(fmt.Sprintf("tbl->Dial error:  (%s)", err))
	}
	t = make(map[string]*TableOtoko)
	t[NameTable] = makeTable(NameTable, IDForm)
	f = t[NameTable].makeTable()
	return f, t
}

func GenFormTree(NameTree, IDForm string) *widget.Tree {

	childUIDs := func(uid string) (c []string) {
		return app_values[NameTree].Tree[NameTree].TreeUIDMapping[uid]
	}

	createNode := func(branch bool) (o fyne.CanvasObject) {
		return widget.NewLabel("")
	}

	// It's a branch if uid exists, and has sub-values
	isBranch := func(uid string) (ok bool) {
		if _, ok := app_values[NameTree].Tree[NameTree].TreeUIDMapping[uid]; ok {
			if len(app_values[NameTree].Tree[NameTree].TreeUIDMapping[uid]) > 0 {
				return true
			}
		}
		return false
	}

	updateNode := func(uid string, branch bool, node fyne.CanvasObject) {
		node.(*widget.Label).SetText(app_values[NameTree].Tree[NameTree].TextForTreeUID[uid])
	}

	return widget.NewTree(childUIDs, isBranch, createNode, updateNode)
}

func GenForm(elemname, id string) {
	var fd FormData = FormData{}
	var f TreeOtoko
	app_values[elemname] = &fd

	top, entr := GenFormElem(elemname, id)
	fd.Entry = entr

	tabl, k := GenFormTable(elemname, elemname)
	fd.Table = k

	fd.Tree = make(map[string]*TreeOtoko)
	f.Tree = GenFormTree(elemname, elemname)
	f.TextForTreeUID = make(map[string]string)
	f.TreeUIDMapping = make(map[string][]string)
	fd.Tree[elemname] = &f
	//bottom := makeCell()
	//left := makeCell()

	//right := makeCell()

	myWindow := myApp.NewWindow("TabContainer Widget")
	myWindow.Resize(fyne.NewSize(1200, 400))
	fd.W = myWindow

	horizontalSplitter := container.NewHSplit(fd.Tree[elemname].Tree, tabl)
	content := container.New(layout.NewBorderLayout(top, nil, nil, nil),
		top, horizontalSplitter)
	//

	// borderLayout := layout.NewBorderLayout(top, nil, left, right)
	// content := fyne.newbNNewContainerWithLayout(border
	// 	Layout, top, nil,left, left, horizontalSplitter)

	myWindow.SetContent(content)
	myWindow.Show()
	GenData(elemname, id)
	GenData(elemname, "")
	fd.Tree[elemname].Tree.Refresh()
}
