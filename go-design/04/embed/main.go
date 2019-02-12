// Embed demonstrates basic struct embedding
package main

import "fmt"

type Point struct{X, Y int}

type Circle struct{
	Point
	Radius int
}

type Wheel struct{
	Circle
	Spokes int
}

func main()  {
	var w Wheel

	w = Wheel{
		Circle{
			Point{8,8},
			5,
		},
		20,
	}
	fmt.Printf("%#v\n", w)
	w.X = 42
	w.Circle.Point.Y = 24
	fmt.Printf("%#v\n", w)
}