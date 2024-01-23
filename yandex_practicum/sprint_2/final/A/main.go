package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	ErrorNotEnoughSpace = errors.New("error")
	ErrorEmptyDequeue   = errors.New("error")
)

const (
	PushFront        string = "push_front"
	PushBack         string = "push_back"
	PopFront         string = "pop_front"
	PopBack          string = "pop_back"
	ForwardDirection int    = 1
	ReverseDirection int    = -1
)

type Deque struct {
	buf  []int
	maxN int
	head int
	tail int
	size int
}

func NewDeque(n int) *Deque {
	return &Deque{
		buf:  make([]int, n),
		maxN: n,
		head: 1,
		tail: 0,
		size: 0,
	}
}

func main() {
	reader := makeReader()
	writer := makeWriter()

	n := readNum(reader)
	m := readNum(reader)
	tasks := readArrStr(n, reader)

	result := Solution(m, tasks)

	printResult(writer, result)
}

func Solution(m int, tasks []string) []string {
	result := make([]string, 0, len(tasks))
	dq := NewDeque(m)

	for _, task := range tasks {
		var command string
		var input int
		arr := strings.Split(task, " ")

		if len(arr) == 2 {
			input, _ = strconv.Atoi(arr[1])
		}

		command = arr[0]

		switch command {
		case PushFront:
			err := dq.PushFront(input)
			if output, ok := checkErr(err); ok {
				result = append(result, output)
			}
		case PopFront:
			output, err := dq.PopFront()
			result = append(result, checkResult(output, err))
		case PushBack:
			err := dq.PushBack(input)
			if output, ok := checkErr(err); ok {
				result = append(result, output)
			}
		case PopBack:
			output, err := dq.PopBack()
			result = append(result, checkResult(output, err))
		default:
			continue
		}
	}

	return result
}

func (dq *Deque) PushBack(item int) error {
	if dq.size != dq.maxN {
		dq.tail = dq.GetNextIndex(ForwardDirection, dq.tail)
		dq.buf[dq.tail] = item
		dq.size += 1

		return nil
	}

	return ErrorNotEnoughSpace
}

func (dq *Deque) PushFront(item int) error {
	if dq.size != dq.maxN {
		dq.head = dq.GetNextIndex(ReverseDirection, dq.head)
		dq.buf[dq.head] = item
		dq.size += 1

		return nil
	}

	return ErrorNotEnoughSpace
}

func (dq *Deque) PopBack() (int, error) {
	if ok := dq.IsEmpty(); ok {
		return 0, ErrorEmptyDequeue
	}

	item := dq.buf[dq.tail]
	dq.buf[dq.tail] = 0
	dq.tail = dq.GetNextIndex(ReverseDirection, dq.tail)
	dq.size -= 1

	return item, nil
}

func (dq *Deque) PopFront() (int, error) {
	if ok := dq.IsEmpty(); ok {
		return 0, ErrorEmptyDequeue
	}

	item := dq.buf[dq.head]
	dq.buf[dq.head] = 0
	dq.head = dq.GetNextIndex(ForwardDirection, dq.head)
	dq.size -= 1

	return item, nil
}

func (dq *Deque) IsEmpty() bool {
	return dq.size == 0
}

func (dq *Deque) GetNextIndex(direction, currentIndex int) int {
	index := 0

	switch direction {
	case ForwardDirection:
		index = (currentIndex + 1) % dq.maxN
	case ReverseDirection:
		index = (((currentIndex - 1) % dq.maxN) + dq.maxN) % dq.maxN
	}

	return index
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

func readArrStr(m int, reader *bufio.Reader) []string {
	strArr := make([]string, 0, m)

	for i := m; i > 0; i-- {
		str, _ := reader.ReadString('\n')

		strArr = append(strArr, strings.TrimSpace(str))
	}

	return strArr
}

func makeWriter() *bufio.Writer {
	writer := bufio.NewWriter(os.Stdout)

	return writer
}

func printResult(writer *bufio.Writer, result []string) {
	for _, val := range result {
		writer.WriteString(val + "\n")
	}
	writer.Flush()
}

func checkResult(x int, err error) string {
	if err != nil {
		return err.Error()
	}

	return strconv.Itoa(x)
}

func checkErr(err error) (string, bool) {
	if err != nil {
		return err.Error(), true
	}

	return "", false
}
