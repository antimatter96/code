package main

import "fmt"

func testHeap() {
	hp := arrayToHeap([]int{9, 6, 5, 2, 3})

	fmt.Println(hp.Empty())

	fmt.Println(hp)

	for !hp.Empty() {
		fmt.Println(hp.remove())
	}
}

type heap struct {
	arr  []int
	size int
}

func (hp *heap) Empty() bool {
	return hp.size == 0
}

func (hp *heap) Size() int {
	return hp.size
}

func (hp *heap) leaf(index int) bool {
	if index >= (hp.size/2) && index <= hp.size {
		return true
	}

	return false
}

func (hp *heap) parent(index int) int {
	return (index - 1) / 2
}

func (hp *heap) insert(item int) {
	hp.arr = append(hp.arr, item)
	hp.size++
	hp.upHeapify(hp.size - 1)
}

func (hp *heap) swap(i, j int) {
	hp.arr[i], hp.arr[j] = hp.arr[j], hp.arr[i]
}

func (hp *heap) upHeapify(index int) {
	for hp.arr[index] > hp.arr[hp.parent(index)] {
		hp.swap(index, hp.parent(index))
		index = hp.parent(index)
	}
}

func (hp *heap) downHeapify(current int) {
	if hp.leaf(current) {
		return
	}
	largest := current
	leftChildIndex := (2 * current) + 1
	rightRightIndex := (2 * current) + 2

	//If current is smallest then return
	if leftChildIndex < hp.size && hp.arr[leftChildIndex] > hp.arr[largest] {
		largest = leftChildIndex
	}
	if rightRightIndex < hp.size && hp.arr[rightRightIndex] > hp.arr[largest] {
		largest = rightRightIndex
	}
	if largest != current {
		hp.swap(current, largest)
		hp.downHeapify(largest)
	}
}

func (hp *heap) remove() int {
	top := hp.arr[0]
	hp.arr[0] = hp.arr[hp.size-1]
	hp.arr = hp.arr[:(hp.size)-1]
	hp.size--
	hp.downHeapify(0)
	return top
}

func arrayToHeap(arr []int) *heap {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	hp := &heap{}
	for i := 0; i < len(inputArray); i++ {
		hp.insert(inputArray[i])
	}

	n := len(inputArray)
	for index := ((n / 2) - 1); index >= 0; index-- {
		hp.downHeapify(index)
	}

	return hp
}
