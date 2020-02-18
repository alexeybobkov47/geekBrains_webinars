package main

import (
	"encoding/json"
	"log"
)

func marshalJSON() error {
	person := PersonJSON{
		FirstName: "Ivan",
		LastName:  "Ivanov",
		Home: AddressJSON{
			Street:   "Leninskie Gory",
			City:     "Moscow",
			PostCode: "1111",
		},
	}

	object, err := json.Marshal(person)
	if err != nil {
		return err
	}

	log.Printf("%s", object)
	return nil
}

func unmarshalJSON() error {
	body, err := openAndReadFile(infoJSON)
	if err != nil {
		return err
	}

	person := new(PersonJSON)
	if err := json.Unmarshal(body, person); err != nil {
		return err
	}

	log.Printf(
		"LastName %v; FirstName %v; PostalCode: %v; Type PostalCode: %T", person.LastName, person.FirstName,
		person.Home.PostCode, person.Home.PostCode)
	return nil
}
