package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"fynegui/ent"
	"fynegui/ent/mdsubsystems"
	"fynegui/ent/mdtabel"
	//"fynegui/ent/mdtypetabel"
	"fynegui/ent/migrate"
	"image/color"
	"strings"
	//"image/color"
	"fyne.io/fyne/v2"
	//	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	//	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var client *ent.Client

// константы метаданные для вывода интерфейса пользователя, номера столбца для вывода
const (
	ID            int = 0  //гуид метаданного(подсистема, таблица, реквизит, форма)
	Nameeng       int = 1  //имя англ
	Synonym       int = 2  //синоним
	ParentID      int = 3  //гуид родителя или владельца метаданного
	OrderOutput   int = 4  //порядок вывода элементов интерфейса
	MestoOutput   int = 5  //область вывода{top,bottom,left,right,middle,""}
	TypeContainer int = 6  //имя контейнера, виджета
	NameContainer int = 7  //имя заголовка контейнера
	width         int = 8  //ширина
	ChildrensID   int = 9  //иерархия дети( подсистемы, табличные части, реквизиты)
	TypeMetaData  int = 10 // тип метаданного
	Fun           int = 11 // тип метаданного
)

func GetTableList(pod string) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	d := GetData{Table: "md_tabels", Type: pod}
	mes := MessageGob{
		Action: "GetTableList",
		Data:   d,
	}
	enc.Encode(mes)
	k := buff.Bytes()
	println(k)
	Cl.Reci <- k
}

//GetMenu - получим описание метаданных приложения
func GetMetaDataApp() {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	d := GetData{Table: "md_sub_systems"}
	mes := MessageGob{
		Action: "GetMetaData",
		Data:   d,
	}
	enc.Encode(mes)
	k := buff.Bytes()
	println(k)
	Cl.Reci <- k
}

func SetMetaDataApp(c *MessageGob) []byte {
	client, _ = ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")
	ctx := context.Background()
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		WriteLog(fmt.Sprintf("failed creating schema resources: %v", err))
	}
	app := c.Data
	for _, b := range app.Data {
		if len(b) > 0 {
			if b[TypeMetaData] == "Подсистема" {
				nw := client.MDSubSystems.Create()
				nw.SetNameeng(b[Nameeng])
				nw.SetSynonym(b[Synonym])
				nw.SetID(b[ID])
				nw.SetPor("0")
				if b[ParentID] != "" {
					nw.SetParentMdsubsystemsID(b[ParentID])
				}
				tabl := strings.Split(b[ChildrensID], ";")
				for _, rec := range tabl {
					Mdtabel, _ := client.MDTabel.Query().Where(mdtabel.IDEQ(rec)).Only(ctx)
					if Mdtabel != nil {
						nw.AddMdtables(Mdtabel).Save(ctx)
					}
				}
				nw.Save(ctx)
				err := nw.OnConflict().UpdateNewValues().Exec(ctx)
				if err != nil {
					fmt.Println(err.Error())
				}
			}

			if b[TypeMetaData] == "Типы" {
				nw := client.MDTypeTabel.Create()
				nw.SetID(b[ID])
				nw.SetNameeng(b[Nameeng])
				nw.SetSynonym(b[Synonym])
				nw.SetPor(b[OrderOutput])
				err := nw.OnConflict().UpdateNewValues().Exec(ctx)
				if err != nil {
					fmt.Println(err.Error())
				}
			}

			if b[TypeMetaData] == "Таблица" {
				nw := client.MDTabel.Create()
				nw.SetNameeng(b[Nameeng])
				nw.SetSynonym(b[Synonym])
				nw.SetID(b[ID])
				nw.SetPor(b[OrderOutput])
				nw.SetFile("")
				if b[ParentID] != "" {
					nw.SetParentMdtabelID(b[ParentID])
				}
				nw.Save(ctx)
				if b[TypeContainer] != "" {
					nw.SetTypesID(b[TypeContainer])
				}
				err := nw.OnConflict().UpdateNewValues().Exec(ctx)
				if err != nil {
					fmt.Println(err.Error())
				}
			}

			if b[TypeMetaData] == "Реквизит" {
				nw := client.MDRekvizit.Create()
				nw.SetNameeng(b[Nameeng])
				nw.SetSynonym(b[Synonym])
				nw.SetID(b[ID])
				nw.SetPor(b[OrderOutput])
				nw.SetType(b[TypeContainer])
				nw.SetWidthSpisok(60)
				if b[TypeContainer] != "" {
					nw.SetOwnerID(b[ChildrensID])
				}
				err := nw.OnConflict().UpdateNewValues().Exec(ctx)
				if err != nil {
					fmt.Println(err.Error())
				}
			}


		}
	}

	SetContent("main")
	return nil
}

