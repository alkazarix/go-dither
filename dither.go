package dither

import (
	"image"
	"image/color"
)

func Monochrome(original image.Image, filter Filter, errorMultiplier float32) image.Image {

	bounds := original.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()

	ydim := len(filter.Matrix) - 1
	xdim := len(filter.Matrix[0]) / 2

	img := image.NewGray(bounds)

	for x := bounds.Min.X; x < dx; x++ {
		for y := bounds.Min.Y; y < dy; y++ {
			pixel := original.At(x, y)
			img.Set(x, y, pixel)
		}
	}

	errors := newErrorMatrix(dx, dy)
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			px := float32(img.GrayAt(x, y).Y)
			px -= errors[x][y] * errorMultiplier
			q, px := quantifyError(px)

			img.SetGray(x, y, color.Gray{Y: uint8(px)})

			for xx := 0; xx <= ydim; xx++ {
				for yy := -xdim; yy <= xdim-1; yy++ {
					if x+xx >= dx || x+xx < 0 || y+yy >= dy || y+yy < 0 {
						continue
					}
					errors[x+xx][y+yy] += q * filter.Matrix[xx][yy+ydim]
				}
			}

		}
	}

	return img
}

func Color(original image.Image, filter Filter, errorMultiplier float32) image.Image {

	bounds := original.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()

	ydim := len(filter.Matrix) - 1
	xdim := len(filter.Matrix[0]) / 2

	img := image.NewRGBA(bounds)

	errorsR := newErrorMatrix(dx, dy)
	errorsG := newErrorMatrix(dx, dy)
	errorsB := newErrorMatrix(dx, dy)

	for x := bounds.Min.X; x < dx; x++ {
		for y := bounds.Min.Y; y < dy; y++ {
			pixel := original.At(x, y)
			img.Set(x, y, pixel)
		}
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {

			r, g, b, a := img.At(x, y).RGBA()
			pr := float32(uint8(r)) - errorsR[x][y]*errorMultiplier
			pb := float32(uint8(b)) - errorsB[x][y]*errorMultiplier
			pg := float32(uint8(g)) - errorsG[x][y]*errorMultiplier

			qr, pr := quantifyError(pr)
			qg, pg := quantifyError(pg)
			qb, pb := quantifyError(pb)

			img.SetRGBA(x, y, color.RGBA{uint8(pr), uint8(pg), uint8(pb), uint8(a)})

			for xx := 0; xx <= ydim; xx++ {
				for yy := -xdim; yy <= xdim-1; yy++ {
					if x+xx >= dx || x+xx < 0 || y+yy >= dy || y+yy < 0 {
						continue
					}
					errorsR[x+xx][y+yy] += qr * filter.Matrix[xx][yy+ydim]
					errorsB[x+xx][y+yy] += qb * filter.Matrix[xx][yy+ydim]
					errorsG[x+xx][y+yy] += qg * filter.Matrix[xx][yy+ydim]
				}
			}

		}
	}

	return img

}

func newErrorMatrix(dx, dy int) [][]float32 {
	errors := make([][]float32, dx)
	for x := 0; x < dx; x++ {
		errors[x] = make([]float32, dy)
		for y := 0; y < dy; y++ {
			errors[x][y] = 0
		}
	}
	return errors
}

func quantifyError(px float32) (float32, float32) {
	var q float32
	if px < 128 {
		q = -px
		px = 0
	} else {
		q = 255 - px
		px = 255
	}

	return q, px
}
