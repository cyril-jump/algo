package main

import "fmt"

// <template>
type ListNode struct {
	data string
	next *ListNode
}

// <template>

/*func main() {
	test()
}*/

func Solution(head *ListNode, elem string) int {
	idx := 0
	for head != nil {
		if head.data == elem {
			return idx
		}
		head = head.next
		idx++
	}

	return -1
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	idx := Solution(&node0, "node1")
	fmt.Println(idx)
	// result is : node0 -> node2 -> node3
}
