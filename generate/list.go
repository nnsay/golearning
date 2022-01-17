package main

import "fmt"

type Node struct {
	Data int
	next *Node
}

type List struct {
	header *Node
}

func (l *List) add(data int) {
	node := Node{
		Data: data,
		next: nil,
	}
	if l.header == nil {
		l.header = &node
		return
	}
	if l.header.next == nil {
		l.header.next = &node
		return
	}
	h := l.header
	l.header = l.header.next // 移动header向后
	l.add(data)
	l.header = h
}

func main() {
	var l List
	fmt.Println(l)
	l.add(12) // header 12 header next nil
	l.add(9)  // header 12 header next 9
	l.add(15) // header 12
	l.add(6)
	for {
		fmt.Println("*", l.header)
		if l.header.next == nil {
			break
		}
		l.header = l.header.next
	}
}
