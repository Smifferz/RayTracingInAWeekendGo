package main

import (
	"go/raytracingweekend/utils"
	"math"
)

type Sphere struct {
	center   utils.Point3
	radius   float64
	material Material
}

func NewSphere(cen utils.Point3, r float64, m Material) Sphere {
	return Sphere{
		center:   cen,
		radius:   r,
		material: m,
	}
}

func (s Sphere) Hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	oc := utils.MinusVec3(r.Origin(), s.center)
	a := r.Direction().LengthSquared()
	halfB := utils.Dot(oc, r.Direction())
	c := oc.LengthSquared() - (s.radius * s.radius)

	discriminant := (halfB * halfB) - (a * c)
	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range
	root := (-halfB - sqrtd) / a
	if root < tMin || tMax < root {
		return false
	}

	rec.t = root
	rec.p = r.At(rec.t)
	outwardNormal := utils.MinusVec3(rec.p, s.center).Divide(s.radius)
	rec.SetFaceNormal(r, *outwardNormal)
	rec.material = s.material

	return true
}
