package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if err := unmarshalJSON(); err != nil {
		log.Println(err)
	}
}

const (
	infoJSON = "info.json"
	infoXML  = "info.xml"
)

func openAndReadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}
