package main

import "fmt"

func testQueue() {
	st := &intQueue{}
	st.Enqueue(1)
	st.Enqueue(2)
	st.Enqueue(3)

	fmt.Println(st)

	fmt.Println(st.Dequeue())

	fmt.Println(st)

	st.Enqueue(3)

	fmt.Println(st)
}

type intQueue struct {
	arr []int
}

func (stk *intQueue) Enqueue(a int) {
	stk.arr = append(stk.arr, a)
}

func (stk *intQueue) Dequeue() int {
	a := stk.arr[0]
	stk.arr = stk.arr[1:]
	return a
}

func (stk *intQueue) Empty() bool {
	return len(stk.arr) == 0
}

type stringQueue struct {
	arr []string
}

func (stk *stringQueue) Enqueue(a string) {
	stk.arr = append(stk.arr, a)
}

func (stk *stringQueue) Dequeue() string {
	a := stk.arr[0]
	stk.arr = stk.arr[1:]
	return a
}

func (stk *stringQueue) Empty() bool {
	return len(stk.arr) == 0
}
