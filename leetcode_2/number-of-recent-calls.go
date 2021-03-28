package main

import "fmt"

type RecentCounter struct {
	arr []int
}

func RecentCounterConstructor() RecentCounter {
	obj := &RecentCounter{}
	obj.arr = make([]int, 0)
	return *obj
}

// func (this *RecentCounter) Ping(t int) int {
// 	this.arr = append(this.arr, t)

// 	min := t - 3000

// 	for i := this.minCounter; i < len(this.arr); i++ {
// 		if this.arr[i] < min {
// 			continue
// 		}

// 		if this.arr[i] >= min {
// 			this.minCounter = i
// 			break
// 		}

// 	}

// 	return len(this.arr) - this.minCounter
// }

func (this *RecentCounter) Ping2(t int) int {
	this.arr = append(this.arr, t)

	min := t - 3000
	counter := -1

	for i := 0; i < len(this.arr); i++ {
		if this.arr[i] < min {
			continue
		}

		if this.arr[i] >= min {
			counter = i
			break
		}

	}

	if counter != -1 {
		this.arr = this.arr[counter:]
	}

	return len(this.arr)
}

func numberOfRecentCalls() {
	obj := RecentCounterConstructor()
	fmt.Println(obj.Ping2(642))
	fmt.Println(obj.Ping2(1849))
	fmt.Println(obj.Ping2(4921))
	fmt.Println(obj.Ping2(5936))
	fmt.Println(obj.Ping2(5957))
}
