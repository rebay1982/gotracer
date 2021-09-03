package gotracer

import "fmt"

type HitableList struct {
	List     []Hitable
	ListSize int
}

func NewHitableList(list []Hitable, listSize int) *HitableList {
	return &HitableList{List: list, ListSize: listSize}
}

func (hl *HitableList) Hit(r Ray, t_min float64, t_max float64, record *HitRecord) bool {
	var temp_rec HitRecord
	var hit_anything bool = false
	var closest_so_far float64 = t_max

	// Go through all the elements in the "hitable" list and hit them --
	// also making sure we find the closest item that is hit.
	for _, hitable := range hl.List {
		if hitable.Hit(r, t_min, closest_so_far, &temp_rec) {

			hit_anything = true
			closest_so_far = temp_rec.T
			record.T = temp_rec.T
			record.Normal = temp_rec.Normal
			record.Point = temp_rec.Point

		}
	}
	return hit_anything
}
