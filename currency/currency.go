package currency

import (
	"encoding/xml"
	"fmt"
)

type Currency struct {
	Code    string `xml:"Ccy"`
	Country string `xml:"CtryNm"`
	Name    string `xml:"CcyNm"`
}

func get() (*[]Currency, error) {
	var list struct {
		Currencies []Currency `xml:"CcyNtry"`
	}
	err := xml.Unmarshal([]byte(Data), &list)
	if err != nil {
		return nil, fmt.Errorf("Failed parsing ISO 4217 (currencies)")
	}

	return &list.Currencies, nil
}
