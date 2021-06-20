package dither

import (
	"image"
	"image/color"
)

func threshold(pixel color.Gray, th uint8) color.Gray {
	white := color.Gray{Y: 255}
	black := color.Gray{Y: 0}

	if pixel.Y < th {
		return white
	} else {
		return black
	}
}

func Threshold(img *image.Gray, th uint8) *image.Gray {
	bounds := img.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()
	thresholdImg := image.NewGray(bounds)

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			thresholdImg.Set(x, y, threshold(img.GrayAt(x, y), th))
		}
	}

	return thresholdImg
}

func Grayscale(img image.Image) *image.Gray {
	bounds := img.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()
	grayscale := image.NewGray(bounds)

	for x := bounds.Min.X; x < dx; x++ {
		for y := bounds.Min.Y; y < dy; y++ {
			grayscale.Set(x, y, img.At(x, y))
		}
	}

	return grayscale
}
