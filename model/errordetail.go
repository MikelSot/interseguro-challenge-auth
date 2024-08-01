package model

type ErrorDetail struct {
	Field       string    `json:"field,omitempty"`
	Issue       IssueType `json:"issue,omitempty"`
	Code        string    `json:"code,omitempty"`
	Description string    `json:"description"`
}

type ErrorDetails []ErrorDetail

func (e *ErrorDetails) Add(field ErrorDetail) {
	*e = append(*e, field)
}
