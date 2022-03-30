package main

import (
//	"context"
	"encoding/json"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/widget"
)

func create_form_md(res Message) {
	var form_data map[string]string
	var w fyne.Window
	var ID string
	var fd *FormData = &FormData{}
	var tab *fyne.Container
	var tab1 *fyne.Container
	var tab2 *fyne.Container
	var content1 *fyne.Container
//ctx := context.Background()
	json.Unmarshal([]byte(res.Parameters), &form_data)

	// обновим сведения о форме
	for key, value := range form_data {
		switch key {
		case "ID":
			ID = value
			app_values[ID] = fd
		case "Title":
			w = myApp.NewWindow(value)
			fd.W = w
			fd.Table = make(map[string]*TableOtoko)
		case "Size":
			words := strings.Split(value, ",")
			wf, err := strconv.ParseFloat(words[0], 32)
			if err != nil {
				wf = 300
			}
			hf, err := strconv.ParseFloat(words[0], 32)
			if err != nil {
				hf = 400
			}
			println(wf)
			println(hf)
			//	w.Resize(fyne.NewSize(float32(wf), float32(hf)))
		}
	}
	

	// for _, value := range res.Child {
	// 	//fmt.Println("Key:", key, "Value:", value)
	// 	switch value.Name {
	// 	case "file":
	// 		//MDTabelReadFromJSON(ctx,clientsqllite, value.Body)

	// 		// var bt = TableOtoko{}
	// 		// json.Unmarshal([]byte(value.Body), &bt)
	// 		// bt.wb = make(map[*widget.Button]int)
	// 		// bt.wc = make(map[*widget.Check]widget.TableCellID)
	// 		// bt.we = make(map[*enterEntry]widget.TableCellID)

	// 		// fd.Table[bt.ID] = &bt
	// 		// tab = bt.makeTable()
	// 	}	
	// }

	content1 = container.NewBorder(
		tab,
		tab1,
		tab2,
		nil,
	)

	if tab != nil {
		content1.Add(tab)

	}

	w.SetContent(content1)

	w.Show()

}
