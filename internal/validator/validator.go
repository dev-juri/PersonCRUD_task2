package validator

import "strings"

type Validator struct {
	Error string
}

func New() *Validator {
	return &Validator{Error: ""}
}

func (v *Validator) Valid() bool {
	return strings.TrimSpace(v.Error) == ""
}

func (v *Validator) AddError(message string) {
	v.Error = message
}
