package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Push  string = "push"
	Pop   string = "pop"
	Peek  string = "peek"
	Size  string = "size"
	Error string = "error\n"
	None  string = "None\n"
)

type Queue struct {
	queue  []int
	writer *bufio.Writer
	maxN   int
	head   int
	tail   int
	size   int
}

func main() {
	reader := makeReader()

	num := readNum(reader)
	size := readNum(reader)
	tasks := readArrStr(num, reader)

	Solution(size, tasks)
}

func Solution(size int, tasks []string) {
	q := NewQueue(size)

	for _, task := range tasks {
		arr := strings.Split(task, " ")
		item := 0

		if len(arr) == 2 {
			item, _ = strconv.Atoi(arr[1])
			log.Println(item)
		}

		switch arr[0] {
		case Push:
			q.Push(item)
		case Pop:
			q.Pop()
		case Size:
			q.Size()
		case Peek:
			q.Peek()

		default:
			continue
		}
	}
	q.writer.Flush()
}

func NewQueue(n int) *Queue {
	writer := bufio.NewWriter(os.Stdout)
	return &Queue{
		queue:  make([]int, n),
		writer: writer,
		maxN:   n,
		head:   0,
		tail:   0,
		size:   0,
	}
}

func (q *Queue) Push(item int) {
	if q.size != q.maxN {
		q.queue[q.tail] = item
		q.tail = (q.tail + 1) % q.maxN
		q.size += 1
		return
	}
	q.writer.WriteString(Error)
}

func (q *Queue) Pop() {
	if q.size == 0 {
		q.writer.WriteString(None)
		return
	}
	x := q.queue[q.head]
	q.queue[q.head] = 0
	q.head = (q.head + 1) % q.maxN
	q.size -= 1
	q.writer.WriteString(fmt.Sprintf("%d\n", x))
}

func (q *Queue) Peek() {
	if q.size == 0 {
		q.writer.WriteString(None)
		return
	}
	x := q.queue[q.head]
	q.writer.WriteString(fmt.Sprintf("%d\n", x))
}

func (q *Queue) Size() {
	q.writer.WriteString(fmt.Sprintf("%d\n", q.size))
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
