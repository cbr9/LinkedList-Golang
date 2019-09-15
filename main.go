package main

import "fmt"

func main() {
	list := newDoublyLinkedList()
	newElements := []interface{}{"John", "Mary", 54, 44, 55.1}
	for _, elem := range newElements {
		list.addToEnd(elem)
	}
	list2 := newDoublyLinkedList("John", "Mary", 54, 44, 55.1)
	list2.toString()
	index, mary := list.searchFor("Mary")
	fmt.Println(index, mary)
}
