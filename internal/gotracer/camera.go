package gotracer

import ()

type Camera struct {
	Origin          Vec3
	LowerLeftCorner Vec3
	Horizontal      Vec3
	Vertical        Vec3
}

func NewCamera() *Camera {
	return &Camera{
		Origin:          *NewVec3(0.0, 0.0, 0.0),
		LowerLeftCorner: *NewVec3(-2.0, -1.0, -1.0),
		Horizontal:      *NewVec3(4.0, 0.0, 0.0),
		Vertical:        *NewVec3(0.0, 2.0, 0.0)}
}

func (c *Camera) GetRay(u, v float64) *Ray {
	direction := c.LowerLeftCorner.Add(*c.Horizontal.ScalarMult(u)).Add(*c.Vertical.ScalarMult(v))
	return NewRay(c.Origin, *direction)
}
