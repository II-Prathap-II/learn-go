package main

import (
	"encoding/json"
	"fmt"
)

type Circle struct {
	Point
	Radius       int
	privateField int
}

type Point struct {
	X int `json:"point x"`
	Y int `json:"point Y"`
}

func main() {
	origin := Point{0, 0}

	fmt.Println("before moving by reference", origin)
	movePointLinearlyBy_reference(&origin, 1)
	fmt.Println("after moving by reference", origin)

	fmt.Println("before moving by value", origin)
	origin = movePointLinearlyBy_value(origin, 1)
	fmt.Println("after moving by value", origin)

	somePoint := Point{10, 10}
	dupeOfSomePoint := Point{10, 10}

	// if the field are comparable it is possible to compare structs of the "same type" using ==
	// use reflect.DeepEqual to compare the slice fields
	fmt.Println("comparing points", somePoint == dupeOfSomePoint)
	fmt.Println("comparing points", somePoint == origin)

	// A field a struct can have the type as another struct.
	// This can act as short hand notation without having to mention the intermediate struct explicitly
	var c Circle
	c.X = 0 // equivalent to c.point.x
	c.Y = 0 // equivalent to c.point.y
	c.Radius = 5
	fmt.Println(c)

	// But note that the shorthand like below doens't work
	// c := circle {0,0,5}
	// c := circle {x:0,y:0,5}
	// maximum shorthand is as below
	d := Circle{Point{0, 0}, 4, 3}
	fmt.Printf("%#v", d)
	fmt.Println("")

	// We have created a slice of nested structs
	// marshalled it using json.marshal and json.MarshalIndent
	// Note that only the public fields of the struct are marshalled.
	// Note that the exported fields' names in the output can be changed by using the field tags

	circles := []Circle{
		{Point: Point{1, 1}, Radius: 5, privateField: 100},
		{Point: Point{1, 1}, Radius: 2, privateField: 100},
	}

	marshalCircles, _ := json.Marshal(circles)
	marshalCirclesIndent, _ := json.MarshalIndent(circles, "", " ")

	fmt.Println(string(marshalCircles))
	fmt.Println("Below is the human readable format")
	fmt.Println(string(marshalCirclesIndent))

	// The incoming json data can be unmarshalled using json.unmarshal
	// We can take only the fields that we need and discard others
	// Note how we take only point x from the json. We have defined a local struct which has a public field with json tags that exactly matched the json
	// Notice how the point y is 0 because the local struct has a private field

	var radiusOfXPoints []struct {
		PointX int `json:"point x"`
	}
	json.Unmarshal(marshalCircles, &radiusOfXPoints)
	fmt.Println(radiusOfXPoints)

	var radiusOfYPoints []struct {
		pointY int `json:"point y"`
	}
	json.Unmarshal(marshalCircles, &radiusOfYPoints)
	fmt.Println(radiusOfYPoints)

}

func movePointLinearlyBy_reference(p *Point, move int) {
	p.X += move
	p.Y += move
}

func movePointLinearlyBy_value(p Point, move int) Point {
	p.X += move
	p.Y += move

	return p
}
