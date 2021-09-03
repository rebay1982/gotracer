package gotracer

type HitRecord struct {
	Point  Vec3
	Normal Vec3
	T      float64
}

type Hitable interface {
	Hit(r Ray, t_min float64, t_max float64, record *HitRecord) bool
}
