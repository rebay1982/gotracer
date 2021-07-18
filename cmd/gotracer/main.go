package main

import (
	//	"bufio"
	"fmt"
	"os"
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

			var r float64 = float64(i) / float64(nx)
			var g float64 = float64(j) / float64(ny)
			var b float64 = 0.2

			var ir = int(255.99 * r)
			var ig = int(255.99 * g)
			var ib = int(255.99 * b)

			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}

	fmt.Println("Done...")
}
