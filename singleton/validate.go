package singleton

import "github.com/go-playground/validator"

var validate *Validator

// Validator custom
type Validator struct {
	validator *validator.Validate
}

// GetValidator Validator
func GetValidator() *Validator {
	if validate == nil {
		validate = &Validator{validator.New()}
	}
	return validate
}

// Validate struct
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// ValidateField var
func (v *Validator) ValidateField(field interface{}, tag string) error {
	return v.validator.Var(field, tag)
}
