package main

import (
	"go/raytracingweekend/utils"
)

type Ray struct {
	orig utils.Point3
	dir  utils.Vec3
}

func NewRay(origin utils.Point3, direction utils.Vec3) Ray {
	return Ray{
		orig: origin,
		dir:  direction,
	}
}

func (r Ray) Origin() utils.Point3 {
	return r.orig
}

func (r Ray) Direction() utils.Vec3 {
	return r.dir
}

func (r Ray) At(t float64) utils.Point3 {
	return (utils.Vec3AddVec3(r.orig, *r.dir.Multiply(t)))
}
