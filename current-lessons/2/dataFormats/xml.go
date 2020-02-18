package main

import (
	"encoding/xml"
	"log"
)

func unmarshalXML() error {
	body, err := openAndReadFile(infoXML)
	if err != nil {
		return err
	}

	person := new(PersonXML)
	if err := xml.Unmarshal(body, person); err != nil {
		return err
	}

	log.Printf(
		"LastName %v; FirstName %v; PhoneNumbers: %v", person.LastName, person.FirstName, person.PhoneNumbers)

	return nil
}

func marshalXML() error {
	person := PersonXML{
		FirstName: "Ivan",
		LastName:  "Ivanov",
		Home: AddressXML{
			Street:   "Leninskie Gory street",
			City:     "Moscow",
			PostCode: "1111",
		},
		PhoneNumbers: []string{"+7 903 3141592", "8 495 271828"},
	}

	body, err := xml.Marshal(person)
	if err != nil {
		return err
	}

	log.Printf("%s", body)
	return nil
}
