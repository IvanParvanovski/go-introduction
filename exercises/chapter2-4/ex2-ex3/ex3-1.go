package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin, cos = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := range cells {
		for j := range cells {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			points := []float64{ax, ay, bx, by, cx, cy, dx, dy}
			invalid := false
			for _, p := range points {
				if math.IsNaN(p) || math.IsInf(p, 0) {
					invalid = true
				}
			}
			if invalid {
				continue
			}

			// z := ((az + bz + cz + dz) / 4.0)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}
func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	
	z := f(x, y)

	sx := width/2 + (x-y)*cos*xyscale
	sy := height/2 + (x+y)*sin*xyscale - z*zscale
	return sx, sy
}

// func f(x, y float64) float64 {
// 	r := math.Hypot(x, y)

// 	if r == math.Inf(-1) || r == math.Inf(1) {
// 		fmt.Println("Error")
// 		return 0

// 	}

// 	return math.Sin(r)
// }

// Egg box
func f(x, y float64) float64 {
	return math.Sin(x) * math.Sin(y) / 4
}

// Mogul
// func f(x, y float64) float64 {
// 	return math.Cos(x) * math.Cos(y) / 4
// }

// Saddle
// func f(x, y float64) float64 {
// 	return (x*x - y*y) 
// }