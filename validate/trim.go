package validate

import (
	"strings"
)

type Trim struct {
	Inside bool
}

func (r Trim) Validate(f Form, path []interface{}) {
	s, ok := f.GetValue(path...).(string)
	if !ok {
		return
	}
	if r.Inside {
		s = strings.Join(strings.Fields(s), " ")
	} else {
		s = strings.TrimSpace(s)
	}
	f.SetValue(s, path...)
}
