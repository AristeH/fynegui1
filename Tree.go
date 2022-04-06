package main

import (
	"fyne.io/fyne/v2/widget"
)

type TreeOtoko struct {
	ID             string
	IDForm         string
	Tree           *widget.Tree
	TextForTreeUID map[string]string
	TreeUIDMapping map[string][]string
}



