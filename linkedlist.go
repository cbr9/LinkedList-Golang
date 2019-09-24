package main

import "fmt"

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewDoublyLinkedList(elements ...interface{}) *DoublyLinkedList {
	list := new(DoublyLinkedList)
	for _, elem := range elements {
		list.AddToEnd(elem)
	}
	return list
}

func (list *DoublyLinkedList) AddToFront(element interface{}) {
	node := NewNode(element)
	if list.IsEmpty() {
		list.tail = node
	} else {
		list.head.previous = node
		node.next = list.head
	}
	list.length++
	list.head = node
}

func (list *DoublyLinkedList) RemoveFromFront() *Node {
	if list.IsEmpty() {
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

func (list *DoublyLinkedList) AddToEnd(element interface{}) {
	node := NewNode(element)
	if list.IsEmpty() {
		list.head = node
	} else {
		list.tail.next = node
		node.previous = list.tail
	}
	list.length++
	list.tail = node
}

func (list *DoublyLinkedList) RemoveFromEnd() *Node {
	if list.IsEmpty() {
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

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.head == nil && list.tail == nil
}

func (list *DoublyLinkedList) ToString() {
	if list.IsEmpty() {
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

func (list *DoublyLinkedList) SearchFor(element interface{}) (index int,  node *Node) {
	node = list.head
	for node.element != element {
		node = node.next
		index++
	}
	return index, node
}
