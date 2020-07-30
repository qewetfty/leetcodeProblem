package data

/*
	并查集的类实现
*/

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	u := new(UnionFind)
	u.count = n
	u.parent = make([]int, n)
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

func (u *UnionFind) Parent(p int) int {
	root, i := p, p
	for root != u.parent[root] {
		root = u.parent[root]
	}
	for u.parent[i] != i {
		i, u.parent[i] = u.parent[i], root
	}
	return root
}

func (u *UnionFind) Union(i, j int) {
	p1, p2 := u.Parent(i), u.Parent(j)
	u.parent[p1] = p2
}
