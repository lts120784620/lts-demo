package main

import "lts-demo/test/linkedList"

func main() {
	l := linkedList.New()
	l.Insert(1, &linkedList.Node{"a", nil})
	l.Insert(2, &linkedList.Node{"b", nil})
	l.Travelse()

	head := linkedList.ReverseList(l)
	linkedList.TravelseHead(head)
}
