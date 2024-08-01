package model

const (
	BindFailed StatusCode = "bind_failed"
	// Failure sends the custom error message and API message from the logic
	Failure         StatusCode = "failure"
	Ok              StatusCode = "ok"
	RecordCreated   StatusCode = "record_created"
	RecordNotFound  StatusCode = "record_not_found"
	InvalidEmail    StatusCode = "invalid_email"
	InvalidPassword StatusCode = "invalid_password"
	// UnexpectedError is a server error
	UnexpectedError StatusCode = "unexpected_error"
	// AuthError is any of authorization errors
	AuthError StatusCode = "authorization_error"
)
