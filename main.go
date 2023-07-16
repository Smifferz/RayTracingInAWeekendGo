package main

import (
	"fmt"
	"go/raytracingweekend/utils"
	"math"
	"os"
)

func HitSphere(center utils.Point3, radius float64, r Ray) float64 {
	oc := utils.MinusVec3(r.Origin(), center)
	a := r.Direction().LengthSquared()
	halfB := utils.Dot(oc, r.Direction())
	c := oc.LengthSquared() - (radius * radius)
	discriminant := (halfB * halfB) - (a * c)
	if discriminant < 0 {
		return -1.0
	} else {
		return (-halfB - math.Sqrt(discriminant)) / a
	}
}

func RayColor(r *Ray, world Hittable) utils.Color {
	var rec HitRecord
	if world.Hit(*r, 0, utils.Infinity, &rec) {
		return (*utils.AddVec3(rec.normal, utils.NewVec3(1, 1, 1)).Multiply(0.5))
	}
	var unitDirection utils.Vec3 = r.Direction().UnitVector()
	t := 0.5 * (unitDirection.Y() + 1.0)
	var whiteColor utils.Color = utils.NewVec3(1.0, 1.0, 1.0)
	var blueColor utils.Color = utils.NewVec3(0.5, 0.7, 1.0)
	return (utils.Vec3AddVec3(*whiteColor.Multiply(1.0 - t), *blueColor.Multiply(t)))
}

func main() {

	// Test Vec3 multiple
	testVec := utils.NewVec3(1.0, 1.0, 1.0)
	testVec = *testVec.Multiply(2.0)
	fmt.Fprintf(os.Stderr, "Computed vector as: %+v\n", testVec.ToString())

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int32(float64(imageWidth) / aspectRatio)
	samplesPerPixel := 100

	// World
	var world HittableList
	world.Add((NewSphere(utils.NewVec3(0, 0, -1), 0.5)))
	world.Add(NewSphere(utils.NewVec3(0, -100.5, -1), 100))

	// Camera
	camera := NewCamera()

	// Render
	fmt.Printf("P3\n%+v %+v\n255\n", imageWidth, imageHeight)
	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %+v\n", j)
		for i := 0; i < imageWidth; i++ {
			var pixelColor utils.Color = utils.NewVec3(0, 0, 0)
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + utils.RandomDouble()) / float64(imageWidth-1)
				v := (float64(j) + utils.RandomDouble()) / float64(imageHeight-1)
				r := camera.GetRay(u, v)
				pixelColor = utils.AddVec3(pixelColor, RayColor(&r, world))
			}
			utils.WriteMultiSampleColor(pixelColor, samplesPerPixel)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone.\n")
}
