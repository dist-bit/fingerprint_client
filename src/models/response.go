package models

type Response struct {
	Status  bool        `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
}
