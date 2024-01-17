package service

import "fmt"

type FieldValidationError struct {
	msg string
}

func (e *FieldValidationError) Error() string {
	return fmt.Sprintf("field validation error: %s", e.msg)
}

func NewFieldValidationError(msg string) error {
	return &FieldValidationError{msg: msg}
}
