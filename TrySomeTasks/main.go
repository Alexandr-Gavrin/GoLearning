package main

import (
	"bufio"
	"fmt"
	"os"
)

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	next *Node
	data int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, size: 0}
}

func NewNode(data int) *Node {
	return &Node{data: data, next: nil}
}

func (l *LinkedList) Add(data int) {
	tmp := NewNode(data)
	if l.head == nil {
		l.head = tmp
		l.tail = tmp
	} else {
		l.tail.next = tmp
		l.tail = tmp
	}
	l.size++
}

func (l *LinkedList) Print() {
	tmp := l.head
	for tmp != nil {
		fmt.Printf("%d ", tmp.data)
		tmp = tmp.next
	}
	fmt.Println()
}

func (l *LinkedList) swapZeros() {
	index := 0
	cur := l.head
	var prev *Node = nil
	for cur != nil && index != l.size {
		next := cur.next

		if cur.data == 0 {
			if prev != nil {
				prev.next = cur.next
			} else {
				l.head = cur.next
			}

			l.tail.next = cur
			l.tail = cur
			l.tail.next = nil
			cur = next
		} else if cur.data != 0 {
			prev = cur
			cur = next
		}
		index++

	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func merge(arr []int, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid
	arr1 := make([]int, n1)
	arr2 := make([]int, n2)
	for i := 0; i < n1; i++ {
		arr1[i] = arr[left+i]
	}
	for i := 0; i < n2; i++ {
		arr2[i] = arr[mid+1+i]
	}
	i := 0
	j := 0
	k := left
	for i < n1 && j < n2 {
		if arr1[i] <= arr2[j] {
			arr[k] = arr1[i]
			i++
		} else {
			arr[k] = arr2[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = arr1[i]
		k++
		i++
	}
	for j < n2 {
		arr[k] = arr2[j]
		k++
		j++
	}
}

func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

type Book struct {
	ISBN string
	name string
	year int
}

func main() {

	var reader *bufio.Reader
	var writer *bufio.Writer

	reader = bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	writer.WriteString(line)
	writer.Flush()

	// var n int
	// fmt.Scan(&n)
	// //	var arr = NewLinkedList()
	// var arr = make([]Book, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&(arr[i]))
	// }

}
