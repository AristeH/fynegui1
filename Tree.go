package main

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/widget"
)

type TreeOtoko struct {
	ID             string
	IDForm         string
	Tree           *widget.Tree
	TextForTreeUID map[string]string
	TreeUIDMapping map[string][]string
}


func (t TreeOtoko) makeTree() *widget.Tree {
	childUIDs := func(uid string) (c []string) {
		return t.TreeUIDMapping[uid]
	}

	createNode := func(branch bool) (o fyne.CanvasObject) {
		return widget.NewLabel("")
	}

	// It's a branch if uid exists, and has sub-values
	isBranch := func(uid string) (ok bool) {
		if _, ok := t.TreeUIDMapping[uid]; ok {
			if len(t.TreeUIDMapping[uid]) > 0 {
				return true
			}
		}
		return false
	}

	updateNode := func(uid string, branch bool, node fyne.CanvasObject) {
		node.(*widget.Label).SetText(t.TextForTreeUID[uid])
	}

	return widget.NewTree(childUIDs, isBranch, createNode, updateNode)
}
