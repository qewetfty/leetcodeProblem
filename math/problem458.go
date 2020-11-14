package math

import (
	"fmt"
	"math"
)

// There are 1000 buckets, one and only one of them is poisonous, while the rest
// are filled with water. They all look identical. If a pig drinks the poison it
// will die within 15 minutes. What is the minimum amount of pigs you need to
// figure out which bucket is poisonous within one hour?
// Answer this question, and write an algorithm for the general case.
//	General case:
//		If there are n buckets and a pig drinking poison will die within m minutes,
//		how many pigs (x) you need to figure out the poisonous bucket within p
//		minutes? There is exactly one bucket with poison.
//	Note:
//		A pig can be allowed to drink simultaneously on as many buckets as one would
//		like, and the feeding takes no time.
//		After a pig has instantly finished drinking buckets, there has to be a cool down time of m minutes. During this
//		time, only observation is allowed and no feedings at all.
//		Any given bucket can be sampled an infinite number of times (by an unlimited number of pigs).

// 分析：🐷一共有 minutesToTest/minutesToDie + 1 种状态，例如：
//	minutesToTest/minutesToDie = 0，表示只有存活状态
//	minutesToTest/minutesToDie = 1， 有存活和死亡状态
// 	minutesToTest/minutesToDie = 2， 存活，第一次测试死亡，第二次测试死亡三种状态
// 	......
// 而x只🐷能够测试 s^x个水桶
// 因此只需找到一个x，使得 s ^ x >= buckets， 即 x >= logstate(buckets)，换底公示即为：x >= log(buckets) / log(states)
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	states := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(states))))
}

func testProblem458() {
	fmt.Println(poorPigs(1000, 15, 60))
	fmt.Println(poorPigs(4, 15, 15))
}
