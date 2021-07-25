package gotracer

import "math"

type Vec3 struct {
	e [3]float64
}

func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{e: [3]float64{x, y, z}}
}

// Basic gettings.
func (v *Vec3) X() float64 {
	return v.e[0]
}

func (v *Vec3) Y() float64 {
	return v.e[1]
}

func (v *Vec3) Z() float64 {
	return v.e[2]
}

func (v *Vec3) R() float64 {
	return v.e[0]
}

func (v *Vec3) G() float64 {
	return v.e[1]
}

func (v *Vec3) B() float64 {
	return v.e[2]
}

// Basic Operators
func (v *Vec3) Add(v2 Vec3) *Vec3 {
	return &Vec3{e: [3]float64{v.e[0] + v2.e[0], v.e[1] + v2.e[1], v.e[2] + v2.e[2]}}
}

func (v *Vec3) Sub(v2 Vec3) *Vec3 {
	return &Vec3{e: [3]float64{v.e[0] - v2.e[0], v.e[1] - v2.e[1], v.e[2] - v2.e[2]}}
}

func (v *Vec3) Mult(v2 Vec3) *Vec3 {
	return &Vec3{e: [3]float64{v.e[0] * v2.e[0], v.e[1] * v2.e[1], v.e[2] * v2.e[2]}}
}

func (v *Vec3) Div(v2 Vec3) *Vec3 {
	return &Vec3{e: [3]float64{v.e[0] / v2.e[0], v.e[1] / v2.e[1], v.e[2] / v2.e[2]}}
}

// Scalar Operators

func (v *Vec3) ScalarMult(t float64) *Vec3 {
	return &Vec3{e: [3]float64{v.e[0] * t, v.e[1] * t, v.e[2] * t}}
}

func (v *Vec3) ScalarDiv(t float64) *Vec3 {
	return &Vec3{e: [3]float64{v.e[0] / t, v.e[1] / t, v.e[2] / t}}
}

// Complex Operators
func (v *Vec3) Dot(v2 Vec3) float64 {
	return v.e[0]*v2.e[0] + v.e[1]*v2.e[1] + v.e[2]*v2.e[2]
}

func (v *Vec3) Cross(v2 Vec3) *Vec3 {
	return &Vec3{e: [3]float64{(v.e[1]*v2.e[2] - v.e[2]*v2.e[1]), -(v.e[0]*v2.e[2] - v.e[2]*v2.e[0]), (v.e[0]*v2.e[1] - v.e[1]*v2.e[0])}}
}

// Attributes
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2])
}

func (v *Vec3) SquaredLength() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

// Modifiers
func (v *Vec3) MakeUnitVector() {
	var k float64 = 1.0 / v.Length()

	v.e[0] = v.e[0] * k
	v.e[1] = v.e[1] * k
	v.e[2] = v.e[2] * k
}

func (v *Vec3) GetUnitVector() *Vec3 {
	return v.ScalarDiv(v.Length())
}
