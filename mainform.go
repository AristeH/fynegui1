package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
	"strings"
)

// константы метаданные для вывода интерфейса пользователя, номера столбца для вывода
const (
	ID            int = 0  //гуид метаданного(подсистема, таблица, реквизит, форма)
	Name          int = 1  //имя англ
	Synonym       int = 2  //синоним
	ParentID      int = 3  //гуид родителя или владельца метаданного
	OrderOutput   int = 4  //порядок вывода элементов интерфейса
	Icon          int = 5  //область вывода{top,bottom,left,right,middle,""}
	TypeContainer int = 6  //имя контейнера, виджета
	NameContainer int = 7  //имя заголовка контейнера
	Style         int = 8  //ширина
	ChildrensID   int = 9  //иерархия дети( подсистемы, табличные части, реквизиты)
	Parameters    int = 10 // тип метаданного
	Fun           int = 11 // тип метаданного
)

// InitForm - инициализация формы шаблон и объект
func InitForm(form string) fyne.Window {
	//var br [][]string
	logger.Tracef("Инициализация формы f:" + form)
	if val, ok := appValues[form]; ok {
		return val.W
	}
	myWindow := myApp.NewWindow("")
	myWindow.Resize(fyne.NewSize(1200, 400))
	appValues[form] = &FormData{
		ID:        form,
		W:         myWindow,
		Button:    make(map[string]ButtonData),
		Container: make(map[string]fyne.CanvasObject),
		Entry:     make(map[string]entryForm),
		Table:     make(map[string]*TableOtoko),
		Tree:      make(map[string]*TreeOtoko),
		form:      make(map[string][]string),
	}
	//myWindow.CenterOnScreen()
	if form != "main" {
		myWindow.SetCloseIntercept(func() {
			myWindow.Hide()
		})
	}
	appValues[form].W = myWindow
	myWindow.Show()
	return myWindow
}

// FormDescription - инициализация структуры формы.
func FormDescription(mes *GetData) {
	app := mes.Data
	fd := appValues[mes.Form]

	for i := range app {
		fd.form[app[i][ID]] = app[i]

		if app[i][ParentID] != "0" {
			fd.Container[strconv.Itoa(i)] = nil
		}
	}
	FormStyle(mes)

}

// FormStyle - заголовок, стиль формы
func FormStyle(mes *GetData) {
	app := mes.DataDescription
	fd := appValues[mes.Form]
	fd.W.SetTitle(app[0][Synonym])
	w := strings.Split(app[0][Style], "x")
	if len(w) == 2 {
		wf, _ := strconv.Atoi(w[0])
		hf, _ := strconv.Atoi(w[1])
		fd.W.Resize(fyne.NewSize(float32(wf), float32(hf)))
	} else {
		fd.W.Resize(fyne.NewSize(1000, 1000))
	}
	fd.W.SetCloseIntercept(func() {
		fd.W.Hide()
	})

	for _, j := range fd.form {
		tc, _ := strconv.Atoi(j[ID])
		mc, _ := strconv.Atoi(mes.Container)
		if j[Parameters] == "true" && tc > mc {
			d := GetData{Form: mes.Form, Action: j[Fun], Container: j[ID]}
			UpdateContainer(d)
			break
		}
	}
}

func UpdateForm(mes *GetData) {
	f := mes.Form
	c := mes.Container
	form := appValues[f].form // map[string][]string
	for i := range form {
		if c == form[i][ID] {
			if fnc, bExist := mfu[form[i][TypeContainer]]; bExist {
				fnc(mes)
			} else {
				logger.Errorf("Не удалось найти функцию: " + form[i][TypeContainer])
			}
			return
		}
	}

}

func nextContainer(mes *GetData) {

	fd := appValues[mes.Form]
	fun := true

	for i := range fd.form {
		tc, _ := strconv.Atoi(fd.form[i][ID])
		mc, _ := strconv.Atoi(mes.Container)
		if fun && tc > mc && fd.form[i][Parameters] == "true" {
			var br [][]string
			output := make([]string, 12)

			output[ID] = fd.form[i][ID] // гуид подсистемы/кнопки
			//
			output[Fun] = fd.form[i][Fun] // функция для кнопки

			br = append(br, output)
			fun = false
			d := GetData{Form: mes.Form, Action: fd.form[i][Fun], Data: br, Container: output[ID]}
			UpdateContainer(d)
			break
		}
	}
	if fun {
		SetContent(mes.Form)
	}

}

