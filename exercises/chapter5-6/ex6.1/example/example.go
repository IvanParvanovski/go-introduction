package main

import (
	"image/color"
	"math"
)

type Point struct{X, Y float64}

type ColoredPoint struct {
	Point
	Color color.RGBA
}


// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point 

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i - 1].Distance(path[i])
		}
	}
	return sum
} 

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}



func main() {
	// perim := Path{
	// 	{1, 1},
	// 	{5, 1}, 
	// 	{5, 4},
	// 	{1, 1},
	// }
	
	// r := &Point{1, 2}
	// r.ScaleBy(2)
	// fmt.Println(*r)

	// p := Point{1, 2}
	// pptr := &p
	// pptr.ScaleBy(2)
	// fmt.Println(p)

	// z := Point{1, 2}
	// (&z).ScaleBy(2)
	// fmt.Println(z)

	// fmt.Println(perim.Distance())
	// fmt.Println(Distance(p, q))
	// fmt.Println(p.Distance(q))
}
