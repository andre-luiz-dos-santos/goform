package validate

import (
	"regexp"
)

type Regexp struct {
	RE      *regexp.Regexp
	Message string
}

func (r Regexp) Validate(f Form, path []interface{}) {
	s, ok := f.GetValue(path...).(string)
	if !ok {
		return
	}
	if r.RE.MatchString(s) {
		return
	}
	f.SetError(r.Message, path...)
}
