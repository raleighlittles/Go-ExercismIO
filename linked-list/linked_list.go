package linkedlist

import "errors"

var ErrEmptyList = errors.New("Empty list")

type Node struct {
	Val interface{}
	prev  *Node
	next *Node
}

type List struct {
	head *Node
	tail *Node
	length int
}

func (node *Node) Next() (*Node) {
	return node.next
}

func (node *Node) Prev() (*Node) {
	return node.prev
}

func NewList(args... interface{}) (l *List) {
	l = &List{nil, nil, 0}
	for _, nodeValue := range args {
		l.PushBack(nodeValue)
	}

	return
}

func (list *List) PushFront(v interface{}) {

	if (list.isEmpty()) {
		newNode := &Node{v, nil, nil}
		list.head = newNode
		list.tail = newNode
	} else {
		newNode := &Node{v, nil, list.head}
		list.head.prev = newNode
		list.head = newNode

		}

	list.length++

}

func (list *List) PushBack(v interface{}) {
	var newNode *Node

	if (list.isEmpty()) {
		newNode = &Node{v, nil, nil}
		list.head = newNode
		list.tail = newNode
	} else {
		newNode = &Node{v, list.tail, nil}
		list.tail.next = newNode
		list.tail = newNode
	}

	list.length++
}

func (list *List) PopFront() (value interface{}, err error) {
	if (list.isEmpty()) {
		value = nil
		err = ErrEmptyList
	} else {
		value = list.head.Val
		err = nil

		if (list.length == 1) {
			list.head = nil
			list.tail = nil
		} else {
			newHeadNode := list.head.next
			newHeadNode.prev = nil
			list.head = newHeadNode
		}

		list.length--
	}

	return
}

func (list *List) PopBack() (value interface{}, err error) {

	if (list.isEmpty()) {
		value = nil
		err = ErrEmptyList
	} else {

		value = list.tail.Val
		err = nil


		if (list.length == 1) {
			// Now, the list is empty -- you popped the last element off it
			list.head = nil
			list.tail = nil

		} else {
			 newLastNode := list.tail.prev
			 newLastNode.next = nil
			 list.tail = newLastNode
		}

		list.length--

	}

	return
}

// Note: The return parameter listed for this function is wrong! It is supposed to modify the list
// in place; not return a new list as the given method signature would lead you to expect.

func (list *List) Reverse()  {
	if (list.length > 1) {
		currentNode := list.head
		list.tail = currentNode

		for (currentNode != nil) {
			temporaryNode := currentNode.prev
			// swap the previous and next pointers as you iterate through the list
			currentNode.prev = currentNode.next
			currentNode.next = temporaryNode

			// now, to go forward in the original list you have to go backwards in the new list

			if( currentNode.prev == nil) {
				// if by going backwards in this new list (aka forwards in the old list) by another node,
				// you'd reach a null node -- then that means you're at the end of the original list
				// which should now be the start of the original list
				list.head = currentNode
				break
			}

			currentNode = currentNode.prev
		}
	}
}

func (list *List) First() (*Node) {
	return list.head
}

func (list *List) Last() (*Node) {
	return list.tail
}

func (list *List) isEmpty() (bool) {
	return list.length == 0
}
