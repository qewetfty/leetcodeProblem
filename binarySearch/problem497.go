package binarySearch

// Given a list of non-overlapping axis-aligned rectangles rects, write a function pick which randomly and uniformily
// picks an integer point in the space covered by the rectangles.
//	Note:
//		An integer point is a point that has integer coordinates.
//		A point on the perimeter of a rectangle is included in the space covered by the rectangles.
//		ith rectangle = rects[i] = [x1,y1,x2,y2], where [x1, y1] are the integer coordinates of the bottom-left corner,
//			and [x2, y2] are the integer coordinates of the top-right corner.
//		length and width of each rectangle does not exceed 2000.
//		1 <= rects.length <= 100
//		pick return a point as an array of integer coordinates [p_x, p_y]
//		pick is called at most 10000 times.
//	Example 1:
//		Input:
//			["Solution","pick","pick","pick"]
//			[[[[1,1,5,5]]],[],[],[]]
//		Output:
//			[null,[4,1],[4,1],[3,3]]
//	Example 2:
//		Input:
//			["Solution","pick","pick","pick","pick","pick"]
//			[[[[-2,-2,-1,-1],[1,0,3,0]]],[],[],[],[],[]]
//		Output:
//			[null,[-1,-2],[2,0],[-2,-1],[3,0],[-2,-2]]
//		Explanation of Input Syntax:
//			The input is two lists: the subroutines called and their arguments. Solution's constructor has one argument,
//			the array of rectangles rects. pick has no arguments. Arguments are always wrapped with a list, even if there aren't any.

type Solution struct {
	rects     [][]int
	nextPoint []int
	cur       int
}

func Constructor(rects [][]int) Solution {
	return Solution{
		rects:     rects,
		nextPoint: []int{rects[0][0], rects[0][1]},
		cur:       0,
	}
}

func (this *Solution) Pick() []int {
	point := this.nextPoint
	curX, curY := point[0], point[1]
	curRect := this.rects[this.cur]
	x1, x2, y2 := curRect[0], curRect[2], curRect[3]
	if curX == x2 && curY == y2 {
		if this.cur == len(this.rects)-1 {
			this.cur = 0
		} else {
			this.cur++
		}
		this.nextPoint = []int{this.rects[this.cur][0], this.rects[this.cur][1]}
	} else if curX == x2 {
		this.nextPoint = []int{x1, curY + 1}
	} else {
		this.nextPoint = []int{curX + 1, curY}
	}
	return point
}
