package main

import "fmt"

func sortArrayByParity(A []int) []int {
	l := 0
	m := 0

	for l < len(A) && m < len(A) {
		for l < len(A) && A[l]%2 == 0 {
			l++
		}

		m = l + 1

		for m < len(A) && A[m]%2 == 1 {
			m++
		}

		if l < len(A) && m < len(A) && A[l]%2 == 1 && A[m]%2 == 0 {
			A[l], A[m] = A[m], A[l]
		}
	}

	return A
}

func sortArrayByParity2(A []int) []int {
	fmt.Println("=========")
	arr := make([]int, len(A))
	copy(arr, A)

	l := 0
	j := len(A) - 1
	for _, ele := range A {
		if ele%2 == 0 {
			arr[l] = ele
			l++
		} else {
			arr[j] = ele
			j--
		}
		fmt.Println(arr)
	}

	return arr
}

func sortArrayByParity3(A []int) []int {
	odd := make([]int, 0, len(A))
	even := make([]int, 0, len(A))

	for _, ele := range A {
		if ele%2 == 0 {
			even = append(even, ele)
		} else {
			odd = append(odd, ele)
		}
	}

	even = append(even, odd...)
	return even
}

func sortArrayByParity4(A []int) []int {
	i := 0
	j := len(A) - 1
	for i < j {
		if A[i]%2 == 1 && A[j]%2 == 0 {
			A[i], A[j] = A[j], A[i]
		}
		if A[i]%2 == 0 {
			i++
		}
		if A[j]%2 == 1 {
			j--
		}
	}

	return A
}

func driver__sortArrayByParity() {
	fmt.Println(sortArrayByParity([]int{1, 2, 3, 4}))
	fmt.Println(sortArrayByParity([]int{1, 2, 4, 3}))
	fmt.Println(sortArrayByParity([]int{1, 3, 2, 4}))
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4, 5}))
	fmt.Println(sortArrayByParity([]int{3, 1, 3, 4, 5}))
	fmt.Println(sortArrayByParity([]int{0, 2, 1}))
	fmt.Println(sortArrayByParity([]int{0, 1, 2}))
	fmt.Println(sortArrayByParity([]int{1, 0, 3, 5}))
	fmt.Println(sortArrayByParity([]int{0, 2, 1, 4}))

	fmt.Println("<=< <==> >=>")
	fmt.Println(sortArrayByParity2([]int{1, 2, 3, 4}))
	fmt.Println(sortArrayByParity2([]int{1, 2, 4, 3}))
	fmt.Println(sortArrayByParity2([]int{1, 3, 2, 4}))
	fmt.Println(sortArrayByParity2([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity2([]int{3, 1, 2, 4, 5}))
	fmt.Println(sortArrayByParity2([]int{3, 1, 3, 4, 5}))
	fmt.Println(sortArrayByParity2([]int{0, 2, 1}))
	fmt.Println(sortArrayByParity2([]int{0, 1, 2}))
	fmt.Println(sortArrayByParity2([]int{1, 0, 3, 5}))
	fmt.Println(sortArrayByParity2([]int{0, 2, 1, 4}))

	fmt.Println("<=< <==> >=>")
	fmt.Println(sortArrayByParity3([]int{1, 2, 3, 4}))
	fmt.Println(sortArrayByParity3([]int{1, 2, 4, 3}))
	fmt.Println(sortArrayByParity3([]int{1, 3, 2, 4}))
	fmt.Println(sortArrayByParity3([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity3([]int{3, 1, 2, 4, 5}))
	fmt.Println(sortArrayByParity3([]int{3, 1, 3, 4, 5}))
	fmt.Println(sortArrayByParity3([]int{0, 2, 1}))
	fmt.Println(sortArrayByParity3([]int{0, 1, 2}))
	fmt.Println(sortArrayByParity3([]int{1, 0, 3, 5}))
	fmt.Println(sortArrayByParity3([]int{0, 2, 1, 4}))
}
