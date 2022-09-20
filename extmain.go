package main

import (
	"bytes"
	"encoding/gob"

	"fmt"

	"github.com/gorilla/websocket"
)

var mfu map[string]func(*GetData)
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

type GetData struct {
	Form            string
	Container       string
	Action          string
	Data            [][]string
	DataDescription [][]string
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
		out := GetData{}
		dec := gob.NewDecoder(z)
		dec.Decode(&out)
		//if out.Action == "" {
		logger.Infof("<-:" + out.Action + " f:" + out.Form + " c:" + out.Container)
		//}
		if out.Form != "" {
			logger.Infof("прочитано:" + out.Action + " f:" + out.Form + " c:" + out.Container)
			if out.Action == "" {
				out.Action = appValues[out.Form].form[out.Container][TypeContainer]
			}
			go Runproc(&out)
		}
	}
}

// RegFunc adds the fu func to a map of functions,
func RegFunc(sName string, fu func(*GetData)) {
	if mfu == nil {
		mfu = make(map[string]func(*GetData))
	}
	mfu[sName] = fu
}

// Runproc выполним процедуру
func Runproc(c *GetData) {

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

			}
			return
		}
		logger.Infof(fmt.Sprintf("write"))
		err := Cl.socket.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("ошибка write:", err)
			break
		}
		//if err := Cl.socket.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
		//	logger.Infof(fmt.Sprintf(шибка ping: (%s)", err))
		//	return
		//}
	}
}

func UpdateContainer(param GetData) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	enc.Encode(param)
	k := buff.Bytes()
	Cl.Reci <- k
}
