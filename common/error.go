package common

import "strings"

type ErrorE struct {
	errInfo string
}

func NewError(err string, more ...string) error {
	var sb strings.Builder
	sb.WriteString(err)
	for _, opt := range more {
		sb.WriteString(opt)
	}
	return &ErrorE{sb.String()}
}

func (e ErrorE) Error() string {
	return e.errInfo
}
