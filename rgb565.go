// Copyright 2020 Lars Hoogestraat
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rgb565

import (
	"image"
	"image/color"
	"math"
)

// RGB565 is an in-memory image whose At method returns RGB565 values.
type RGB565 struct {
	// Pix holds the image's pixels, as RGB565 values in big-endian format. The pixel at
	// (x, y) starts at Pix[(y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

// Model is the model for RGB565 colors.
var Model = color.ModelFunc(rgb565Model)

// Color represents a RGB565 color.
type Color struct {
	RGB565 uint16
}

// NewRGB565 returns a new RGB565 image with the given bounds.
func NewRGB565(r image.Rectangle) *RGB565 {
	return &RGB565{
		Pix:    make([]uint8, 2*r.Dx()*r.Dy()),
		Stride: 2 * r.Dx(),
		Rect:   r,
	}
}

func (p *RGB565) Bounds() image.Rectangle {
	return p.Rect
}

func (p *RGB565) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return Color{0}
	}

	i := p.PixOffset(x, y)

	return Color{uint16(p.Pix[i])<<8 | uint16(p.Pix[i+1])}
}

func (p *RGB565) PixOffset(x, y int) int {
	return 2*(x-p.Rect.Min.X) + (y-p.Rect.Min.Y)*p.Stride
}

func (p *RGB565) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)

	c1 := Model.Convert(c).(Color)
	p.Pix[i+0] = uint8(c1.RGB565 >> 8)
	p.Pix[i+1] = uint8(c1.RGB565)
}

func (p *RGB565) ColorModel() color.Model {
	return Model
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.RGB565&0b1111100000000000) >> 11
	g = uint32(c.RGB565&0b0000011111100000) >> 5
	b = uint32(c.RGB565&0b0000000000011111) >> 0

	r = uint32(math.Round(float64(r)*255/31)) << 8
	g = uint32(math.Round(float64(g)*255/63)) << 8
	b = uint32(math.Round(float64(b)*255/31)) << 8

	a = 0xffff

	return
}

func rgb565Model(c color.Color) color.Color {
	if _, ok := c.(Color); ok {
		return c
	}

	r, g, b, _ := c.RGBA()

	return Color{uint16((r<<8)&0b1111100000000000 | (g<<3)&0b0000011111100000 | (b>>3)&0b0000000000011111)}
}
