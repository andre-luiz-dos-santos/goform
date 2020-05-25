package goform

import (
	"github.com/andre-luiz-dos-santos/goform/validate"
	"reflect"
)

type Runner struct {
	*Form
	Paths [][]interface{}
}

func (f *Form) NewRunner() *Runner {
	return &Runner{f, [][]interface{}{{}}}
}

func (run *Runner) New(paths [][]interface{}) *Runner {
	return &Runner{run.Form, paths}
}

func (f *Form) In(path ...interface{}) *Runner {
	return f.NewRunner().In(path...)
}

func (run *Runner) In(path ...interface{}) *Runner {
	var ps [][]interface{}
	for _, p := range run.Paths {
		ps = append(ps, append(append(intf, p...), path...))
	}
	return run.New(ps)
}

func (f *Form) For(keys ...string) *Runner {
	return f.NewRunner().For(keys...)
}

func (run *Runner) For(keys ...string) *Runner {
	var ps [][]interface{}
	for _, p := range run.Paths {
		for _, key := range keys {
			ps = append(ps, append(append(intf, p...), key))
		}
	}
	return run.New(ps)
}

func (f *Form) Range() *Runner {
	return f.NewRunner().Range()
}

func (run *Runner) Range() *Runner {
	var ps [][]interface{}
	for _, p := range run.Paths {
		v := reflect.ValueOf(run.GetValue(p...))
		for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.IsValid() && v.Kind() == reflect.Slice {
			for i := 0; i < v.Len(); i++ {
				ps = append(ps, append(append(intf, p...), i))
			}
		}
	}
	return run.New(ps)
}

func (run *Runner) Run(validators ...validate.Validator) *Runner {
	for _, p := range run.Paths {
		for _, validator := range validators {
			validator.Validate(run.Form, p)
		}
	}
	return run
}
