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

func getNodeByIndex(node *ListNode, index int) *ListNode {
	for index > 0 {
		if node == nil {
			return nil
		}
		node = node.next
		index--
	}
	return node
}

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		return head.next
	}
	previousNode := getNodeByIndex(head, idx-1)
	if previousNode == nil {
		return head
	}

	del := previousNode.next
	previousNode.next = del.next

	return head
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	newHead := Solution(&node0, 1)
	for newHead != nil {
		fmt.Println(newHead.data)
		newHead = newHead.next
	}
	// result is : node0 -> node2 -> node3
}
