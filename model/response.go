package model

type Response struct {
	Code    StatusCode `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
}

type Responses []Response

type MessageResponse struct {
	Data     interface{} `json:"data,omitempty"`
	Errors   Responses   `json:"errors,omitempty"`
	Messages Responses   `json:"messages,omitempty"`
}
