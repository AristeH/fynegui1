package main

import (
	// "context"
	"fynegui/ent"
	// "fynegui/ent/mdsubsystems"
	// "strings"

	// "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)



// список форм
var app_values = make(map[string]*FormData)
var myApp = app.New()
var Clientsqllite *ent.Client

var mfulocal map[string]func(*FormData,  string) []byte

func main() {
	RegFunc("GetFile", GetFile)
	RegFunc("PutData", PutData)
	go connectServer()
	myWindow := mainform()
	myWindow.ShowAndRun()
}


// RegFunc adds the fu func to a map of functions,
func RegFuncLocal(sName string, fu func(*FormData,  string) []byte) {
	if mfulocal == nil {
		mfulocal = make(map[string]func(*FormData,  string) []byte)
	}
	mfulocal[sName] = fu
}

// Runproc выполним процедуру
func RunprocLocal(fd *FormData, sName string) {
	if fnc, bExist := mfulocal[sName]; bExist {
		fnc(fd,sName)
	}
}


