package validate

type IsBoolean struct {
}

func (r IsBoolean) Validate(f Form, path []interface{}) {
	_, ok := f.GetValue(path...).(bool)
	if !ok {
		f.SetError("Não é boolean", path...)
	}
}
