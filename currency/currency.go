package currency

import (
	"encoding/xml"
	"fmt"
)

// Currency wraps ISO 4217 names and codes
// https://github.com/datasets/currency-codes
type Currency struct {
	Code       string `xml:"Ccy"`
	Country    string `xml:"CtryNm"`
	Name       string `xml:"CcyNm"`
	MinorUnits string `xml:"CcyMnrUnts"`
}

// Getall returns a list of available Currencies
func Getall() (*[]Currency, error) {
	var list struct {
		Currencies []Currency `xml:"CcyNtry"`
	}
	err := xml.Unmarshal([]byte(Data), &list)
	if err != nil {
		return nil, fmt.Errorf("Failed parsing ISO 4217 (currencies)")
	}

	return &list.Currencies, nil
}
