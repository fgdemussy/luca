package data

import (
	"encoding/json"
	"io"
)

// ToJSON writes the JSON representation of v to the stream
func ToJSON(v interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(v)
}

// FromJSON writes decoded JSON to v from io.Reader r
func FromJSON(v interface{}, r io.Reader) error {
	d := json.NewDecoder(r)

	return d.Decode(v)
}
