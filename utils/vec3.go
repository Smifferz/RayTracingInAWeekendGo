package utils

import (
	"fmt"
	"math"
)

type Vec3 struct {
	e [3]float64
}

func NewVec3(coords ...float64) Vec3 {
	if coords != nil {
		return Vec3{
			e: [3]float64{coords[0], coords[1], coords[2]},
		}
	} else {
		return Vec3{
			e: [3]float64{0, 0, 0},
		}
	}
}

// Vec3 Helper functions
func (v3 Vec3) X() float64 {
	return v3.e[0]
}

func (v3 Vec3) Y() float64 {
	return v3.e[1]
}

func (v3 Vec3) Z() float64 {
	return v3.e[2]
}

func (v3 Vec3) Inverse() Vec3 {
	return NewVec3(
		-v3.e[0],
		-v3.e[1],
		-v3.e[2],
	)
}

func (v3 Vec3) At(i int) float64 {
	return v3.e[i]
}

func (v3 Vec3) Add(v *Vec3) *Vec3 {
	v3.e[0] += v.e[0]
	v3.e[1] += v.e[1]
	v3.e[2] += v.e[2]
	return &v3
}

func (v3 Vec3) Multiply(t float64) *Vec3 {
	v3.e[0] *= t
	v3.e[1] *= t
	v3.e[2] *= t
	return &v3
}

func (v3 Vec3) Divide(t float64) *Vec3 {
	v3.e[0] /= t
	v3.e[1] /= t
	v3.e[2] /= t
	return &v3
}

func (v3 Vec3) Length() float64 {
	return math.Sqrt(v3.LengthSquared())
}

func (v3 Vec3) LengthSquared() float64 {
	return v3.e[0]*v3.e[0] + v3.e[1]*v3.e[1] + v3.e[2]*v3.e[2]
}

// Vec3 Utility functions
func (v3 Vec3) ToString() string {
	return fmt.Sprintf("%+v %+v %+v", v3.e[0], v3.e[1], v3.e[2])
}

func Vec3AddVec3(u Vec3, v Vec3) Vec3 {
	return Vec3{
		e: [3]float64{u.e[0] + v.e[0], u.e[1] + v.e[1], u.e[2] + v.e[2]},
	}
}

func AddVec3(vecs ...Vec3) Vec3 {
	addedX := 0.0
	addedY := 0.0
	addedZ := 0.0
	for _, vec := range vecs {
		addedX += vec.e[0]
		addedY += vec.e[1]
		addedZ += vec.e[2]
	}

	return Vec3{
		e: [3]float64{addedX, addedY, addedZ},
	}
}

func Vec3MinusVec3(u Vec3, v Vec3) Vec3 {
	return Vec3{
		e: [3]float64{u.e[0] - v.e[0], u.e[1] - v.e[1], u.e[2] - v.e[2]},
	}
}

func MinusVec3(vecs ...Vec3) Vec3 {
	initialized := false
	var x float64
	var y float64
	var z float64
	for _, vec := range vecs {
		if !initialized {
			x = vec.e[0]
			y = vec.e[1]
			z = vec.e[2]
			initialized = true
		} else {
			x -= vec.e[0]
			y -= vec.e[1]
			z -= vec.e[2]
		}
	}

	return Vec3{
		e: [3]float64{x, y, z},
	}
}

func MultiplyVec3ByVec3(u Vec3, v Vec3) Vec3 {
	return Vec3{
		e: [3]float64{u.e[0] * v.e[0], u.e[1] * v.e[1], u.e[2] * v.e[2]},
	}
}

func MultiplyVec3ByFloat(v Vec3, t float64) Vec3 {
	return Vec3{
		e: [3]float64{t * v.e[0], t * v.e[1], t * v.e[2]},
	}
}

func DivideVec3ByFloat(v Vec3, t float64) Vec3 {
	return MultiplyVec3ByFloat(v, (1 / t))
}

func Dot(u Vec3, v Vec3) float64 {
	return (u.e[0]*v.e[0] + u.e[1]*v.e[1] + u.e[2]*v.e[2])
}

func Cross(u Vec3, v Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			u.e[1]*v.e[2] - u.e[2]*v.e[1],
			u.e[2]*v.e[0] - u.e[0]*v.e[2],
			u.e[0]*v.e[1] - u.e[1]*v.e[0],
		},
	}
}

func (v Vec3) UnitVector() Vec3 {
	return *(v.Divide(v.Length()))
}

type Point3 = Vec3
type Color = Vec3
