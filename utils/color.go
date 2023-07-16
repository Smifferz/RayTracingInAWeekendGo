package utils

import (
	"fmt"
	"math"
)

func WriteColor(pixelColor Color) {
	fmt.Printf("%+v %+v %+v\n", int32(255.999*pixelColor.X()),
		int32(255.999*pixelColor.Y()),
		int32(255.999*pixelColor.Z()))
}

func WriteMultiSampleColor(pixelColor Color, samplesPerPixel int) {
	r := pixelColor.X()
	g := pixelColor.Y()
	b := pixelColor.Z()

	// Divide the color by the number of samples
	scale := 1.0 / float64(samplesPerPixel)
	r = math.Sqrt(scale * r)
	g = math.Sqrt(scale * g)
	b = math.Sqrt(scale * b)

	// Write the translated [0,255] value of each color component
	fmt.Printf("%+v %+v %+v\n", int32(256*Clamp(r, 0.0, 0.999)), int32(256*Clamp(g, 0.0, 0.999)), int32(256*Clamp(b, 0.0, 0.999)))
}
