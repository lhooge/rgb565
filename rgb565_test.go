package rgb565_test

import (
	"image/color"
	"testing"

	"git.hoogi.eu/snafu/rgb565"
)

var rgb565TestCases = []struct {
	rgb    color.RGBA
	rgb565 rgb565.Color
}{
	{rgb: color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}, rgb565: 0x0000},
	{rgb: color.RGBA{R: 0x84, G: 0x00, B: 0x00, A: 0xff}, rgb565: 0x8000},
	{rgb: color.RGBA{R: 0x00, G: 0x82, B: 0x00, A: 0xff}, rgb565: 0x0400},
	{rgb: color.RGBA{R: 0x00, G: 0x00, B: 0x84, A: 0xff}, rgb565: 0x0010},
	{rgb: color.RGBA{R: 0x00, G: 0x2d, B: 0x84, A: 0xff}, rgb565: 0x170},
	{rgb: color.RGBA{R: 0x84, G: 0x2d, B: 0x84, A: 0xff}, rgb565: 0x8170},
	{rgb: color.RGBA{R: 0x84, G: 0xff, B: 0xff, A: 0xff}, rgb565: 0x87ff},
	{rgb: color.RGBA{R: 0x84, G: 0x82, B: 0xff, A: 0xff}, rgb565: 0x841f},
	{rgb: color.RGBA{R: 0x84, G: 0x82, B: 0x84, A: 0xff}, rgb565: 0x8410},
	{rgb: color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, rgb565: 0xFFFF},
}

func TestRGBToRGB565(t *testing.T) {
	for _, test := range rgb565TestCases {
		got := rgb565.Model.Convert(test.rgb)
		want := test.rgb565
		if got != want {
			t.Errorf("unexpected RGB565 value for %+v: got: %016b, want: %016b", test.rgb, got, want)
		}
	}
}

func TestRGB565ToRGBA(t *testing.T) {
	for _, test := range rgb565TestCases {
		got := color.RGBAModel.Convert(test.rgb565).(color.RGBA)
		want := test.rgb
		if got != want {
			t.Errorf("unexpected RGBA value for %x: got: %+v, want: %+v", test.rgb565, got, want)
		}
	}
}
