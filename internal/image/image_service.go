package image

import (
	"bytes"

	"github.com/h2non/bimg"
)

func ResizeImage(buf *bytes.Buffer) ([]byte, error) {
	// TODO: check that the image format is appropriate for the application
	resized, err := bimg.NewImage(buf.Bytes()).Process(bimg.Options{
		Width:  500,
		Height: 500,
	})
	if err != nil {
		return nil, err
	}

	return resized, nil
}
