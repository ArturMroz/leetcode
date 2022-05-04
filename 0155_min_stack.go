// https://leetcode.com/problems/min-stack/
// Design a stack that supports push, pop, top, and retrieving the minimum
// element in constant time.

package main

type Element struct {
	val, min int
}

type MinStack struct {
	vals []Element
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	el := Element{val: val}
	if len(ms.vals) == 0 || val < ms.vals[len(ms.vals)-1].min {
		el.min = val
	} else {
		el.min = ms.vals[len(ms.vals)-1].min
	}

	ms.vals = append(ms.vals, el)
}

func (ms *MinStack) Pop() {
	ms.vals = ms.vals[:len(ms.vals)-1]
}

func (ms *MinStack) Top() int {
	return ms.vals[len(ms.vals)-1].val
}

func (ms *MinStack) GetMin() int {
	return ms.vals[len(ms.vals)-1].min
}
