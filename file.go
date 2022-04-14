package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"os"

//	"github.com/gorilla/websocket"

)



func GetFile(c *MessageGob) []byte {
	p := &c.File
	file, err := os.Create(p.Name)
     if err != nil {
        log.Fatal(err)
    }
	buffer := bufio.NewWriter(file)
	buffer.Write(p.Content)
	defer file.Close()
	return nil

}

func GetMetaData(c *MessageGob) []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	jsonMessage, _ := json.Marshal([]string {"Table: MDSubSystems"})
	mes := MessageGob{
		Action:     "MDSubSystemsGob",
		Parameters: jsonMessage,
	}
	enc.Encode(mes)
	k:=buff.Bytes()
	println(k)
	//c:=[]byte("vbh")
//Cl.socket.WriteMessage(websocket.TextMessage, c)

// Action,_ := json.Marshal("login")
// 	mes1 := Message{
// 		Action: Action,
// 	}

// 	jsonMessage, _ := json.Marshal(&mes1)
// 	Cl.Reci <- jsonMessage
//	Cl.Reci <-k
return nil
}
