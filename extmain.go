package main

import (
	"bytes"
	"encoding/gob"
	//"encoding/json"
	"fmt"

	"os"

	"github.com/gorilla/websocket"
)
var sLogName = "aristeh.log"
var mfu map[string]func([]byte) []byte

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



func Init(sOpt string, constr string, quit chan string) int {
	VCH = make(chan string)
	CH = quit
	connectServer()
	return 0

}

func connectServer() string {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080/telephon", nil)
	if err != nil {
		WriteLog(fmt.Sprintf("Dial error:  (%s)", err))
		return "Error connection"
	}
	Cl = Client{id: "", socket: conn, Send: make(chan []byte), Reci: make(chan []byte)}
	go readC()
	go write()
	// Action,_ := json.Marshal("login")
	// mes := Message{
	// 	Action: Action,
	// }

	// jsonMessage, _ := json.Marshal(&mes)
	// Cl.Reci <- jsonMessage

	// mes = Message{
	// 	Action: "GetFile",
	// 	Parameters:[]byte("Name:md.db"),
	// }

	// jsonMessage, _ = json.Marshal(&mes)
	// Cl.Reci <- jsonMessage
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
				WriteLog(fmt.Sprintf("error: %v", err))
			}
			break
		}
 		z:=bytes.NewBuffer(message)
		out := MessageGob{}
		dec := gob.NewDecoder(z)
		dec.Decode(&out)
		go Runproc(&out)

	}
}

// RegFunc adds the fu func to a map of functions,
func RegFunc(sName string, fu func([]byte) []byte) {
	if mfu == nil {
		mfu = make(map[string]func([]byte) []byte)
	}
	mfu[sName] = fu
}

// Runproc выполним процедуру
func Runproc(c *MessageGob) {
	if fnc, bExist := mfu[c.Action]; bExist {
		var ap []byte
		if len(c.Parameters) > 1 {
			ap =[]byte( c.Parameters)
		}
		fnc(ap)
	}
}

// WriteLog writes the sText to a log file aristeh.log.
func WriteLog(sText string) {
	f, err := os.OpenFile(sLogName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString(sText)
}

func write() {
	defer func() {
		Cl.socket.Close()
	}()

	for {
		 message, ok := <-Cl.Reci
			println("пишем в сокет")
			if !ok {
				err := Cl.socket.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					WriteLog(fmt.Sprintf("WebSocket Close Error: (%s)", err))
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


