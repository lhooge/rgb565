package rgb565_test

import (
	"image/color"
	"testing"

	"git.hoogi.eu/snafu/rgb565"
)

var rgb565TestCases = []struct {
	rgbColor    color.RGBA
	rgb565Color rgb565.Color
}{
	{rgbColor: color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}, rgb565Color: rgb565.Color{RGB565: 0x0000}},
	{rgbColor: color.RGBA{R: 0x84, G: 0x00, B: 0x00, A: 0xff}, rgb565Color: rgb565.Color{RGB565: 0x8000}},
	{rgbColor: color.RGBA{R: 0x00, G: 0x82, B: 0x00, A: 0xff}, rgb565Color: rgb565.Color{RGB565: 0x0400}},
	{rgbColor: color.RGBA{R: 0x00, G: 0x00, B: 0x84, A: 0xff}, rgb565Color: rgb565.Color{RGB565: 0x0010}},
	{rgbColor: color.RGBA{R: 0x00, G: 0x2d, B: 0x84, A: 0xff}, rgb565Color: rgb565.Color{RGB565: 0x170}},
	{rgbColor: color.RGBA{R: 0x84, G: 0x2d, B: 0x84, A: 0xff}, rgb565Color: rgb565.Color{RGB565: 0x8170}},
	{rgbColor: color.RGBA{R: 0x84, G: 0xff, B: 0xff, A: 0xff}, rgb565Color: rgb565.Color{RGB565: 0x87ff}},
	{rgbColor: color.RGBA{R: 0x84, G: 0x82, B: 0xff, A: 0xff}, rgb565Color: rgb565.Color{RGB565: 0x841f}},
	{rgbColor: color.RGBA{R: 0x84, G: 0x82, B: 0x84, A: 0xff}, rgb565Color: rgb565.Color{RGB565: 0x8410}},
	{rgbColor: color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, rgb565Color: rgb565.Color{RGB565: 0xFFFF}},
}

func TestRGBToRGB565(t *testing.T) {
	for _, test := range rgb565TestCases {
		got := rgb565.Model.Convert(test.rgbColor)
		want := test.rgb565Color
		if got != want {
			t.Errorf("unexpected RGB565 value for %+v: got: %016b, want: %016b", test.rgbColor, got, want)
		}
	}
}

func TestRGB565ToRGBA(t *testing.T) {
	for _, test := range rgb565TestCases {
		got := color.RGBAModel.Convert(test.rgb565Color).(color.RGBA)
		want := test.rgbColor
		if got != want {
			t.Errorf("unexpected RGBA value for %x: got: %+v, want: %+v", test.rgb565Color, got, want)
		}
	}
}
