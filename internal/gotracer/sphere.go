package gotracer

import (
	"math"
)

type Sphere struct {
	center Vec3
	radius float64
}

func NewSphere(c Vec3, r float64) *Sphere {
	return &Sphere{center: c, radius: r}

}

func (s *Sphere) Hit(r Ray, tMin float64, tMax float64, record HitRecord) bool {
	origin := r.GetOrigin()
	oc := origin.Sub(s.center)

	direction := r.GetDirection()
	var a float64 = direction.SquaredLength()
	var half_b float64 = oc.Dot(r.GetDirection())
	var c float64 = oc.SquaredLength() - (s.radius * s.radius)

	var discriminant float64 = half_b*half_b - a*c

	// If the discriminant is negative, we're looking at the sphere's surface
	// from inside the sphere. (or behind it)
	if discriminant < 0 {
		return false

	} else {

		// We're going to find the nearest root that falls between tMin and tMax
		var sqrtd = math.Sqrt(discriminant)
		var root = (-half_b - sqrtd) / a

		if root < tMin || root > tMax {
			root = (-half_b + sqrtd) / a

			if root < tMin || root > tMax {
				return false

			}
		}

		record.t = root
		record.point = *r.PointAtParameter(record.t)
		record.normal = *record.point.Sub(s.center).ScalarDiv(s.radius)

		return true
	}
}