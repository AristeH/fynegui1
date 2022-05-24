package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"fynegui/ent"
	"fynegui/ent/mdrekvizit"
	"fynegui/ent/mdtabel"

	//"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Message struct {
	Action     []byte // имя  функции
	Parameters []byte // параметры
}

func GenFormLayout(fd map[string]entryForm, rek []*ent.MDRekvizit) *fyne.Container {
	//var pages = make(map[string]string)
	var columns = make(map[float64]*widget.Form)
	// разделим вывод
	for _, v := range rek {
		if _, ok := columns[v.WidthSpisok]; !ok {
			columns[v.WidthSpisok] = widget.NewForm()
		}
	}
	gri := container.New(layout.NewGridLayout(len(columns)))
	for _, v := range columns {
		gri.Add(v)
	}
	for _, v := range rek {
		if v.Nameeng != "id" && v.Type != "String,0" {
			if strings.HasPrefix(v.Type, "bool") {
				con := widget.NewCheck("", nil)
				columns[v.WidthSpisok].Items = append(columns[v.WidthSpisok].Items, widget.NewFormItem(v.Synonym, con))
			} else {
				input := widget.NewEntry()
				input.SetPlaceHolder(v.Nameeng)
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
				columns[v.WidthSpisok].Items = append(columns[v.WidthSpisok].Items, widget.NewFormItem(v.Synonym, con))
			}
		}
	}
	return gri
}

func GenData(elemname string, id string) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	jsonMessage, _ := json.Marshal([]string{"ID: " + id})
	d := GetData{Table: elemname, ID: id}
	mes := MessageGob{
		Action:     elemname + "GetData",
		Parameters: jsonMessage,
		Data:       d,
	}
	enc.Encode(mes)
	k := buff.Bytes()
	println(k)

	Cl.Reci <- k

}

