package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filesToConcat := os.Args[1:]
	var buf bytes.Buffer
	for _, file := range filesToConcat {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		buf.Write(b)
	}
	err := ioutil.WriteFile("merged.txt", buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
