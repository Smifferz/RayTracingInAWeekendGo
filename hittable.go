package main

import "go/raytracingweekend/utils"

type HitRecord struct {
	p         utils.Point3
	normal    utils.Vec3
	material  Material
	t         float64
	frontFace bool
}

func (rec *HitRecord) SetFaceNormal(r Ray, outwardNormal utils.Vec3) {
	rec.frontFace = utils.Dot(r.Direction(), outwardNormal) < 0
	if rec.frontFace {
		rec.normal = outwardNormal
	} else {
		rec.normal = outwardNormal.Inverse()
	}
}

type Hittable interface {
	Hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool
}
