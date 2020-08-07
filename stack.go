package main

import (
	"fmt"
)

type Data struct {
	arr  []int
	size int
	top  int
}

func (d *Data) Show() []int {
	return d.arr
}

func (d *Data) Push(e int) (string, error) {
	if d.size == d.top {
		return "", fmt.Errorf("Stack is at maximum capacity")
	}
	d.top++
	d.arr[d.top] = e
	return "An element was added", nil
}

func (d *Data) Pop() (string, error) {
	if d.top == -1 {
		return "", fmt.Errorf("Stack are empty")
	}
	d.arr[d.top] = 0
	d.top--
	return "Stack top element was deleted", nil
}

type StackInterface interface {
	Show() []int

	Push(int) (string, error)

	Pop() (string, error)
}

func NewData(s int) StackInterface {
	d := &Data{}
	d.arr = make([]int, s)
	d.top = -1
	d.size = int(s) - 1
	return d
}

func main() {

	// s := 2
	// d := &Data{}
	// d.arr = make([]int, s)
	// d.top = -1
	// d.size = int(s) - 1

	d := NewData(2)

	m1, e1 := d.Push(1)
	fmt.Println("r: ", m1, e1, d.Show())
	m2, e2 := d.Push(2)
	fmt.Println("r: ", m2, e2, d.Show())
	m3, e3 := d.Push(3)
	fmt.Println("r: ", m3, e3, d.Show())
	m4, e4 := d.Pop()
	fmt.Println("r: ", m4, e4, d.Show())
	m5, e5 := d.Pop()
	fmt.Println("r: ", m5, e5, d.Show())
	m6, e6 := d.Pop()
	fmt.Println("r: ", m6, e6, d.Show())
}
