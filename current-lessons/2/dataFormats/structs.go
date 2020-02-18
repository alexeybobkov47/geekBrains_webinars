package dataFormats

type PersonJSON struct {
	FirstName    string      `json:"firstName"`
	LastName     string      `json:"lastName"`
	Home         AddressJSON `json:"address"`
	PhoneNumbers []string    `json:"phoneNumber"`
}

type AddressJSON struct {
	Street   string     `json:"streetAddress"`
	City     string     `json:"city"`
	PostCode PostalCode `json:"postalCode"`
}

type PostalCode string

type PersonXML struct {
	FirstName    string     `xml:"firstName"`
	LastName     string     `xml:"lastName"`
	Home         AddressXML `xml:"address"`
	PhoneNumbers []string   `xml:"phoneNumber"`
}

type AddressXML struct {
	Street   string     `xml:"streetAddress"`
	City     string     `xml:"city"`
	PostCode PostalCode `xml:"postalCode"`
}
