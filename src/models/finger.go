package models

type FingerPrint struct {
	Image []byte `json:"image"`
}

type Finger struct {
	Name  string `json:"name,omitempty"`
	Nfiq  int    `json:"nfiq"`
	Image []byte `json:"image,omitempty"`
}

type Fingers struct {
	Fingers []Finger `json:"fingers,omitempty"`
}
