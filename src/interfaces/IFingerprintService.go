package interfaces

import "hackaton/src/models"

type IFingerprintDetectorService interface {
	GetNfiqScoreService(file []byte) int
	GenerateWSQService(file []byte) ([]byte, error)
	GetFingerprintsService(file []byte, hand int) (*models.Fingers, error)
	ProcessFingerprintService(file []byte) (*models.Finger, error)
	GetISOFIleService() ([]byte, error)
}
