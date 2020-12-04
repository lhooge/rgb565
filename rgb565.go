// Copyright 2020 Lars Hoogestraat
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rgb565

import (
	"encoding/binary"
	"image"
	"image/color"
	"math"
)

// RGBA is an in-memory image whose At method returns color.RGBA values.
type RGB565 struct {
	// Pix holds the image's pixels, as RGB565 values. The pixel at
	// (x, y) starts at Pix[(y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

//Model is the Model for RGB565 colors.
var Model = color.ModelFunc(rgb565Model)

//Color represents a RGB565 color.
type Color uint16

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
		return Color(0)
	}

	i := p.PixOffset(x, y)

	return Color(binary.LittleEndian.Uint16(p.Pix[i : i+2]))
}

func (p *RGB565) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2
}

func (p *RGB565) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	binary.LittleEndian.PutUint16(p.Pix[i:i+2], uint16(Model.Convert(c).(Color)))
}

func (p *RGB565) ColorModel() color.Model {
	return Model
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c&0b1111100000000000) >> 11
	g = uint32(c&0b0000011111100000) >> 5
	b = uint32(c&0b0000000000011111) >> 0

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

	return Color((r<<8)&0b1111100000000000 | (g<<3)&0b0000011111100000 | (b>>3)&0b0000000000011111)
}
