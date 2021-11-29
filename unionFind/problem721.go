package unionFind

import "sort"

//Given a list of accounts where each element accounts[i] is a list of strings,
//where the first element accounts[i][0] is a name, and the rest of the elements
//are emails representing emails of the account.
//
// Now, we would like to merge these accounts. Two accounts definitely belong
//to the same person if there is some common email to both accounts. Note that even
//if two accounts have the same name, they may belong to different people as
//people could have the same name. A person can have any number of accounts initially,
//but all of their accounts definitely have the same name.
//
// After merging the accounts, return the accounts in the following format: the
//first element of each account is the name, and the rest of the elements are
//emails in sorted order. The accounts themselves can be returned in any order.
//
//
// Example 1:
//
//
//Input: accounts = [["John","johnsmith@mail.com","john_newyork@mail.com"],[
//"John","johnsmith@mail.com","john00@mail.com"],["Mary","mary@mail.com"],["John",
//"johnnybravo@mail.com"]]
//Output: [["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.
//com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
//Explanation:
//The first and second John's are the same person as they have the common email
//"johnsmith@mail.com".
//The third John and Mary are different people as none of their email addresses
//are used by other accounts.
//We could return these lists in any order, for example the answer [['Mary',
//'mary@mail.com'], ['John', 'johnnybravo@mail.com'],
//['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']]
//would still be accepted.
//
//
// Example 2:
//
//
//Input: accounts = [["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin",
//"Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],["Ethan","Ethan5@m.co","Ethan4@m.co",
//"Ethan0@m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@
//m.co","Fern1@m.co","Fern0@m.co"]]
//Output: [["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.
//co","Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.
//co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co",
//"Fern1@m.co","Fern5@m.co"]]
//
//
//
// Constraints:
//
//
// 1 <= accounts.length <= 1000
// 2 <= accounts[i].length <= 10
// 1 <= accounts[i][j] <= 30
// accounts[i][0] consists of English letters.
// accounts[i][j] (for j > 0) is a valid email.
//
// Related Topics Array String Depth-First Search Breadth-First Search Union
//Find 👍 3394 👎 613

//leetcode submit region begin(Prohibit modification and deletion)
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
	u := newUnionFind(l)
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