func ToolBar(mes *GetData) {
	f := mes.Form
	c := mes.Container
	tb := ToolBarCreate(f, c, mes.Data, color.Gray{Y: 230})
	appValues[f].Container[c] = tb
	createParent(f, appValues[f].form[c][ParentID])
	nextContainer(mes)
}

func Accordion(mes *GetData) {
	f := mes.Form
	c := mes.Container
	tbl := mes.Data
	fd := appValues[f]
	contCatalog := container.NewVBox()
	contDocument := container.NewVBox()
	acc := widget.NewAccordion()
	for _, b := range tbl {
		d := widget.NewButton(b[Synonym], nil)
		mp := make(map[string]string)
		mp["form"] = b[ID]
		mp["type"] = b[NameContainer]
		fd.Button[b[ID]] = ButtonData{Fun: b[Fun], Container: c, Parameters: mp, Widget: d}
		d.OnTapped = func() {
			_, param := findButton(d)

			//mp := strings.Split(param.Parameters, ";")
			w := InitForm(param.Parameters["form"])
			d := GetData{Form: param.Parameters["form"], Action: "FormDescription"}
			UpdateContainer(d)
			w.Show()
			w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
				if activeContainer == nil {
					return
				}
				i := activeContainer.Selected
				switch k.Name {
				case "Down":
					if len(activeContainer.Data)-1 > activeContainer.Selected.Row {
						activeContainer.Selected = widget.TableCellID{Col: i.Col, Row: i.Row + 1}
					}
				case "Up":
					if i.Row > 1 {
						activeContainer.Selected = widget.TableCellID{Col: i.Col, Row: i.Row - 1}
					}
				case "Left":
					r := activeContainer.Selected.Col
					for r >= 1 && activeContainer.ColumnStyle[r-1].Width == 0 && 0 < r {
						r--
					}
					if r >= 1 {
						activeContainer.Selected = widget.TableCellID{Col: r - 1, Row: i.Row}
					}
				case "Right":

					r := activeContainer.Selected.Col
					if len(activeContainer.Data[0])-1 > r {
						r++
						for activeContainer.ColumnStyle[r].Width == 0 && len(activeContainer.Data[0])-1 > r {
							r++
						}

						activeContainer.Selected = widget.TableCellID{Col: r, Row: i.Row}
					}
				case "KP_Enter", "Return":
					Entry := widget.NewEntry()
					Entry.Validator = getValidator(activeContainer.ColumnStyle[i.Row].Type)
					items := []*widget.FormItem{
						widget.NewFormItem("Username", Entry),
					}
					dialog.ShowForm("введите", "", "cancel", items, func(b bool) {
						if !b {
							return
						}
						fmt.Println("KP_Enter", Entry.Text)
						activeContainer.Data[i.Row][i.Col] = Entry.Text
					}, w)

				}
				activeContainer.Table.ScrollTo(activeContainer.Selected)
				//activeContainer.Table.Refresh()
			})
		}
		switch b[TypeContainer] {
		case "Справочник":
			contCatalog.Add(d)
		case "Документ":
			contDocument.Add(d)
		}
	}
	acc.Append(&widget.AccordionItem{Title: "Документы", Detail: contDocument})
	acc.Append(&widget.AccordionItem{Title: "Справочники", Detail: contCatalog})
	d := container.NewVBox(container.NewVScroll(acc))
	d.Layout = layout.NewMaxLayout()
	appValues[f].Container[c] = d
	createParent(f, appValues[f].form[c][ParentID])
	nextContainer(mes)

}

func top(f, c string) {
	form := appValues[f]
	fd := appValues[f].form

	mp := strings.Split(fd[c][ChildrensID], ";")
	top := container.New(layout.NewGridLayoutWithRows(len(mp) - 1))
	for _, rec := range mp {
		if rec != "" {
			if form.Container[rec] != nil {
				top.Add(form.Container[rec])
			}
		}
	}

	appValues[f].Container[c] = top
	createParent(f, fd[c][ParentID])
}

