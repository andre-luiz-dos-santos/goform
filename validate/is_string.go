package validate

type IsString struct {
}

func (r IsString) Validate(f Form, path []interface{}) {
	_, ok := f.GetValue(path...).(string)
	if !ok {
		f.SetError("Não é string", path...)
	}
}
