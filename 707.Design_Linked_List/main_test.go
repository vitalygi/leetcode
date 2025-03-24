package main

import "testing"

type Node struct {
	Value int
	Next  *Node
}

type MyLinkedList struct {
	Head *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	currentNode := this.Head
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}
	return currentNode.Value
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size || index < 0 {
		return
	}
	this.size++
	if index == 0 {
		this.Head = &Node{
			Value: val,
			Next:  this.Head,
		}
		return
	}
	currentNode := this.Head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}
	oldNode := currentNode.Next
	currentNode.Next = &Node{
		Value: val,
		Next:  oldNode,
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size || index < 0 {
		return
	}
	this.size--
	if index == 0 {
		this.Head = this.Head.Next
		return
	}
	currentNode := this.Head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}
	currentNode.Next = currentNode.Next.Next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func TestGet(t *testing.T) {
	list := MyLinkedList{
		size: 3,
		Head: &Node{
			Value: 1,
			Next: &Node{
				Value: 2,
				Next: &Node{
					Value: 3,
					Next:  nil,
				},
			},
		},
	}
	tests := []struct {
		name     string
		list     MyLinkedList
		index    int
		expected int
	}{
		{
			name:     "basic test",
			list:     list,
			index:    0,
			expected: 1,
		},
		{
			name:     "basic test",
			list:     list,
			index:    1,
			expected: 2,
		},
		{
			name:     "basic test",
			list:     list,
			index:    2,
			expected: 3,
		},
		{
			name:     "wrong index",
			list:     list,
			index:    4,
			expected: -1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.list.Get(test.index)
			if result != test.expected {
				t.Errorf("Get failed. Expected %v, got %v", test.expected, result)
			}
		})
	}

}
