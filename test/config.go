package test

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/M-Quadra/kazaana/v2"
)

var imgAria image.Image

func init() {
	f, err := os.Open("aria.jpeg")
	if kazaana.HasError(err) {
		os.Exit(-1)
	}
	defer f.Close()

	var (
		_ = jpeg.DefaultQuality
		_ = png.DefaultCompression
	)

	img, _, err := image.Decode(f)
	if kazaana.HasError(err) {
		os.Exit(-1)
	}
	imgAria = img
}
