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
				BoundMin: image.Point{0, 0},
				BoundMax: image.Point{100, 100},
				At:       color.RGBA{101, 101, 255, 255},
			},
		},
		{
			img: Image{200, 50, 0},
			expectedResults: testCaseImages{
				BoundMin: image.Point{0, 0},
				BoundMax: image.Point{200, 50},
				At:       color.RGBA{1, 1, 255, 255},
			},
		},
	}

	for i, c := range cases {
		i, c := i, c
		t.Run(fmt.Sprintf("Test Case N%d", i), func(t *testing.T) {
			t.Parallel()

			img := c.img

			rectangle := img.Bounds()
			if rectangle.Min != c.expectedResults.BoundMin {
				t.Fatalf("Unexpected retangle Min value. \n Got: %v \n Expected: %v", rectangle.Min, c.expectedResults.BoundMin)
			}

			if rectangle.Max != c.expectedResults.BoundMax {
				t.Fatalf("Unexpected retangle Max value. \n Got: %v \n Expected: %v", rectangle.Max, c.expectedResults.BoundMax)
			}

			if img.At(1, 1) != c.expectedResults.At {
				t.Fatalf("Unexpected At value. \n Got: %v \n Expected: %v", img.At(1, 1), c.expectedResults.At)
			}
		})
	}

}
