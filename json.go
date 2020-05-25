package goform

import (
	"encoding/json"
	"io"
)

func (f *Form) ReadJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(&f.Values)
}
