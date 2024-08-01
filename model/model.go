package model

import (
	"errors"
	"reflect"
)

// Errors
var (
	ErrNilPointer = errors.New("el modelo recibido es nulo")
	ErrInvalidID  = errors.New("el id recibido no es valido")
)

// ValidateStructNil returns an error if the model is nil
func ValidateStructNil(i interface{}) error {
	// omit struct type
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		return nil
	}

	// Type: nil, Value: nil
	if i == nil {
		return ErrNilPointer
	}

	// Type: StructPointer, Value: nil
	if reflect.ValueOf(i).IsNil() {
		return ErrNilPointer
	}

	// Type: StructPointer, Value: ZeroValue
	return nil
}
