package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Message struct {
	Action     []byte // имя  функции
	Parameters []byte // параметры
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

