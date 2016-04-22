package errors

import (
	"fmt"
	"strings"
)

type stringer interface {
	String() string
}

type stringerStrct struct {
	e error
}

func (s stringerStrct) String() string {
	return s.e.Error()
}

type Error struct {
	Meta map[string]interface{}
	Err  stringer
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	msg := e.Err.String()

	if len(e.Meta) != 0 {
		ret := make([]string, 0)
		for k, v := range e.Meta {
			ret = append(ret, fmt.Sprintf("%s: %s", k, v))
		}
		msg += "{" + strings.Join(ret, ";") + "}"
	}

	return msg
}
