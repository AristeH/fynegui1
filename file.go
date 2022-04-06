package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

)

type fileUP struct{
	Name string
	Dir string
	Content []byte
} 

func GetFile(param []byte) []byte {
	var p *fileUP
	json.Unmarshal(param, &p)
	file, err := os.Create(p.Name)
     if err != nil {
        log.Fatal(err)
    }
	buffer := bufio.NewWriter(file)
	buffer.Write(p.Content)
	defer file.Close()
	return nil

}