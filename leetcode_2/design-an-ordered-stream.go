package main

type OrderedStream struct {
	arr []string
	n   int
	ptr int
}

func Constructor(n int) OrderedStream {
	obj := &OrderedStream{n: n}
	obj.arr = make([]string, n)
	return *obj
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.arr[idKey-1] = value

	i := this.ptr
	j := this.ptr
	for i := this.ptr; i < this.n; i++ {
		if this.arr[i] != "" {
			this.ptr++
			j++
		} else {
			break
		}
	}

	return this.arr[i:j]
}
