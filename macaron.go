package goform

import (
	"gopkg.in/macaron.v1"
	"net/http"
)

func NewMacaron(ctx *macaron.Context) (*Form, error) {
	return NewMacaronSize(ctx, 100_000)
}

func NewMacaronSize(ctx *macaron.Context, size int64) (*Form, error) {
	f := New()
	err := f.ReadMacaronSize(ctx, size)
	return f, err
}

func (f *Form) ReadMacaronSize(ctx *macaron.Context, size int64) error {
	r := http.MaxBytesReader(ctx.Resp, ctx.Req.Body().ReadCloser(), size)
	defer r.Close()
	return f.ReadJSON(r)
}

func (f *Form) WriteMacaron(ctx *macaron.Context) {
	body := map[string]interface{}{
		"values": f.Values,
		"errors": f.Errors,
	}
	if f.HasError() {
		body["success"] = false
		body["error"] = "Erro de validação"
		ctx.JSON(http.StatusNotAcceptable, body)
	} else {
		body["success"] = true
		ctx.JSON(http.StatusOK, body)
	}
}
