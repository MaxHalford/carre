package carre

import "image"

func BreakAndPaint(in *image.NRGBA, out *image.NRGBA, threshold uint8) {
	var (
		bounds = in.Bounds()
		avg    = getAvgColor(in)
		diff   = getAvgDiff(in, avg)
	)
	if (diff < threshold) || (bounds.Dx() < 2) || (bounds.Dy() < 2) {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				out.SetNRGBA(x, y, avg)
			}
		}
	} else {
		var (
			// Center coordinates
			cx = (bounds.Min.X + bounds.Max.X) / 2
			cy = (bounds.Min.Y + bounds.Max.Y) / 2
			// Upper left corner
			ul = image.Rectangle{
				Min: bounds.Min,
				Max: image.Point{cx, cy},
			}
			// Upper right corner
			ur = image.Rectangle{
				Min: image.Point{bounds.Min.X, cy},
				Max: image.Point{cx, bounds.Max.Y},
			}
			// Lower right corner
			lr = image.Rectangle{
				Min: image.Point{cx, cy},
				Max: bounds.Max,
			}
			// Lower left corner
			ll = image.Rectangle{
				Min: image.Point{cx, bounds.Min.Y},
				Max: image.Point{bounds.Max.X, cy},
			}
		)
		for _, subRect := range []image.Rectangle{ul, ur, lr, ll} {
			var subNRGBA = ImageToNRGBA(in.SubImage(subRect))
			BreakAndPaint(subNRGBA, out, threshold)
		}
	}
}
