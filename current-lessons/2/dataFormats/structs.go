package main

import (
	"encoding/json"
	"strconv"
)

type PersonJSON struct {
	FirstName    string      `json:"firstName"`
	LastName     string      `json:"lastName"`
	Home         AddressJSON `json:"address"`
	PhoneNumbers []string    `json:"phoneNumbers,omitempty"`
}

type AddressJSON struct {
	Street   string     `json:"streetAddress"`
	City     string     `json:"city"`
	PostCode PostalCode `json:"postalCode"`
}

type PostalCode string

func (code *PostalCode) UnmarshalJSON(data []byte) error {
	var c int
	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}

	*code = PostalCode(strconv.Itoa(c))
	return nil
}

func (code PostalCode) MarshalJSON() ([]byte, error) {
	c, err := strconv.Atoi(string(code))
	if err != nil {
		return nil, err
	}

	return json.Marshal(c)
}

type PersonXML struct {
	FirstName    string     `xml:"firstName"`
	LastName     string     `xml:"lastName"`
	Home         AddressXML `xml:"address"`
	PhoneNumbers []string   `xml:"phoneNumbers>phoneNumber"`
}

type AddressXML struct {
	Street   string     `xml:"streetAddress"`
	City     string     `xml:"city"`
	PostCode PostalCode `xml:"postalCode"`
}
