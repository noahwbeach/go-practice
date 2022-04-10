package main

func main() {
	l := LinkedList{}
	l.PushFront("there").PushFront("Hello").PushFront("Your mom")

	l.Reverse()
	l.PrintAll()
}
