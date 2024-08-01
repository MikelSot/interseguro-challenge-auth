package model

import (
	"fmt"
)

// StatusCode type errors of the API
type StatusCode string

// Error records an error, the http code and the message that caused it.
type Error struct {
	code       StatusCode
	err        error
	where      string
	who        string
	endpoint   string
	statusHTTP int
	data       interface{}
	apiMessage string
	Fields     ErrorDetails
}

// NewError returns a new pointer Error
func NewError() *Error {
	return &Error{}
}

// Error implements the interface error
func (e *Error) Error() string {
	return fmt.Sprintf("Status: %d:%s | Endpoint: %s | Where: %s | Who: %s | Message: %s  | Err: %v | Fields: %v",
		e.statusHTTP, e.code, e.endpoint, e.where, e.who, e.apiMessage, e.err, e.Fields)
}

// SetError sets the error field
func (e *Error) SetError(err error) { e.err = err }

// Code gets the code field
func (e *Error) Code() StatusCode { return e.code }

// SetCode sets the code field
func (e *Error) SetCode(code StatusCode) { e.code = code }

// HasCode returns true if the struct has the code field
func (e *Error) HasCode() bool { return e.code != "" }

// Where gets the where field
func (e *Error) Where() string { return e.where }

// Endpoint gets the endpoint field
func (e *Error) Endpoint() string { return e.endpoint }

// SetEndpoint sets the endpoint field
func (e *Error) SetEndpoint(endpoint string) { e.endpoint = endpoint }

// SetWhere sets the where field
func (e *Error) SetWhere(where string) { e.where = where }

// Who gets who field
func (e *Error) Who() string { return e.who }

// SetWho sets who field
func (e *Error) SetWho(who string) { e.who = who }

// APIMessage gets the api message field
func (e *Error) APIMessage() string { return e.apiMessage }

// SetAPIMessage sets the api message field
func (e *Error) SetAPIMessage(message string) { e.apiMessage = message }

// SetErrorAsAPIMessage sets the error fields as the api message
func (e *Error) SetErrorAsAPIMessage() { e.apiMessage = fmt.Sprintf("%v", e.err) }

// HasAPIMessage returns true if the struct has the api message field
func (e *Error) HasAPIMessage() bool { return e.apiMessage != "" }

// StatusHTTP gets the status http field
func (e *Error) StatusHTTP() int { return e.statusHTTP }

// SetStatusHTTP sets the status http field
func (e *Error) SetStatusHTTP(status int) { e.statusHTTP = status }

// HasStatusHTTP returns true if the struct has the status http field
func (e *Error) HasStatusHTTP() bool { return e.statusHTTP != 0 }

// Data gets the data field
func (e *Error) Data() interface{} { return e.data }

// SetData sets the data field
func (e *Error) SetData(data interface{}) { e.data = data }

// HasData returns true if the struct has the data field
func (e *Error) HasData() bool { return e.data != nil }
