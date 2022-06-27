package main

import (
	"bytes"
	"encoding/gob"

	"fmt"

	"github.com/gorilla/websocket"
)

var sLogName = "aristeh.log"
var mfu map[string]func(*MessageGob)
var mfuLocal map[string]func(*FormData, *ButtonData)

// Client  - структура
type Client struct {
	id     string
	socket *websocket.Conn
	Send   chan []byte
	Reci   chan []byte
}

var Cl Client

var CH chan string
var VCH chan string

type MessageGob struct {
	Action     string  // Имя функуции
	Parameters []byte  //параметры??
	Data       GetData // Сведенеия о данных
	File       FileUP  // сведения о передаваемом файле
}

type GetData struct {
	ID        string
	Container string
	Data      [][]string
}

type FileUP struct {
	Name    string
	Dir     string
	Content []byte
}

func Init(sOpt string, constr string, quit chan string) int {
	VCH = make(chan string)
	CH = quit
	connectServer()
	return 0
}

func connectServer() string {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/telephon", nil)
	if err != nil {
		logger.Infof(fmt.Sprintf("Dial error:  (%s)", err))
		return "Error connection"
	}
	Cl = Client{id: "", socket: conn, Send: make(chan []byte), Reci: make(chan []byte)}
	go readC()
	go write()
	return ""
}

func readC() {
	defer func() {
		Cl.socket.Close()
	}()
	for {
		mt, message, err := Cl.socket.ReadMessage()
		if err != nil {
			fmt.Println("read mt:", mt)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Infof(fmt.Sprintf("error: %v", err))
			}
			break
		}
		z := bytes.NewBuffer(message)
		out := MessageGob{}
		dec := gob.NewDecoder(z)
		dec.Decode(&out)
		go Runproc(&out)
	}
}

// RegFunc adds the fu func to a map of functions,
func RegFunc(sName string, fu func(*MessageGob)) {
	if mfu == nil {
		mfu = make(map[string]func(*MessageGob))
	}
	mfu[sName] = fu
}

// Runproc выполним процедуру
func Runproc(c *MessageGob) {
	if fnc, bExist := mfu[c.Action]; bExist {
		fnc(c)
	}
}

func write() {
	defer func() {
		Cl.socket.Close()
	}()

	for {
		message, ok := <-Cl.Reci
		if !ok {
			err := Cl.socket.WriteMessage(websocket.CloseMessage, []byte{})
			if err != nil {
				logger.Infof(fmt.Sprintf("WebSocket Close Error: (%s)", err))
			}
			return
		}
		err := Cl.socket.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("ошибка write:", err)
			break
		}
		if err := Cl.socket.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			fmt.Println("ошибка ping:", err)
			return
		}
	}
}

// RegFuncLocal adds the fu func to a map of functions,
func RegFuncLocal(sName string, fu func(*FormData, *ButtonData)) {
	if mfuLocal == nil {
		mfuLocal = make(map[string]func(*FormData, *ButtonData))
	}
	mfuLocal[sName] = fu
}

// RunprocLocal выполним процедуру
func RunprocLocal(fd *FormData, sName *ButtonData) {
	if fnc, bExist := mfuLocal[sName.Fun]; bExist {
		fnc(fd, sName)
	}
}
