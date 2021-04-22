package main

import (
	"fmt"
	"sort"
)

func testBasicHeap() {
	hp := arrayToBasicHeap([]int{9, 6, 5, 2, 3})

	fmt.Println(hp.Empty())

	fmt.Println(hp)

	for !hp.Empty() {
		fmt.Println(hp.Remove())
	}

	arr := []int{6, 5, 3, 7, 2, 8}

	for _, ele := range arr {
		hp.Insert(ele)
		fmt.Println(ele, hp)
	}
}

type basicHeap struct {
	arr []int
}

func (hp *basicHeap) Empty() bool {
	return len(hp.arr) == 0
}

func (hp *basicHeap) Size() int {
	return len(hp.arr)
}

func (hp *basicHeap) Insert(item int) {
	hp.arr = append(hp.arr, item)

	j := len(hp.arr) - 2
	for ; j >= 0 && hp.arr[j] < item; j-- {
		hp.arr[j+1] = hp.arr[j]
	}
	hp.arr[j+1] = item
}

func (hp *basicHeap) Remove() int {
	top := hp.arr[0]
	hp.arr = hp.arr[1:]
	return top
}

func arrayToBasicHeap(arr []int) *basicHeap {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	hp := &basicHeap{arr}
	return hp
}
