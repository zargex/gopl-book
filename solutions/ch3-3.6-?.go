// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

type newColor struct {
	r uint32
	g uint32
	b uint32
	a uint32
}

func (c newColor) RGBA() (r, g, b, a uint32) {
	return c.a, c.g, c.b, c.a
}

func averageColor(x, y, z, zz color.Color) color.Color {
	xR, xG, xB, xA := x.RGBA()
	yR, yG, yB, yA := y.RGBA()
	zR, zG, zB, zA := z.RGBA()
	zzR, zzG, zzB, zzA := zz.RGBA()
	r := (xR + yR + zR + zzR) / 4
	g := (xG + yG + zG + zzG) / 4
	b := (xB + yB + zB + zzB) / 4
	a := (xA + yA + zA + zzA) / 4
	return newColor{r, g, b, a}
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			subpixel1 := complex(x, y)
			subpixel2 := complex(float64(px+1)/width*(xmax-xmin)+xmin, y)
			subpixel3 := complex(x, float64(py+1)/height*(ymax-ymin)+ymin)
			subpixel4 := complex(float64(px+1)/width*(xmax-xmin)+xmin, float64(py+1)/height*(ymax-ymin)+ymin)

			subpixel1Color := mandelbrot(subpixel1)
			subpixel2Color := mandelbrot(subpixel2)
			subpixel3Color := mandelbrot(subpixel3)
			subpixel4Color := mandelbrot(subpixel4)

			zColor := averageColor(subpixel1Color, subpixel2Color, subpixel3Color, subpixel4Color)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, zColor)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
