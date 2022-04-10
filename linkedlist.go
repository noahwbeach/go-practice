package main

import "fmt"

type node struct {
	value string
	next  *node
}

type LinkedList struct {
	head *node
	tail *node
}

func (l LinkedList) Empty() bool {
	if l.head == nil {
		return true
	}
	return false
}

func (l *LinkedList) PushFront(value string) *LinkedList {
	new_node := node{value, nil}

	if l.head == nil {
		l.tail = &new_node
	} else {
		new_node.next = l.head
	}
	l.head = &new_node

	return l
}

func (l *LinkedList) PopFront() string {
	if l.head == nil {
		return ""
	}
	value := l.head.value
	l.head = l.head.next

	return value
}

func (l LinkedList) ValueAt(index int) string {
	n := l.head
	for i := 0; i <= index; i++ {
		if i == index {
			return n.value
		} else {
			n = n.next
		}
	}
	panic("ERROR: index out of range of LinkedList")
}

func (l *LinkedList) Reverse() *LinkedList {
	fwd := l.head.next
	back := &node{}
	l.head.next = nil
	l.tail = l.head

	for fwd != nil {
		back = l.head
		l.head = fwd
		fwd = fwd.next

		l.head.next = back
	}

	return l
}

func (l LinkedList) PrintAll() {
	list_str := ""

	if l.head != nil {
		list_str += l.head.value
		l.head = l.head.next

		for l.head != nil {
			list_str += ", " + l.head.value
			l.head = l.head.next
		}

		fmt.Println(list_str)
	} else {
		fmt.Println("")
	}
}
