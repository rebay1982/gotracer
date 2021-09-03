package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	gt "github.com/rebay1982/gotracer/internal/gotracer"
)

// random_in_unit_sphere : This functions uses a rejection algorithm
// to find a random point in a unit sphere. We evaluate a randomly selected
// point in a unit cube and reject it if the randomly selected point isn't
// in the unit sphere centered in the cube.
//
// Do this until you find a random point that fits this requirement.
func random_in_unit_sphere() gt.Vec3 {
	var s gt.Vec3

	for ok := true; ok; ok = !(s.SquaredLength() >= 1.0) {
		unit := gt.NewVec3(1.0, 1.0, 1.0)
		s = *gt.NewVec3(rand.Float64(), rand.Float64(), rand.Float64()).ScalarMult(2.0).Sub(*unit)
	}

	return s
}

func color(r *gt.Ray, world gt.Hitable, depth int) *gt.Vec3 {
	var rec gt.HitRecord

	if depth < 0 {
		return gt.NewVec3(0.0, 0.0, 0.0)
	}

	// 0.001 minimum ignores rays that hit the object they're reflectiing off of at not exactly 0.
	// See te end of chapter 7 for explanation.
	if world.Hit(*r, 0.001, math.MaxFloat64, &rec) {
		target := rec.Point.Add(rec.Normal).Add(random_in_unit_sphere())
		return color(gt.NewRay(rec.Point, *target.Sub(rec.Point)), world, depth-1).ScalarMult(0.5)

	} else {
		direction := r.GetDirection()
		unitDirection := direction.GetUnitVector()

		t := 0.5 * (unitDirection.Y() + 1.0)
		aVecScalar := 1.0 - t
		aVec := gt.NewVec3(aVecScalar, aVecScalar, aVecScalar)
		bVec := gt.NewVec3(0.5, 0.7, 1.0).ScalarMult(t)

		return aVec.Add(*bVec)
	}
}

func main() {

	// Open a file.
	f, err := os.Create("./output.ppm")
	if err != nil {
		fmt.Errorf("Unable to open file: %v", err)
	}
	defer f.Close()

	// ns, ny is the image size
	// ns is the number of samples for the anti-aliasing.
	// max_depth is the max iterations that a ray will bounce around.
	var nx, ny, ns, max_depth int = 200, 100, 100, 50
	// Write PPM file header.
	fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	var worldList []gt.Hitable = make([]gt.Hitable, 2)
	worldList[0] = gt.NewSphere(*gt.NewVec3(0.0, 0.0, -1.0), 0.5)
	worldList[1] = gt.NewSphere(*gt.NewVec3(0.0, -100.5, -1.0), 100)
	var world gt.Hitable = gt.NewHitableList(worldList, 2)

	var camera *gt.Camera = gt.NewCamera()

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			var colour gt.Vec3 = *gt.NewVec3(0.0, 0.0, 0.0)

			for s := 0; s < ns; s++ {
				var u float64 = (float64(i) + rand.Float64()) / float64(nx)
				var v float64 = (float64(j) + rand.Float64()) / float64(ny)
				ray := camera.GetRay(u, v)

				colour = *colour.Add(*color(ray, world, max_depth))
			}

			colour = *colour.ScalarDiv(float64(ns)) // Devide the sum with the number of samples.
			colour = *gt.NewVec3(math.Sqrt(colour.R()), math.Sqrt(colour.G()), math.Sqrt(colour.B()))

			var ir = int(255.99 * colour.R())
			var ig = int(255.99 * colour.G())
			var ib = int(255.99 * colour.B())

			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}

	fmt.Println("Done...")
}
