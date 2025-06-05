package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := range cells {
		for j := range cells {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			// SKIP INVALID POLYGONS
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

			// calculate height
			z := ((az + bz + cz + dz) / 4.0)
			color := lerpRedBlue(z/20.0 + 1.0) // experimental

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, zscale := ripple(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z * zscale
}

func ripple(x, y float64) (float64, float64) {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r, 0.4 * height
}

func saddle(x, y float64) (float64, float64) {
	return x*x - y*y, height / (2 * xyrange * xyrange)
}

func lerpRedBlue(t float64) string {
	// clamp t to [0,1] to avoid surprises
	t = math.Max(0, math.Min(1, t))

	// Red fades 255→0, Blue grows 0→255, Green stays 0
	r := uint8(math.Round((1 - t) * 255))
	g := uint8(0)
	b := uint8(math.Round(t * 255))

	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
