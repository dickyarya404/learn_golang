package main

import "fmt"

// Struktur node
type Node struct {
	value int
	next  *Node
}

// Cetak semua node
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.value, " -> ")
		head = head.next
	}
	fmt.Println("nil")
}

// Menambahkan node baru di akhir
func appendNode(head *Node, val int) {
	newNode := &Node{value: val}
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func main() {
	// Buat node awal
	head := &Node{value: 1}
	head.next = &Node{value: 2}
	head.next.next = &Node{value: 3}

	// Tambah node baru di akhir
	appendNode(head, 4)
	appendNode(head, 5)

	// Cetak hasil akhir
	printList(head)
}