func TableList(f, c string) {
	var t fyne.CanvasObject
	var m fyne.CanvasObject
	form := appValues[f]
	fd := appValues[f].form

	mp := strings.Split(fd[c][ChildrensID], ";")

	if val, ok := form.Container[mp[0]]; ok {
		t = val
	}
	if val, ok := form.Container[mp[1]]; ok {
		m = val
		content := container.NewBorder(
			t,
			nil,
			nil,
			nil,
			m,
		)
		appValues[f].Container[c] = content
	}

	createParent(f, fd[c][ParentID])
}

func border(f, c string) {
	var top fyne.CanvasObject
	var left fyne.CanvasObject
	var right fyne.CanvasObject
	var bottom fyne.CanvasObject
	var middle fyne.CanvasObject

	form := appValues[f]
	fd := appValues[f].form

	mp := strings.Split(fd[c][ChildrensID], ";")

	for _, rec := range mp {
		if rec != "" {
			if form.Container[rec] != nil {
				switch fd[rec][TypeContainer] {
				case "top":
					top = form.Container[rec]

				case "bottom":
					bottom = form.Container[rec]

				case "left":
					left = form.Container[rec]

				case "right":
					right = form.Container[rec]

				case "middle":
					middle = form.Container[rec]

				}
			}
		}
	}
	content := container.New(layout.NewBorderLayout(top, bottom, left, right))
	if left != nil {
		content.Add(left)
	}
	if top != nil {
		content.Add(top)
	}
	if middle != nil {
		content.Add(middle)
	}
	form.Container[c] = content

	//appValues[f].Container[c] = top
	//createParent(f, fd[c][ParentID])
}

func SetContent(f string) {
	var content fyne.CanvasObject
	form := appValues[f].form // map[string][]string
	for i := range form {
		if "init" == form[i][TypeContainer] {
			root := strings.Split(form[i][ChildrensID], ";")
			for _, rec := range root {
				if rec != "" {
					content = appValues[f].Container[rec]
				}
			}
			break
		}
	}

	appValues[f].W.SetContent(content)
	appValues[f].W.Show()
}

func createParent(f, c string) {
	form := appValues[f].form

	switch form[c][TypeContainer] {
	case "top":
		top(f, c)

	case "Border":
		border(f, c)
	case "TableList":
		TableList(f, c)
	case "left":
		top(f, c)
	case "middle":
		top(f, c)
	}
}

// SendMessage -  отправить сообщение серверу
func SendMessage(d GetData) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	enc.Encode(d)
	k := buff.Bytes()
	logger.Infof("ОТПР- Форма:" + d.Form + "Контейнер:" + d.Container + "Функция:" + d.Action + " параметры:" + d.Data[0][Parameters])
	Cl.Reci <- k
}

func UpdateFormContent(param GetData) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	enc.Encode(param)
	Cl.Reci <- buff.Bytes()
}

func FieldsCreate(mes *GetData) {

	var vb []*fyne.Container
	f := mes.Form
	c := mes.Container
	fd := mes.Data
	fEntry := make(map[string]entryForm)
	v := container.New(layout.NewHBoxLayout())

	for i := 0; i < len(fd); i++ {
		for k := len(vb); k < len(fd[i])*2; k++ {
			vb = append(vb, container.New(layout.NewVBoxLayout()))
			vb = append(vb, container.New(layout.NewVBoxLayout()))
		}
		for j := 0; j < len(fd[i]); j++ {
			vb[j*2].Add(widget.NewLabel(fd[i][j])) //.Name
			entry := widget.NewEntry()
			if fd[i][j] == "" {
				entry.PlaceHolder = fd[i][j] //.Title
			} else {
				entry.Text = fd[i][j] //.Value
			}
			entry.Wrapping = fyne.TextWrapOff
			vb[j*2+1].Add(entry)
			fEntry[fd[i][j]] = entryForm{Value: fd[i][j], Widget: entry} //.Value
		}
	}

	for i := 0; i < len(vb); i++ {
		v.Add(vb[i])
	}
	beans := appValues[f]
	beans.Entry = fEntry
	appValues[f] = beans
	for j := 0; j < len(fd[0]); j++ {
		h := container.New(layout.NewVBoxLayout())
		for i := 0; i < len(fd); i++ {
			h.Add(widget.NewLabel(fd[i][j]))
			// entry := widget.NewEntry()
			// entry.PlaceHolder = f[i][j].Title
			// h.Add(entry)
		}
		v.Add(h)
	}
	createParent(f, appValues[f].form[c][ParentID])
	nextContainer(mes)
}
