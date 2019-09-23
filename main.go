package main

import "fmt"

func main() {
	list := NewDoublyLinkedList()
	newElements := []interface{}{"John", "Mary", 54, 44, 55.1}
	for _, elem := range newElements {
		list.AddToEnd(elem)
	}
	list2 := NewDoublyLinkedList("John", "Mary", 54, 44, 55.1)
	list2.ToString()
	index, mary := list.SearchFor("Mary")
	fmt.Println(index, mary)
}
