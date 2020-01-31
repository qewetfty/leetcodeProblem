package data

type DulNode struct {
	Key   int
	Value int
	Prev  *DulNode
	Next  *DulNode
}
