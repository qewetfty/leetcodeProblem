package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	l, r, result := make([]int, n), make([]int, n), make([]int, n)

	l[0] = 1
	for i := 1; i < n; i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	r[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		result[i] = l[i] * r[i]
	}
	return result
}

func main() {

}
