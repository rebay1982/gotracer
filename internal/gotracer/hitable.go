package gotracer

type HitRecord struct {
	point  Vec3
	normal Vec3
	t      float64
}

type Hitable interface {
	Hit(r Ray, t_min float64, t_max float64, record *HitRecord) bool
}
