package main

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head  *Node
	count int
}

func (l *LinkedList) Insert(value int) {
	newNode := Node{}
	newNode.value = value
	if l.count == 0 {
		l.head = &newNode
		l.count++
		return
	}
	ptr := l.head
	for i := 0; i < l.count; i++ {
		if ptr.next == nil {
			ptr.next = &newNode
			l.count++
			return
		}
		ptr = ptr.next
	}
}
