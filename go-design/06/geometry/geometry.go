
package geometry

import "math"

type Point struct{X, Y float64}
// 普通的函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
// Point类型的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
/*
附加参数 p 称为方法的接收者
p.Distance(q)
表达式 p.Distance 称为选择子（selector），因为它为接受者 p 选择合适的 Distance 方法。
*/

type Path []Point
// 不同类型的 Distance 接收者不同
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}