package manyerrors


import (
	"bytes"
)

type Errors interface {
	Error() string
	Errors() []error
}


type internalErrors struct {
	errors []error
}


func New(errors ...error) Errors {
	return &internalErrors{
		errors:errors,
	}
}


func (errors *internalErrors) Errors() []error {
	return errors.errors
}

func (errors *internalErrors) Error() string {
	var buffer bytes.Buffer

	for _, err := range errors.errors {
		buffer.WriteString(err.Error())
		buffer.WriteRune('\n')
	}

	return buffer.String()
}
