package services

import (
	"hackaton/src/interfaces"
	"hackaton/src/models"
)

type FingerprintDetectorService struct {
	interfaces.IFingerprintDetector
}

func (service *FingerprintDetectorService) GetNfiqScoreService(file []byte) int {
	return service.GetNfiqScore(file)
}

func (service *FingerprintDetectorService) GenerateWSQService(file []byte) ([]byte, error) {
	return service.GenerateWSQ(file)
}

func (service *FingerprintDetectorService) GetFingerprintsService(file []byte, hand int) (*models.Fingers, error) {
	return service.GetFingerprints(file, hand)
}

func (service *FingerprintDetectorService) ProcessFingerprintService(file []byte) (*models.Finger, error) {
	return service.ProcessFingerprint(file)
}

func (service *FingerprintDetectorService) GetISOFIleService() ([]byte, error) {
	return service.GetISOFIle()
}
