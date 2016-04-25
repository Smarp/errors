package errors

import (
	"fmt"
	"strings"
)

type Errors []*Error

func (e Errors) Error() string {
	if len(e) == 0 {
		return "no error description"
	}

	ret := make([]string, 0)
	for _, err := range e {
		ret = append(ret, err.Error())
	}
	return strings.Join(ret, ";")
}

func (e Errors) Type() interface{} {
	if len(e) == 0 {
		return nil
	}

	switch ee := e[0].Err.(type) {
	case stringerStrct:
		return ee.e
	case stringer:
		return ee
	default:
		return nil
	}
}

func (e Errors) Add(ee error) Errors {
	return append(e, Wrap(ee)...)
}

func New(e stringer, meta map[string]interface{}) Errors {
	if e == nil {
		return nil
	}

	ee := Errors{{Err: e, Meta: meta}}
	return ee
}

func Wrap(e error) Errors {
	var err error

	switch e := e.(type) {
	case Errors:
		return e
	case error:
		err = e
	default:
		err = fmt.Errorf("%v", e)
	}

	errStringer := stringerStrct{
		e: err,
	}

	return New(errStringer, nil)
}
