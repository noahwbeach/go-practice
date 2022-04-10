package main

import "fmt"

func main() {
	l := LinkedList{}
	l.PushFront("there").PushFront("Hello").PushFront("Your mom")

	l.Reverse()
	fmt.Println(l.PopFront())
	l.PrintAll()
}
