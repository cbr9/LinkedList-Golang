package main

type Node struct {
	previous *Node
	next     *Node
	element  interface{}
}

func NewNode(element interface{}) *Node {
	node := new(Node)
	node.element = element
	return node
}
