package main

import (
	//	"bufio"
	"fmt"
	"os"

	gt "github.com/rebay1982/gotracer/internal/gotracer"
)

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

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			colour := gt.NewVec3(float64(i)/float64(nx), float64(j)/float64(ny), 0.2)

			var ir = int(255.99 * colour.R())
			var ig = int(255.99 * colour.G())
			var ib = int(255.99 * colour.B())

			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}

	fmt.Println("Done...")
}
