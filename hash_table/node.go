package hash

import (
	"fmt"
)

type node struct {
	key   string
	value int
	next  *node
}

func lastNode(node *node) *node {
	if node.next == nil {
		return node
	}
	return lastNode(node.next)
}

func linkedLookup(node *node, key string) *node {
	if node == nil {
		return nil
	}
	if node.key == key {
		return node
	}
	return linkedLookup(node.next, key)
}

func linkedRemoveNode(node *node, key string) *node {
	head := node
	fmt.Println(head.key)
	if node.key == key {
		fmt.Println(head.next)
		return head.next
	}
	if node.next.key == key {
		head.next = node.next.next
		return head
	}
	return linkedRemoveNode(head.next, key)

}

func lenNode(node *node) int{
	if node == nil {
		return 0
	}
	return 1 + lenNode(node.next)
}

func lookup(node *node) {
	if node == nil {
		return
	}
	printer(node)
	lookup(node.next)
}

func printer(node *node) {
	fmt.Println("key: ", node.key)
	fmt.Println("value: ", node.value)
}
