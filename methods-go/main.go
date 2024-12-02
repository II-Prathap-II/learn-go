package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y      int
	distance1 func(p Point) int
}

type Path []Point

// Methods are functions which are that can be defined for any named types
// the function distance is defined for type Point - this can be identified by the additional paramter before the function name. This is called function receiver

func (p Point) distance(q Point) int {
	return int(math.Hypot(float64(q.X-p.X), float64(q.Y-p.Y)))
}

// Methods can not only be defined for named struct types, but also for any named types. For example, here Path is of type []Point. But the methods are defined for this type as well. Check out the method receiver before the function name.

func (p Path) distance() int {
	var sum float64

	for i := range p {
		if i > 0 {
			sum += float64(p[i-1].distance(p[i]))
		}
	}

	return int(sum)
}

func main() {
	// Methods should not be confused with the function type fields in the structs.
	// distance1 is an example of a field function. The implementation of the function was not defined in the struct, but after the creation of an object of the struct.

	p := Point{X: 1, Y: 1}
	p.distance1 = func(q Point) int {
		return int(math.Hypot(float64(p.X-q.X), float64(p.Y-q.Y)))
	}
	fmt.Println(p.distance1(Point{X: 2, Y: 2}))

	path := Path{
		Point{X: 1, Y: 1},
		Point{X: 5, Y: 1},
		Point{X: 5, Y: 4},
		Point{X: 1, Y: 1},
	}
	fmt.Println("Perimter is ", path.distance())
}
