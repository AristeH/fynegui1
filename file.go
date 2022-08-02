package main

import (
	"bufio"
	"log"
	"os"
)

func GetFile(c *MessageGob) {
	p := &c.File
	file, err := os.Create(p.Name)
	if err != nil {
		log.Fatal(err)
	}
	buffer := bufio.NewWriter(file)
	buffer.Write(p.Content)
	defer file.Close()
}
