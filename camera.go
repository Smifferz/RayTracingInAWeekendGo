package main

import "go/raytracingweekend/utils"

type Camera struct {
	origin          utils.Point3
	lowerLeftCorner utils.Point3
	horizontal      utils.Vec3
	vertical        utils.Vec3
}

func NewCamera() Camera {
	aspectRatio := 16.0 / 9.0
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	var origin utils.Point3 = utils.NewVec3(0, 0, 0)
	horizontal := utils.NewVec3(viewportWidth, 0, 0)
	vertical := utils.NewVec3(0, viewportHeight, 0)
	lowerLeftCorner := utils.MinusVec3(origin, *horizontal.Divide(2),
		*vertical.Divide(2), utils.NewVec3(0, 0, focalLength))

	return Camera{
		origin:          origin,
		lowerLeftCorner: lowerLeftCorner,
		horizontal:      horizontal,
		vertical:        vertical,
	}
}

func (c Camera) GetRay(u float64, v float64) Ray {
	return NewRay(c.origin, utils.AddVec3(c.lowerLeftCorner, *c.horizontal.Multiply(u), *c.vertical.Multiply(v), c.origin.Inverse()))
}
