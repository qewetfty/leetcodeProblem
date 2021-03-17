package math

import (
	"math"
	"math/rand"
)

// Given the radius and x-y positions of the center of a circle, write a function
// randPoint which generates a uniform random point in the circle.
//	Note:
//		input and output values are in floating-point.
//		radius and x-y position of the center of the circle is passed into the class constructor.
//		a point on the circumference of the circle is considered to be in the circle.
//		randPoint returns a size 2 array containing x-position and y-position of the random point, in that order.
//	Example 1:
//		Input:
//			["Solution","randPoint","randPoint","randPoint"]
//			[[1,0,0],[],[],[]]
//		Output: [null,[-0.72939,-0.65505],[-0.78502,-0.28626],[-0.83119,-0.19803]]
//	Example 2:
//		Input:
//			["Solution","randPoint","randPoint","randPoint"]
//			[[10,5,-7.5],[],[],[]]
//		Output: [null,[11.52438,-8.33273],[2.46992,-16.21705],[11.13430,-12.42337]]
//	Explanation of Input Syntax:
//		The input is two lists: the subroutines called and their arguments. Solution's
//		constructor has three arguments, the radius, x-position of the center, and
//		y-position of the center of the circle. randPoint has no arguments. Arguments
//		are always wrapped with a list, even if there aren't any.

type Solution struct {
	x      float64
	y      float64
	radius float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		x:      x_center,
		y:      y_center,
		radius: radius,
	}
}

func (this *Solution) RandPoint() []float64 {
	x0, y0 := this.x-this.radius, this.y-this.radius
	for {
		xg, yg := x0+rand.Float64()*this.radius*2, y0+rand.Float64()*this.radius*2
		if math.Sqrt((xg-this.x)*(xg-this.x)+(yg-this.y)*(yg-this.y)) <= this.radius {
			return []float64{xg, yg}
		}
	}
}
