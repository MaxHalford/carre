package carre

import (
	"image"
	"image/png"
	"os"
)

// LoadImage reads and loads an image from a file path.
func LoadImage(path string) (image.Image, error) {
	infile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer infile.Close()
	img, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// SaveImagePNG saves an image to a PNG file.
func SaveImagePNG(img image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	png.Encode(f, img)
	return nil
}

func absUint8(a, b uint8) uint8 {
	if a > b {
		return a - b
	}
	return b - a
}
