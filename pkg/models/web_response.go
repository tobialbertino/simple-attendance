package models

type WebResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

type WebResponseError struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}
