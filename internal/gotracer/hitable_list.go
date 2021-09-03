package gotracer

type HitableList struct {
	list     []Hitable
	listSize int
}

func NewHitableList(list []Hitable, listSize int) *HitableList {
	return &HitableList{list: list, listSize: listSize}
}

func (hl *HitableList) Hit(r Ray, t_min float64, t_max float64, record *HitRecord) bool {
	var temp_rec HitRecord
	var hit_anything bool = true
	var closest_so_far float64 = t_max

	// Go through all the elements in the "hitable" list and hit them --
	// also making sure we find the closest item that is hit.
	for _, hitable := range hl.list {
		if hitable.Hit(r, t_min, closest_so_far, &temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.t
			record = &temp_rec
		}
	}
	return hit_anything
}
