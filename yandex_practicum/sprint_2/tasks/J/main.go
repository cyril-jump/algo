package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Get   string = "get"
	Put   string = "put"
	Size  string = "size"
	Error string = "error\n"
)

type Node struct {
	value int
	next  *Node
}

type LinkedQueue struct {
	size   int
	head   *Node
	tail   *Node
	writer *bufio.Writer
}

func NewLinkedQueue() *LinkedQueue {
	writer := bufio.NewWriter(os.Stdout)
	return &LinkedQueue{
		size:   0,
		head:   nil,
		tail:   nil,
		writer: writer,
	}
}

func main() {
	reader := makeReader()
	num := readNum(reader)
	tasks := readArrStr(num, reader)
	Solution(tasks)
}

func (lq *LinkedQueue) Put(item int) {
	newNode := &Node{
		value: item,
		next:  lq.head,
	}

	if lq.size == 0 {
		lq.tail = newNode
		lq.head = newNode

	} else {
		lq.tail.next = newNode
		lq.tail = lq.tail.next
	}

	lq.size++
}

func (lq *LinkedQueue) Get() {
	if lq.size == 0 {
		lq.writer.WriteString(Error)
		return
	}

	tmp := lq.head.value
	lq.head = lq.head.next

	if lq.size == 1 {
		lq.tail = nil
	}

	lq.size--
	lq.writer.WriteString(fmt.Sprintf("%d\n", tmp))
}

func (lq *LinkedQueue) Size() {
	lq.writer.WriteString(fmt.Sprintf("%d\n", lq.size))
}

func makeReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)

	return reader
}

func readNum(reader *bufio.Reader) int {
	num, _, _ := reader.ReadLine()
	intNum, _ := strconv.Atoi(string(num))

	return intNum
}

func readArrStr(num int, reader *bufio.Reader) []string {
	strArr := make([]string, 0, 1)

	for i := num; i > 0; i-- {
		str, _ := reader.ReadString('\n')
		strArr = append(strArr, strings.TrimSpace(str))
	}

	return strArr
}

func Solution(tasks []string) {
	lq := NewLinkedQueue()

	for _, task := range tasks {
		arr := strings.Split(task, " ")
		item := 0

		if len(arr) == 2 {
			item, _ = strconv.Atoi(arr[1])
		}

		switch arr[0] {
		case Put:
			lq.Put(item)
		case Get:
			lq.Get()
		case Size:
			lq.Size()

		default:
			continue
		}
	}
	lq.writer.Flush()
}
