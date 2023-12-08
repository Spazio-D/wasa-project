package utils

import (
	"image/jpeg"
	"io/fs"
	"os"

	"github.com/nfnt/resize"
)

func SaveAndResizeImage(path string, data []byte, umask fs.FileMode) error {
	if err := os.WriteFile(path, data, 0666); err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { err = file.Close() }()

	image, err := jpeg.Decode(file)
	if err != nil {
		return err
	}
	image = resize.Resize(1080, 1080, image, resize.NearestNeighbor)
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { err = out.Close() }()
	if err := jpeg.Encode(out, image, nil); err != nil {
		return err
	}

	return err
}
