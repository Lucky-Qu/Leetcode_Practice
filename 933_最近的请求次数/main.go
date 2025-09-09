package main

import "fmt"

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)

	// 移除所有过期请求
	for len(this.requests) > 0 && this.requests[0] < t-3000 {
		this.requests = this.requests[1:]
	}

	return len(this.requests)
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))    // 1
	fmt.Println(obj.Ping(100))  // 2
	fmt.Println(obj.Ping(3001)) // 3
	fmt.Println(obj.Ping(3002)) // 3
}
