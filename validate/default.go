package validate

type Default struct {
	Value interface{}
}

func (r Default) Validate(f Form, path []interface{}) {
	v := f.GetValue(path...)
	if v != nil {
		return
	}
	f.SetValue(r.Value, path...)
}
