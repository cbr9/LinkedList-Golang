package main

import "fmt"

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func newDoublyLinkedList(elements ...interface{}) *DoublyLinkedList {
	list := new(DoublyLinkedList)
	for _, elem := range elements {
		list.addToEnd(elem)
	}
	return list
}

func (list *DoublyLinkedList) addToFront(element interface{}) {
	node := newNode(element)
	if list.isEmpty() {
		list.tail = node
	} else {
		list.head.previous = node
		node.next = list.head
	}
	list.length++
	list.head = node
}

func (list *DoublyLinkedList) removeFromFront() *Node {
	if list.isEmpty() {
		return nil
	}
	if list.head.next == nil {
		list.tail = nil
	} else {
		list.head.next.previous = nil
	}
	removedNode := list.head
	list.head = list.head.next
	list.length--
	removedNode.next = nil
	return removedNode
}

func (list *DoublyLinkedList) addToEnd(element interface{}) {
	node := newNode(element)
	if list.isEmpty() {
		list.head = node
	} else {
		list.tail.next = node
		node.previous = list.tail
	}
	list.length++
	list.tail = node
}

func (list *DoublyLinkedList) removeFromEnd() *Node {
	if list.isEmpty() {
		return nil
	}
	if list.tail.previous == nil {
		list.head = nil
	} else {
		list.tail.previous.next = nil
	}
	list.length--
	removedNode := list.tail
	list.tail = list.tail.previous
	removedNode.previous = nil
	return removedNode
}

func (list *DoublyLinkedList) isEmpty() bool {
	return list.head == nil && list.tail == nil
}

func (list *DoublyLinkedList) toString() {
	if list.isEmpty() {
		fmt.Println("Linked List is empty.")
		return
	}
	current := list.head
	for current != nil {
		fmt.Print(current.element, " <=> ")
		current = current.next
	}
	fmt.Println("nil")
}

func (list *DoublyLinkedList) searchFor(element interface{}) (index int,  node *Node) {
	node = list.head
	for node.element != element {
		node = node.next
		index++
	}
	return index, node
}
