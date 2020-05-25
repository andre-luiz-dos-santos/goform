package validate

type Validator interface {
	Validate(f Form, path []interface{})
}

type Form interface {
	GetValue(path ...interface{}) interface{}
	SetValue(value interface{}, path ...interface{}) error
	SetError(error string, path ...interface{})
}
