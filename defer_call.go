package main

import (
	"fmt"
)

var ch chan int = make(chan int, 1)

func foo(id int) { //id: 这个routine的标号
	ch <- id
}

func main1() {
	var ch2 chan int = nil
	ch2 <- 3
	for i := 0; i < 5; i++ {
		go foo(i)
	}

	// 取出信道中的数据
	for i := 0; i < 5; i++ {
		fmt.Print(<-ch)
	}
	fmt.Println(uint(10) - uint(1))
	fmt.Println(deferCall())
	pase_student()
}

func deferCall() string {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	return "main"
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Println(m)
}
