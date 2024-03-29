
package main

import (
	"fmt"
	"math"
)

import "image/color"

type Point struct{X, Y float64}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX+dY*dY)
}


func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255,255}
	var p = ColoredPoint{Point{1,1}, red}
	// Point 方法被纳入到 ColoredPoint 类型中，以这种方式，内嵌允许构成复杂的类型，
	// 该类型由许多字段构成，每个字段提供一些方法

	var q = ColoredPoint{Point{5,4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}

func init() {
	p := Point{1,2}
	q := Point{4,6}

	distance := Point.Distance
	fmt.Println(distance(p,q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p,2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}

func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}

	p := ColoredPoint{&Point{1,1}, red}
	q := ColoredPoint{&Point{5,4}, blue}
	fmt.Println(p.Distance(*q.Point))
	q.Point = p.Point
	
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point)
	// 选择子 p.Distance 可以赋予一个方法变量，它是一个函数，把方法(Point.Distance)
	// 绑定到一个接收者 p 上。函数只需要提供实参而不需要提供接收者就能够调用。
	distanceFromP := p.Distance
	var origin Point
	fmt.Println(distanceFromP(origin))
}