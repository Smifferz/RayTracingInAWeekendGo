package main

import "go/raytracingweekend/utils"

type Material interface {
	Scatter(rIn Ray, rec HitRecord, attenuation *utils.Color, scattered *Ray) bool
}

type Lambertian struct {
	albedo utils.Color
}

func NewLambertian(a utils.Color) Lambertian {
	return Lambertian{
		albedo: a,
	}
}

func (l Lambertian) Scatter(rIn Ray, rec HitRecord, attenuation *utils.Color, scattered *Ray) bool {
	scatterDirection := utils.AddVec3(rec.normal, utils.RandomUnitVector())

	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = rec.normal
	}

	*scattered = NewRay(rec.p, scatterDirection)
	*attenuation = l.albedo
	return true
}

type Metal struct {
	albedo utils.Color
}

func NewMetal(a utils.Color) Metal {
	return Metal{
		albedo: a,
	}
}

func (m Metal) Scatter(rIn Ray, rec HitRecord, attenuation *utils.Color, scattered *Ray) bool {
	var reflected utils.Vec3 = utils.Reflect(rIn.Direction().UnitVector(), rec.normal)
	*scattered = NewRay(rec.p, reflected)
	*attenuation = m.albedo
	return utils.Dot(scattered.Direction(), rec.normal) > 0
}
