package errors

import "fmt"

type ErrValidation struct {
	Property    string
	Given       *string
	Description string
}

func (e *ErrValidation) Error() string {
	// e.g.
	// the property "Name" given "" cannot be empty.
	// the property "Email" given "given-email" is invalid email address.
	if e.Given == nil {
		return fmt.Sprintf("the property \"%s\" given 'nil' %s.", e.Property, e.Description)
	}
	return fmt.Sprintf("the property \"%s\" given \"%s\" %s.", e.Property, *e.Given, e.Description)
}

func (e *ErrValidation) errorType() string {
	return "validation"
}

func (e *ErrValidation) As(t interface{}) bool {
	if x, ok := t.(interface{ errorType() string }); ok && x.errorType() == e.errorType() {
		return true
	}
	return false
}

type ErrNotFound struct {
	Object string
	Id     string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("\"%s\" not found (id: \"%s\")", e.Object, e.Id)
}

func (e *ErrNotFound) errorType() string {
	return "not found"
}

func (e *ErrNotFound) As(t interface{}) bool {
	if x, ok := t.(interface{ errorType() string }); ok && x.errorType() == e.errorType() {
		return true
	}
	return false
}
