package main

import (
	"encoding/json"
	"encoding/xml"
	"log"

	"gopkg.in/go-playground/validator.v8"
)

type JsonXML struct {
	Key   string `json:"key" xml:"key"`
	Value string `json:"value" xml:"value" check:"required"`
}

func validation(jsonXml JsonXML) error {
	return validate.Struct(jsonXml)
}

var validate = validator.New(&validator.Config{
	TagName: "check",
})

func main() {
	jx := JsonXML{
		Key: "key",
	}

	if err := validation(jx); err != nil {
		log.Print(err)
		return
	}

	bytesJSON, err := json.Marshal(&jx)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print(string(bytesJSON))

	bytesXML, err := xml.Marshal(&jx)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Print(string(bytesXML))
}
