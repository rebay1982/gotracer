package gotracer

type Ray struct {
	a, b Vec3
}

func NewRay(a, b Vec3) *Ray {
	return &Ray{a: a, b: b}
}

func (r *Ray) GetOrigin() Vec3 {
	return r.a
}

func (r *Ray) GetDirection() Vec3 {
	return r.b
}

func (r *Ray) PointAtParameter(t float64) *Vec3 {
	return r.a.Add(*r.b.ScalarMult(t)) // A + t*B
}
