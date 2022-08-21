package trainer

/*
#cgo CFLAGS: -I../../headers
#cgo LDFLAGS: -L../../libs -lfingerprint
#cgo linux LDFLAGS: -L../../libs -lfingerprint -Wl,-rpath=./libs

#include <stdlib.h>
#include <fingerprint.h>
*/
import "C"

import (
	"bytes"
	"errors"
	"fmt"
	"hackaton/src/models"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"unsafe"

	"github.com/disintegration/imaging"
)

type FingerDetections struct {
	detections []int
}

type FingerprintDetector struct {
}

func (net *FingerprintDetector) Init() {
	C.initFingersModel()
}

var fingersName = []string{"index", "middle", "ring", "little"}

func (net *FingerprintDetector) ProcessFingerprint(file []byte) (*models.Finger, error) {

	status := C.extractFingerPrints((*C.uchar)(&file[0]), C.int(len(file)))

	if !bool(status) {
		return nil, errors.New("cant extract fingerprints")
	}

	fingerPath := "wsq/data.jpeg"

	saved, err := ioutil.ReadFile(fingerPath)
	if err != nil {
		return nil, errors.New("cant decode file")
	}

	err = os.Remove(fingerPath)
	if err != nil {
		return nil, err
	}

	nfiq := net.GetNfiqScore(saved)

	finger := models.Finger{
		Image: saved,
		Name:  "finger",
		Nfiq:  nfiq,
	}

	return &finger, nil
}

func (net *FingerprintDetector) GetFingerprints(file []byte, hand int) (*models.Fingers, error) {

	tmp := make([]byte, len(file))
	copy(tmp, file)

	inference := C.detectFingerprints((*C.uchar)(&file[0]), C.int(len(file)))
	results := C.FingerDetections(*inference)

	defer C.free(unsafe.Pointer(inference))

	var fingers []models.Finger

	img, err := imaging.Decode(bytes.NewReader(file), imaging.AutoOrientation(true))

	if err != nil {
		return nil, err
	}

	for i, rect := range results.detections {

		if err != nil {
			return nil, errors.New("image with bad quality")
		}

		x0 := int(rect[0])
		y0 := int(rect[1])
		x1 := int(rect[2]) + int(rect[0])
		y1 := int(rect[3]) + int(rect[1])

		crop := imaging.Crop(img,
			image.Rect(
				x0,
				y0,
				x1,
				y1,
			),
		)

		var dst *image.NRGBA
		if int(rect[2]) > int(rect[3]) {
			if hand == 0 {
				// left
				dst = imaging.Rotate(crop, -270, color.NRGBA{0, 0, 0, 0})
			} else {
				// right
				dst = imaging.Rotate(crop, -90, color.NRGBA{0, 0, 0, 0})
				//dst = imaging.Rotate90(crop)
			}
		} else {
			if hand == 0 {
				// left
				dst = imaging.Rotate(crop, -270, color.NRGBA{0, 0, 0, 0})
			} else {
				// right
				dst = imaging.Rotate(crop, -90, color.NRGBA{0, 0, 0, 0})
				//dst = imaging.Rotate90(crop)
			}
		}

		// Convert to array of bytes
		buf := new(bytes.Buffer)

		err = imaging.Encode(buf, dst, imaging.JPEG, imaging.JPEGQuality(100))
		if err != nil {
			return nil, err
		}

		if err != nil {
			fmt.Println(err)
			return nil, errors.New("image with bad quality")
		}

		out := buf.Bytes()

		status := C.extractFingerPrints((*C.uchar)(&out[0]), C.int(len(out)))

		if !bool(status) {
			return nil, errors.New("cant extract fingerprints")
		}

		fingerPath := "wsq/data.jpeg"

		saved, err := ioutil.ReadFile(fingerPath)
		if err != nil {
			return nil, errors.New("cant decode file")
		}

		err = os.Remove(fingerPath)
		if err != nil {
			return nil, err
		}

		nfiq := net.GetNfiqScore(saved)

		finger := models.Finger{
			Image: saved,
			Name:  fingersName[i],
			Nfiq:  nfiq,
		}

		fingers = append(fingers, finger)

	}

	return &models.Fingers{
		Fingers: fingers,
	}, nil
}

func (net *FingerprintDetector) GetNfiqScore(file []byte) int {
	score := C.getNFiq((*C.uchar)(&file[0]), C.int(len(file)))
	return int(score)
}

func (net *FingerprintDetector) GetISOFIle() ([]byte, error) {
	isoPath := "fingerprint.iso"
	iso, err := ioutil.ReadFile(isoPath)
	if err != nil {
		return nil, errors.New("cant decode file")
	}

	os.Remove(isoPath)

	return iso, nil
}

func (net *FingerprintDetector) GenerateWSQ(file []byte) ([]byte, error) {
	C.generateWSQ((*C.uchar)(&file[0]), C.int(len(file)))

	webpImage := "wsq/fingerprint.jpg"
	//jpgImage := "wsq/fingerprint.jpg"
	wsqPath := "wsq/fingerprint.wsq"

	wsq, err := ioutil.ReadFile(wsqPath)
	if err != nil {
		return nil, errors.New("cant decode file")
	}

	os.Remove(wsqPath)
	os.Remove(webpImage)
	//os.Remove(jpgImage)
	return wsq, nil
}

func (net *FingerprintDetector) SaveFileFingerprint(report string, img []byte) error {
	path := fmt.Sprintf("files/fingerprint_%s.jpg", report)
	if err := ioutil.WriteFile(path, img, 0644); err != nil {
		return err
	}

	return nil
}
