package model

const (
	BindFailed      StatusCode = "bind_failed"
	Failure         StatusCode = "failure"
	Ok              StatusCode = "ok"
	RecordCreated   StatusCode = "record_created"
	RecordNotFound  StatusCode = "record_not_found"
	InvalidEmail    StatusCode = "invalid_email"
	InvalidPassword StatusCode = "invalid_password"
	UnexpectedError StatusCode = "unexpected_error"
	AuthError       StatusCode = "authorization_error"

	TokenExpired StatusCode = "token_expired"
	Unauthorized StatusCode = "unauthorized"
)
