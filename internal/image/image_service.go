package image

import (
	"bytes"

	"github.com/h2non/bimg"
)

func ResizeImage(buf *bytes.Buffer, width int, height int) ([]byte, error) {
	imageToResize := bimg.NewImage(buf.Bytes())
	imageType := imageToResize.Type()
	if !IsImageTypeAllowed(imageType) {
		return nil, ImageTypeNotAllowedError
	}

	resized, err := imageToResize.Process(bimg.Options{
		Width:  width,
		Height: height,
	})
	if err != nil {
		return nil, err
	}

	return resized, nil
}