func PutData(c *MessageGob) []byte {
	app := c.Data
	if app.ID == "" {
		if app.Table == "md_sub_systems"{
			println("jhjhjhjhjhjhjhjhjhjhjhjh")
		}
		client, _ := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
		tbl, _ := client.MDTabel.Query().Where(mdtabel.NameengEQ(app.Table)).All(context.Background())
		rec, err := client.MDRekvizit.Query().Order(ent.Asc(mdrekvizit.FieldPor)).Where(mdrekvizit.OwnerID(tbl[0].ID)).All(context.Background())
		if err != nil {
			WriteLog(fmt.Sprintf("tbl->Dial error:  (%s)", err))
			return nil
		}
		// посчитаем сколько столбцов нужно отобразить
		kolstolb := 0
		for _, field := range rec {
			if field.Nameeng == "id" {
				kolstolb++ // столбец для гуид
				continue
			}
			if field.WidthSpisok > 0 {
				kolstolb++ // столбец с шириной > 0
			}
		}
		kolstolb++ // столбец для нумерации

		ColumnsName := make([]string, kolstolb)
		ColumnsType := make([]string, kolstolb)
		ColumnsWidth := make([]float32, kolstolb)
		ColumnsNameRecive := make([]int, kolstolb)
		ColumnsWidth[1] = 40
		ColumnsType[1] = "string"
		ColumnsName[1] = "N"
		kolstolb = 2
		for _, field := range rec {
			// пропустим столбец нулевой длины
			if field.WidthSpisok == 0 {
				continue
			}
			for i := 0; i < len(app.Data[0]); i++ {
				if strings.EqualFold(app.Data[0][i], field.Nameeng) || strings.EqualFold(app.Data[0][i], field.Nameeng+"Name") {
					if field.Nameeng == "id" {
						ColumnsNameRecive[0] = i
						ColumnsWidth[0] = 0
						ColumnsType[0] = "string"
						ColumnsName[0] = "id"
					} else {
						ColumnsNameRecive[kolstolb] = i
						if strings.HasPrefix(strings.ToUpper(field.Type), "STRING") {
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
						break
					}
				}
			}

		}
		fd := app_values[app.Table]
		ParentID := 0
		Name := 0
		ID := 0
		// name stolb
		tabl := make([][]string, len(app.Data))
		tabl[0] = make([]string, kolstolb)
		copy(tabl[0], ColumnsName)
		tree := false
		for i := 0; i < len(app.Data[0]); i++ {
			if app.Data[0][i] == "ParentID" {
				ParentID = i
				tree = true
			} else if app.Data[0][i] == "Name" {
				Name = i
			} else if app.Data[0][i] == "ID" {
				ID = i
			}
		}
		if tree {
			for i := 1; i < len(app.Data); i++ {
				fd.Tree[app.Table].TextForTreeUID[app.Data[i][ID]] = app.Data[i][Name]
				k := fd.Tree[app.Table].TreeUIDMapping[app.Data[i][ParentID]]
				k = append(k, app.Data[i][ID])
				fd.Tree[app.Table].TreeUIDMapping[app.Data[i][0]] = k
			}
			fd.Tree[app.Table].Tree.Refresh()
		} else {
			fd.Tree[app.Table].Tree.Hide()
		}
		//table
		for i := 1; i < len(app.Data); i++ {
			tabl[i] = make([]string, kolstolb)
			for j := 0; j < len(ColumnsName); j++ {
				if ColumnsType[j] == "Time" {
					tabl[i][j] = app.Data[i][ColumnsNameRecive[j]]
				} else {
					tabl[i][j] = app.Data[i][ColumnsNameRecive[j]]
				}
				ColumnsWidth[j] = float32(len(app.Data[i][ColumnsNameRecive[j]])) * 9
			}
			tabl[i][1] = strconv.Itoa(i)
		}
		ColumnsWidth[0] = 0
		ColumnsWidth[1] = float32(len(strconv.Itoa(len(app.Data)))) * 12

	

		fd.Table[app.Table].ColumnsType = ColumnsType
		fd.Table[app.Table].ColumnsWidth = ColumnsWidth
		fd.Table[app.Table].ColumnsName = ColumnsName

		for ic, v := range ColumnsWidth {
			fd.Table[app.Table].Table.SetColumnWidth(ic, v)
			if ColumnsName[ic] == "Ссылка" {
				fd.Table[app.Table].Table.SetColumnWidth(ic, 0)
			}
		}
		fd.Table[app.Table].Table.Refresh()

	} else {
		for row, st := range app.Data { //получим строку
			for col := range st { // колонку
				if row > 0 {
					for _, v3 := range app_values {
						for k, v4 := range v3.Entry {
							if k == app.Data[0][col] || k+"Name" == app.Data[0][col] {
								v4.Value = app.Data[row][col]
								v4.Widget.SetText(app.Data[row][col])
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

func GenFormTable(NameTable, IDForm string) (f *fyne.Container, t *TableOtoko) {
	Clientsqllite, _ := ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
	ctx := context.Background()
	_, err := Clientsqllite.MDTabel.Query().WithMdrekvizits().Where(mdtabel.NameengEQ(NameTable)).All(ctx)
	if err != nil {
		WriteLog(fmt.Sprintf("tbl->Dial error:  (%s)", err))
	}

	f, t = newTableOtoko(NameTable, IDForm)

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

// Функция для генерации формы таблицы БД
// если id = "" генерится форма списка
// иначе форма элемента
func GenForm(NameTable, id string) {
	var fd FormData = FormData{}

	var top, tabl *fyne.Container
	var tree *widget.Tree

	var entr map[string]entryForm
	// форма элемента
	
	if id != "" {
		app_values[id] = &fd
		top, entr = GenFormElem(NameTable, id)
		fd.Entry = entr
		println(top)
	} else {
		// форма списка
		app_values[NameTable] = &fd
		tree = GenFormTree(NameTable, NameTable)
		if tree != nil {
			var f TreeOtoko
			fd.Tree = make(map[string]*TreeOtoko)
			f.Tree = tree
			f.TextForTreeUID = make(map[string]string)
			f.TreeUIDMapping = make(map[string][]string)
			fd.Tree[NameTable] = &f
		}
		var ta *TableOtoko
		tabl, ta = GenFormTable(NameTable, NameTable)
		if ta != nil {
			fd.Table = make(map[string]*TableOtoko)
			fd.Table[NameTable] = ta
		}

	}

	myWindow := myApp.NewWindow("TabContainer Widget")
	myWindow.Resize(fyne.NewSize(1200, 400))
	fd.W = myWindow
	var horizontalSplitter *container.Split
	var content *fyne.Container


//if tree!=nil{
	horizontalSplitter = container.NewHSplit(tree, tabl)
	content = container.New(layout.NewBorderLayout(nil, nil, nil, nil),
		 horizontalSplitter)
// }else{
// 	content = container.New(layout.NewBorderLayout(top, nil, nil, nil),
// 		top)
// }
	
	//

	// borderLayout := layout.NewBorderLayout(top, nil, left, right)
	// content := fyne.newbNNewContainerWithLayout(border
	// 	Layout, top, nil,left, left, horizontalSplitter)

	myWindow.SetContent(content)
	myWindow.Show()
	//	GenData(NameTable, id)
	GenData(NameTable, "")
	//fd.Tree[eleNameTablemname].Tree.Refresh()
}
