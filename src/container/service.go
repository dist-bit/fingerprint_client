package container

import (
	"hackaton/src/controllers"
	"hackaton/src/response"
	"hackaton/src/services"
	"hackaton/src/trainer"
	"sync"
)

type kernel struct{}

type IServiceContainer interface {
	InjectTrainerController() controllers.TrainerController
}

func (k *kernel) InjectTrainerController() controllers.TrainerController {
	detector := &trainer.FingerprintDetector{}

	detector.Init()

	return controllers.TrainerController{
		IFingerprintDetectorService: &services.FingerprintDetectorService{
			IFingerprintDetector: detector,
		},
		IResponse: &response.Response{},
	}
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
