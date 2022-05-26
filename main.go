package main

import (
	// "context"
	"context"
	"fmt"
	"fynegui/ent"
	"fynegui/ent/migrate"

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
	client, _ = ent.Open("sqlite3", "C:/проект/fynegui/md.db?_fk=1")

	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		WriteLog(fmt.Sprintf("failed creating schema resources: %v", err))
	}
	RegFunc("GetFile", GetFile)
	RegFunc("PutData", PutData)
	RegFunc("SetMetaDataApp", SetMetaDataApp)
	go connectServer()
	
	myWindow := mainform()
	GetMetaDataApp()
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


