package data

type IntHeapAsc []int

func (h IntHeapAsc) Len() int {
	return len(h)
}

func (h IntHeapAsc) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeapAsc) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeapAsc) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeapAsc) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntHeapDesc []int

func (h IntHeapDesc) Len() int {
	return len(h)
}

func (h IntHeapDesc) Less(i, j int) bool {
	return h[j] < h[i]
}

func (h IntHeapDesc) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeapDesc) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeapDesc) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
