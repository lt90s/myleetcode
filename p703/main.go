package main

	import (
		"container/heap"
		"fmt"
	)

	type intHeap []int

	// first implement sort.Interface
	func (h intHeap) Len() int {
		return len(h)
	}

	func (h intHeap) Less(i, j int) bool {
		return h[i] < h[j]
	}

	func (h intHeap) Swap(i, j int) {
		h[i], h[j] = h[j], h[i]
	}

	// implement Push and Pop

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) KthValue() int {
	return (*h)[0]
}

func (h *intHeap) Pop() interface{} {
	len := len(*h)
	if len == 0 {
		return nil
	}

	value := (*h)[len-1]
	*h = (*h)[:len-1]
	return value
}


type KthLargest struct {
	heap *intHeap
	k int
}


func Constructor(k int, nums []int) KthLargest {
	h := &intHeap{}
	kl :=  KthLargest{
		heap: h,
		k: k,
	}

	for _, num := range nums {
		kl.Add(num)
	}

	return kl
}



func (this *KthLargest) Add(val int) int {
	if len(*this.heap) < this.k {
		heap.Push(this.heap, val)
	} else if val > this.heap.KthValue() {
		heap.Pop(this.heap)
		heap.Push(this.heap, val)
	}
	return this.heap.KthValue()
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */


func main() {
	k := 3
	nums := []int{4,5,8,2}

	obj := Constructor(k, nums)
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(11))
}
