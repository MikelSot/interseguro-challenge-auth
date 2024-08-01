package model

type IssueType string

const (
	IssueMissingRequiredParam     IssueType = "MISSING_REQUIRED_PARAMETER"
	IssueMissingRequiredHeader    IssueType = "MISSING_REQUIRED_HEADER"
	IssueMissingRequiredFieldBody IssueType = "MISSING_REQUIRED_FIELD_BODY"
	IssueInvalidSyntaxParam       IssueType = "INVALID_PARAMETER_SYNTAX"
	IssueBodyError                IssueType = "INVALID_BODY"
	IssueViolatedValidation       IssueType = "VIOLATED_VALIDATION"
	IssueViolatedMaxLimit         IssueType = "VIOLATED_MAX_LIMIT"
	IssueResourceNotFound         IssueType = "RESOURCE_NOT_FOUND"
)
