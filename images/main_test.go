package main

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func Test_Image(t *testing.T) {
	type testCaseImages struct {
		BoundMin   image.Point
		BoundMax   image.Point
		ColorModel color.RGBA
		At         color.RGBA
	}
	cases := []struct {
		img             image.Image
		expectedResults testCaseImages
	}{
		{
			img: Image{100, 100, 100},
			expectedResults: testCaseImages{
				BoundMin:   image.Point{0, 0},
				BoundMax:   image.Point{100, 100},
				ColorModel: color.RGBA{101, 101, 255, 255},
				At:         color.RGBA{101, 101, 255, 255},
			},
		},
	}

	for i, c := range cases {
		i, c := i, c
		t.Run(fmt.Sprintf("Test Case N%d", i), func(t *testing.T) {
			img := c.img

			rectangle := img.Bounds()
			if rectangle.Min != c.expectedResults.BoundMin {
				t.Fatalf("Unexpected retangle Min value. \n Got: %v \n Expected: %v", rectangle.Min, c.expectedResults.BoundMin)
			}

			if rectangle.Max != c.expectedResults.BoundMax {
				t.Fatalf("Unexpected retangle Max value. \n Got: %v \n Expected: %v", rectangle.Max, c.expectedResults.BoundMax)
			}
		})
	}
}
