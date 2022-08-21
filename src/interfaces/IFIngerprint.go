package interfaces

import "hackaton/src/models"

type IFingerprintDetector interface {
	GetNfiqScore(file []byte) int
	GenerateWSQ(file []byte) ([]byte, error)
	GetFingerprints(file []byte, hand int) (*models.Fingers, error)
	ProcessFingerprint(file []byte) (*models.Finger, error)
	GetISOFIle() ([]byte, error)
}
