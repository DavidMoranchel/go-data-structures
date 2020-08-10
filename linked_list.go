package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	size int
	head *Node
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Append(n *Node) {
	if l.size == 0 {
		l.head = n
	} else {
		cNode := l.head
		for cNode.next != nil {
			cNode = cNode.next
		}
		cNode.next = n
	}
	l.size++
}

func (l *LinkedList) InsertAfter(cNode *Node, nNode *Node) (string, error) {
	if l.size == 0 {
		return "", fmt.Errorf("Linked list is empty")
	}

	nNode.next = cNode.next
	cNode.next = nNode
	l.size++
	return fmt.Sprintf("Node added after %v", cNode.value), nil
}

func (l *LinkedList) Remove(n *Node) (string, error) {
	if l.size == 0 {
		return "", fmt.Errorf("Linked list is empty")
	}
	var pNode *Node
	cNode := l.head
	for cNode.value != n.value {
		pNode = cNode
		cNode = cNode.next
	}

	pNode.next = cNode.next
	l.size--
	return fmt.Sprintf("Node %v removed", n.value), nil

}

func (l *LinkedList) Show() (string, error) {
	if l.size == 0 {
		return "", fmt.Errorf("Linked list is empty")
	}
	list := make([]int, l.size)
	cNode := l.head
	for i := 0; i < l.size; i++ {
		list[i] = cNode.value
		cNode = cNode.next
	}
	return fmt.Sprintln(list), nil
}

func main() {

	linkedList := &LinkedList{}
	linkedList.size = 0
	// size 0
	fmt.Println(linkedList.Size())
	// Show
	_, se1 := linkedList.Show()
	fmt.Println("ERROR", se1)

	n1 := &Node{}
	n1.value = 10
	n2 := &Node{}
	n2.value = 30
	n3 := &Node{}
	n3.value = 20

	// ERROR: InsertAfter in empty list
	_, e1 := linkedList.InsertAfter(n1, n2)
	fmt.Println("ERROR", e1)
	// ERROR: Remove in empty list
	_, e2 := linkedList.Remove(n1)
	fmt.Println("ERROR", e2)

	// append
	linkedList.Append(n1)
	// size 1
	fmt.Println("Size", linkedList.Size())
	// Show
	s1, _ := linkedList.Show()
	fmt.Println("Show", s1)

	// append
	linkedList.Append(n2)
	// size 2
	fmt.Println("Size", linkedList.Size())
	// Show
	s2, _ := linkedList.Show()
	fmt.Println("Show", s2)

	// InsertAfter in empty list
	m2, _ := linkedList.InsertAfter(n1, n3)
	fmt.Println(m2)
	// size 3
	fmt.Println("Size", linkedList.Size())
	// Show
	s3, _ := linkedList.Show()
	fmt.Println("Show", s3)

	// Remove node n2
	r1, _ := linkedList.Remove(n3)
	fmt.Println(r1)
	// size 2
	fmt.Println("Size", linkedList.Size())
	// Show
	s4, _ := linkedList.Show()
	fmt.Println("Show", s4)
}
