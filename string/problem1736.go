package string

import (
	"fmt"
	"strings"
)

// You are given a string time in the form of hh:mm, where some of the digits in the string are hidden (represented by ?).
// The valid times are those inclusively between 00:00 and 23:59.
// Return the latest valid time you can get from time by replacing the hidden digits.
//	Example 1:
//		Input: time = "2?:?0"
//		Output: "23:50"
//		Explanation: The latest hour beginning with the digit '2' is 23 and the latest minute ending with the digit '0' is 50.
//	Example 2:
//		Input: time = "0?:3?"
//		Output: "09:39"
//	Example 3:
//		Input: time = "1?:22"
//		Output: "19:22"
//	Constraints:
//		time is in the format hh:mm.
//		It is guaranteed that you can produce a valid time from the given string.

var hourMap = map[string]string{"0?": "09", "1?": "19", "2?": "23", "?0": "20", "?1": "21", "?2": "22", "?3": "23",
	"?4": "14", "?5": "15", "?6": "16", "?7": "17", "?8": "18", "?9": "19"}

var minuteMap = map[string]string{"0?": "09", "1?": "19", "2?": "29", "3?": "39", "4?": "49", "5?": "59", "?0": "50",
	"?1": "51", "?2": "52", "?3": "53", "?4": "54", "?5": "55", "?6": "56", "?7": "57", "?8": "58", "?9": "59"}

func maximumTime(time string) string {
	split := strings.Split(time, ":")
	hour, minute := split[0], split[1]
	result := ""
	if "??" == hour {
		result = result + "23"
	} else if hour[0] != '?' && hour[1] != '?' {
		result = result + hour
	} else {
		result = result + hourMap[hour]
	}
	result = result + ":"
	if "??" == minute {
		result = result + "59"
	} else if minute[0] != '?' && minute[1] != '?' {
		result = result + minute
	} else {
		result = result + minuteMap[minute]
	}
	return result
}

func testProblem1736() {
	fmt.Println(maximumTime("2?:?0"))
	fmt.Println(maximumTime("0?:3?"))
	fmt.Println(maximumTime("1?:22"))
	fmt.Println(maximumTime("??:??"))
}
