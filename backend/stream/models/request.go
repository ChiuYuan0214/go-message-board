package models

type TextRequest struct {
	DataType int    `json:"dataType"`
	Data     string `json:"data"`
}
