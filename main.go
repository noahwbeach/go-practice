package main

import "fmt"

func main() {
	l := LinkedList{}
	l.PushFront("there").PushFront("Hello")

	l.Reverse()
	fmt.Println(l.PopFront())
	l.PrintAll()
}
