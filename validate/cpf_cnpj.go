package validate

type CPFCNPJ struct {
}

func (r CPFCNPJ) Validate(f Form, path []interface{}) {
	s, ok := f.GetValue(path...).(string)
	if !ok {
		return
	}
	if !isCPF(s) && !isCNPJ(s) {
		f.SetError("CPF/CNPJ é inválido", path...)
	}
}
