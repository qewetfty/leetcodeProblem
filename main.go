package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	l := len(accounts)
	// 构造email的map
	emailMap := make(map[string][]int)
	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			emailMap[account[j]] = append(emailMap[account[j]], i)
		}
	}
	// 合并unionFind
	u := NewUnionFind(l)
	for _, needUnion := range emailMap {
		if len(needUnion) < 2 {
			continue
		}
		n1 := needUnion[0]
		for i := 1; i < len(needUnion); i++ {
			u.Union(n1, needUnion[i])
		}
	}
	// 构造对应去重结果
	unionMap := make(map[int]map[string]bool)
	for i, account := range accounts {
		var unionEmailMap map[string]bool
		root := u.Parent(i)
		if _, ok := unionMap[root]; ok {
			unionEmailMap = unionMap[root]
		} else {
			unionEmailMap = make(map[string]bool)
		}
		for j := 1; j < len(account); j++ {
			unionEmailMap[account[j]] = true
		}
		unionMap[root] = unionEmailMap
	}
	result := make([][]string, 0)
	for i, emails := range unionMap {
		name := accounts[i][0]
		emailList := make([]string, 0)
		for email := range emails {
			emailList = append(emailList, email)
		}
		sort.Strings(emailList)
		emailList = append([]string{name}, emailList...)
		result = append(result, emailList)
	}
	return result
}

type unionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) unionFind {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	return unionFind{
		parent: nums,
		count:  n,
	}
}

func (u *unionFind) Parent(n int) int {
	root, i := n, n
	for root != u.parent[root] {
		root = u.parent[root]
	}
	for i != u.parent[i] {
		i, u.parent[i] = u.parent[i], root
	}
	return root
}

func (u *unionFind) Union(i, j int) {
	m, n := u.Parent(i), u.Parent(j)
	if m == n {
		return
	}
	u.parent[m] = n
	u.count--
}

func main() {
	fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}))
}
