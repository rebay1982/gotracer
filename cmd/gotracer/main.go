package main

import (
	"fmt"
	"math"
	"os"

	gt "github.com/rebay1982/gotracer/internal/gotracer"
)

func color(r *gt.Ray, world gt.Hitable) *gt.Vec3 {
	var rec gt.HitRecord

	if world.Hit(*r, 0.0, math.MaxFloat64, &rec) {
		//fmt.Println("Hit something")
		return gt.NewVec3(rec.Normal.X()+1.0, rec.Normal.Y()+1.0, rec.Normal.Z()+1.0).ScalarMult(0.5)

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

	// Image size
	var nx, ny int = 200, 100

	// Write PPM file header.
	fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	// Define some vectors used for the computation -- see page 10.
	var lowerLeftCorner *gt.Vec3 = gt.NewVec3(-2.0, -1.0, -1.0)
	var origin *gt.Vec3 = gt.NewVec3(0.0, 0.0, 0.0)
	var horizontal *gt.Vec3 = gt.NewVec3(4.0, 0.0, 0.0)
	var vertical *gt.Vec3 = gt.NewVec3(0.0, 2.0, 0.0)

	var worldList []gt.Hitable = make([]gt.Hitable, 2)
	worldList[0] = gt.NewSphere(*gt.NewVec3(0.0, 0.0, -1.0), 0.5)
	worldList[1] = gt.NewSphere(*gt.NewVec3(0.0, -100.5, -1.0), 100)

	var world gt.Hitable = gt.NewHitableList(worldList, 2)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			var u float64 = float64(i) / float64(nx)
			var v float64 = float64(j) / float64(ny)
			direction := lowerLeftCorner.Add(*horizontal.ScalarMult(u)).Add(*vertical.ScalarMult(v))
			ray := gt.NewRay(*origin, *direction)

			colour := color(ray, world) //gt.NewVec3(float64(i)/float64(nx), float64(j)/float64(ny), 0.2)

			var ir = int(255.99 * colour.R())
			var ig = int(255.99 * colour.G())
			var ib = int(255.99 * colour.B())

			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}

	fmt.Println("Done...")
}
