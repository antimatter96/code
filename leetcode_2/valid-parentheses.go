package main

import "fmt"

type stackRune struct {
	arr []rune
}

func (stk *stackRune) Pop() (rune, error) {
	n := len(stk.arr)
	if n < 1 {
		return '-', fmt.Errorf("empty")
	}
	top := stk.arr[n-1]

	stk.arr = stk.arr[:n-1]

	return top, nil
}

func (stk *stackRune) Push(r rune) {
	if stk.arr == nil {
		stk.arr = make([]rune, 0, 1)
	}

	stk.arr = append(stk.arr, r)
}

func (stk *stackRune) Top() rune {
	n := len(stk.arr)

	return stk.arr[n-1]
}

func (stk *stackRune) Empty() bool {
	return len(stk.arr) == 0
}

func isValid(s string) bool {
	stk := &stackRune{}

	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stk.Push(r)
		} else {
			top, _ := stk.Pop()

			if r == ')' && top == '(' {

			} else if r == '}' && top == '{' {

			} else if r == ']' && top == '[' {

			} else {
				return false
			}
		}
	}

	return stk.Empty()
}

func driver__isValid() {
	test_cases := []struct {
		input  string
		answer bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, tc := range test_cases {
		fmt.Println(isValid(tc.input) == tc.answer)
	}
}
