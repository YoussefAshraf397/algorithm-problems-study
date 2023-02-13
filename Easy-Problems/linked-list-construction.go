package main

import "fmt"

type NodeLL struct {
	Value      int
	Prev, Next *NodeLL
}
type DoublyLinkedList struct {
	Head, Tail *NodeLL
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) insertBefore(node, nodeToInsert *NodeLL) {
	if nodeToInsert == ll.Tail && nodeToInsert == ll.Head {
		return
	}
	ll.remove(nodeToInsert)
	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node

	if node.Prev == nil {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}
	node.Prev = nodeToInsert
}

func (ll *DoublyLinkedList) insertAfter(node, nodeToInsert *NodeLL) {
	if nodeToInsert == ll.Tail && nodeToInsert == ll.Head {
		return
	}
	ll.remove(nodeToInsert)
	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next

	if node.Next == nil {
		ll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}
	node.Next = nodeToInsert
}

// SetHead O(1) time | O(1) space
func (ll *DoublyLinkedList) SetHead(node *NodeLL) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.insertBefore(ll.Head, node)
}

// SetTail O(1) time | O(1) space
func (ll *DoublyLinkedList) SetTail(node *NodeLL) {
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}
}

// remove O(1) time | O(1) space
func (ll *DoublyLinkedList) removeWithValue(value int) {
	node := ll.Head
	for node != nil {
		nodeToRemove := node
		node = node.Next
		if nodeToRemove.Value == value {
			ll.remove(nodeToRemove)
		}
	}
}

// remove O(1) time | O(1) space
func (ll *DoublyLinkedList) remove(node *NodeLL) {
	if ll.Head == node {
		ll.Head = ll.Head.Next
	}
	if ll.Tail == node {
		ll.Tail = ll.Head.Prev
	}
	removeNodeBindings(node)
}
func removeNodeBindings(node *NodeLL) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	node.Prev = nil
	node.Next = nil
}

func (ll *DoublyLinkedList) containsNodeWithValue(value int) bool {
	node := ll.Head
	for node != nil && node.Value != value {
		node = node.Next
	}
	return node != nil
}

func (ll *DoublyLinkedList) insertAtPosition(position int, nodeToInsert *NodeLL) {
	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}
	node := ll.Head
	currentPosition := 1
	for node != nil && position != currentPosition {
		node = node.Next
		currentPosition = currentPosition + 1
	}
	if node != nil {
		ll.insertBefore(node, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

func (ll *DoublyLinkedList) printList() {
	node := ll.Head
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}

func main() {
	node1 := NodeLL{1, nil, nil}
	node2 := NodeLL{3, nil, nil}
	node3 := NodeLL{5, nil, nil}

	linkdeList := DoublyLinkedList{&node1, &node1}
	linkdeList.insertAfter(&node1, &node2)
	linkdeList.insertBefore(&node1, &node3)
	linkdeList.remove(&node1)
	linkdeList.removeWithValue(3)

	linkdeList.printList()

}
