package models

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Status string

const (
	OK     Status = "OK"
	ERROR  Status = "ERROR"
	DENIED Status = "DENIED"
)
