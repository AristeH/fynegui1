package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"os"
)


func GetFile(c *MessageGob){
	p := &c.File
	file, err := os.Create(p.Name)
     if err != nil {
        log.Fatal(err)
    }
	buffer := bufio.NewWriter(file)
	buffer.Write(p.Content)
	defer file.Close()
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
    return nil
}
