package carre

import (
	"image"
	"image/color"
)

// ImageToNRGBA converts an image.Image into an image.NRGBA.
func ImageToNRGBA(img image.Image) *image.NRGBA {
	var (
		bounds = img.Bounds()
		nrgba  = image.NewNRGBA(bounds)
	)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			nrgba.Set(x, y, img.At(x, y))
		}
	}
	return nrgba
}

func getAvgColor(img *image.NRGBA) color.NRGBA {
	var (
		bounds  = img.Bounds()
		r, g, b int
	)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			var c = img.NRGBAAt(x, y)
			r += int(c.R)
			g += int(c.G)
			b += int(c.B)
		}
	}
	return color.NRGBA{
		R: uint8(r / (bounds.Dx() * bounds.Dy())),
		G: uint8(g / (bounds.Dx() * bounds.Dy())),
		B: uint8(b / (bounds.Dx() * bounds.Dy())),
		A: 255,
	}
}

func getAvgDiff(img *image.NRGBA, color color.NRGBA) uint8 {
	var (
		bounds = img.Bounds()
		diff   int
	)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			var c = img.NRGBAAt(x, y)
			diff += int(absUint8(c.R, color.R))
			diff += int(absUint8(c.G, color.G))
			diff += int(absUint8(c.B, color.B))
		}
	}
	return uint8(diff / (3 * bounds.Dx() * bounds.Dy()))
}