func SetContent(w string) {
	var top *fyne.Container
	win := app_values[w].W
	var b [][]string
	var b1 [][]string
	var f []*ent.MDSubSystems
	//var t []*ent.MDTabel
	ctx := context.Background()
	f, _ = client.MDSubSystems.Query().Where(mdsubsystems.ParentIsNil()).All(ctx)

	for i := range f {
		output := make([]string, 12)
		output[ID] = f[i].ID
		output[Nameeng] = f[i].Nameeng
		output[Synonym] = f[i].Synonym
		output[OrderOutput] = f[i].Por
		output[ParentID] = f[i].Parent
		output[Fun] = "podsystem"
		b = append(b, output)
	}

	top = container.NewVBox()
	tb := ToolBarCreate("main", b, color.Gray{230})
	app_values[w].ToolBar["main1"] = tb
	top.Add(tb)
	f, _ = client.MDSubSystems.Query().Where(mdsubsystems.ParentEQ(f[0].ID)).All(ctx)
	for i := range f {
		output := make([]string, 12)
		output[ID] = f[i].ID
		output[Nameeng] = f[i].Nameeng
		output[Synonym] = f[i].Synonym
		output[OrderOutput] = f[i].Por
		output[ParentID] = f[i].Parent
		output[Fun] = "tabl"
		b1 = append(b1, output)
	}
	tb = ToolBarCreate("main", b1, color.Gray{240})
	app_values[w].ToolBar["main2"] = tb
	top.Add(tb)
	ch := toolMain21(f[0].ID)
	content := container.New(layout.NewBorderLayout(top, nil, ch, nil), top, ch)
	win.SetContent(content)

}

// toolMain функция отображающая таблицы подсистемы
func toolMain21(sub string) fyne.CanvasObject {

	tbl, err := client.MDSubSystems.Query().Where(mdsubsystems.IDEQ(sub)).QueryMdtables().All(context.Background())
	if err != nil {
		println(err)
	}
	ch := widget.NewAccordion()
	contCatalog := container.NewVBox()
	contDocument := container.NewVBox()

	for _, b := range tbl {
		d := widget.NewButton(b.Synonym, nil)
		d.OnTapped = func() {
			param, _ := findButton(d)
			mp := strings.Split(param.Parameters, ";")
			GenForm(mp[0], mp[1])
		}
		p := b.Nameeng + ";" + "0005bfbd-e65c-11e8-8828-3440b5b05858"
		app_values["main"].Button[b.Synonym] = ButtonData{Fun: b.Nameeng + "GenForm", Parameters: p, Widget: d}
		switch b.TypesID {
		case "Справочник":
			contCatalog.Add(d)

		case "Документ":
			contDocument.Add(d)

		}
	}

	ch.Append(&widget.AccordionItem{Title: "Документы", Detail: contDocument})
	ch.Append(&widget.AccordionItem{Title: "Справочники", Detail: contCatalog})

	return ch
}

func mainform() fyne.Window {
	RegFunc("SetMetaDataApp", SetMetaDataApp)
	//RegFunc("SetTableList", SetTableList)
	//	var top *fyne.Container
	//	var left *fyne.Container
	myWindow := myApp.NewWindow("Телефоны")
	myWindow.Resize(fyne.NewSize(1200, 400))
	app_values["main"] = &FormData{}
	app_values["main"].W = myWindow
	app_values["main"].Button = make(map[string]ButtonData)
	app_values["main"].ToolBar = make(map[string]*fyne.Container)
	//middle := canvas.NewText("content", color.White)
	//SetContent("main")
	GetMetaDataApp()
	return myWindow
}
