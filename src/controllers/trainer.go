package controllers

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"hackaton/src/interfaces"
	"io"
	"strconv"
	"sync"
)

var lock sync.Mutex

const (
	MB = 1 << 20
)

type Sizer interface {
	Size() int64
}

type TrainerController struct {
	interfaces.IFingerprintDetectorService
	interfaces.IResponse
}

func (controller *TrainerController) GetFingerprintsController(c *fiber.Ctx) error {

	lock.Lock()
	defer lock.Unlock()

	hand := c.FormValue("hand")

	position, err := strconv.Atoi(hand)

	if err != nil {
		return controller.Error("invalid hand", c)
	}

	imageForm, err := c.FormFile("image")

	if err != nil {
		return err
	}

	file, err := imageForm.Open()

	if err != nil {
		return err
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return controller.Error(err.Error(), c)
	}

	fingers, err := controller.GetFingerprintsService(buf.Bytes(), position)

	if err != nil {
		return controller.Error(err.Error(), c)
	}
	return controller.Success(fingers, c)
}

func (controller *TrainerController) ProcessFingerprintController(c *fiber.Ctx) error {

	lock.Lock()
	defer lock.Unlock()

	imageForm, err := c.FormFile("image")

	if err != nil {
		return err
	}

	file, err := imageForm.Open()

	if err != nil {
		return err
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return controller.Error(err.Error(), c)
	}

	fingers, err := controller.ProcessFingerprintService(buf.Bytes())

	if err != nil {
		return controller.Error(err.Error(), c)
	}
	return controller.Success(fingers, c)
}

func (controller *TrainerController) GetWSQFingerprintController(c *fiber.Ctx) error {

	lock.Lock()
	defer lock.Unlock()

	imageForm, err := c.FormFile("front")

	if err != nil {
		return err
	}

	file, err := imageForm.Open()

	if err != nil {
		return err
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return controller.Success(err.Error(), c)
	}

	wsq, err := controller.GenerateWSQService(buf.Bytes())

	if err != nil {
		return controller.Success(err.Error(), c)
	}
	return controller.File(wsq, c)
}

func (controller *TrainerController) GetFingerprintNfiqScoreController(c *fiber.Ctx) error {

	lock.Lock()
	defer lock.Unlock()

	imageForm, err := c.FormFile("image")

	if err != nil {
		return err
	}

	file, err := imageForm.Open()

	if err != nil {
		return err
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return controller.Success(err.Error(), c)
	}

	score := controller.GetNfiqScoreService(buf.Bytes())

	if err != nil {
		return controller.Success(err.Error(), c)
	}

	return controller.Success(score, c)
}

func (controller *TrainerController) GetISOFIle(c *fiber.Ctx) error {
	file, err := controller.GetISOFIleService()

	if err != nil {
		return controller.Success(err.Error(), c)
	}

	return controller.File(file, c)
}
