package main

import "fmt"

func testStack() {
	st := &intStack{}
	st.Push(1)
	st.Push(2)
	st.Push(3)

	fmt.Println(st)

	fmt.Println(st.Pop())

	fmt.Println(st)

	st.Push(3)

	fmt.Println(st)
}

type intStack struct {
	arr []int
}

func (stk *intStack) Push(a int) {
	stk.arr = append(stk.arr, a)
}

func (stk *intStack) Pop() int {
	n := len(stk.arr)
	a := stk.arr[n-1]
	stk.arr = stk.arr[:n-1]
	return a
}

func (stk *intStack) Empty() bool {
	return len(stk.arr) == 0
}

type stringStack struct {
	arr []string
}

func (stk *stringStack) Push(a string) {
	stk.arr = append(stk.arr, a)
}

func (stk *stringStack) Pop() string {
	n := len(stk.arr)
	a := stk.arr[n-1]
	stk.arr = stk.arr[:n-1]
	return a
}

func (stk *stringStack) Empty() bool {
	return len(stk.arr) == 0
}
