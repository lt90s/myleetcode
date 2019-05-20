package main

import (
	"fmt"
	"strconv"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	single bool
	value interface{}
}
// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return n.single
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	i, _ := n.value.(int)
	return i
}
// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.single = true
	n.value = value
}
// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	if n.single != false {
		panic("should be list")
	}

	if n.value == nil {
		n.value = make([]*NestedInteger, 0, 1)
	}
	list, _ := n.value.([]*NestedInteger)
	n.value = append(list, &elem)
}
// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	if n.single {
		return nil
	} else {
		list, _ := n.value.([]*NestedInteger)
		return list
	}
}

func (n NestedInteger) display() {
	if n.single {
		fmt.Print(n.GetInteger())
	} else {
		fmt.Print("[")
		for i, nn := range n.GetList() {
			if i != 0 {
				fmt.Print(",")
			}
			nn.display()
		}
		fmt.Print("]")
	}
}


func deserialize(s string) *NestedInteger {
	n := &NestedInteger{}

	i := 0
	if s[i] == '[' {
		deserializeList(s, n, i+1)
	} else {
		deserializeInteger(s, n, i)
	}
	return n
}

func deserializeList(s string, n *NestedInteger, i int) int {
	for i < len(s) {
		nn := NestedInteger{}
		if s[i] == '[' {
			i += 1
			i = deserializeList(s, &nn, i)
		} else if s[i] == ']'{
			i += 1
			break
		} else if s[i] == ',' {
			i += 1
			continue
		} else {
			i = deserializeInteger(s, &nn, i)
		}
		n.Add(nn)
	}
	return i
}

func deserializeInteger(s string, n *NestedInteger, i int) int {
	extra := 0
	j := i
	for j < len(s) {
		if s[j] == ',' {
			extra = 1
			break
		}
		if s[j] == ']' {
			break
		}
		j += 1
	}
	i, _ = strconv.Atoi(s[i:j])
	j += extra
	n.SetInteger(i)
	return j
}

func main() {
	s := "[123,[456,[789],[[12,23]]]]"
	n := deserialize(s)
	n.display()
}
