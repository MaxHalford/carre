package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/MaxHalford/carre/carre"
)

var (
	inPath    = flag.String("in", "", "Input path")
	outPath   = flag.String("out", "", "Output path")
	threshold = flag.Int("threshold", 15, "Threshold")
)

func main() {
	// Parse the command-line arguments
	flag.Parse()

	// Open the input image
	var img, err = carre.LoadImage(*inPath)
	if err != nil {
		fmt.Printf("Failed to open image '%s'\n", *inPath)
		os.Exit(1)
	}

	// Convert it to NRGBA
	var nrgba = carre.ImageToNRGBA(img)

	// Apply the algorithm and save the output
	var output = image.NewNRGBA(nrgba.Bounds())
	carre.BreakAndPaint(nrgba, output, uint8(*threshold))
	carre.SaveImagePNG(output, *outPath)
}
