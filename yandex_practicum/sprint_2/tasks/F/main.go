package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Push   string = "push"
	Pop    string = "pop"
	GetMax string = "get_max"
	Error  string = "error\n"
	None   string = "None\n"
)

type Stack struct {
	items  []int
	max    int
	writer *bufio.Writer
}

func NewStack() *Stack {
	writer := bufio.NewWriter(os.Stdout)
	return &Stack{
		items:  []int{},
		max:    0,
		writer: writer,
	}
}

func main() {
	reader := makeReader()

	num := readNum(reader)
	tasks := readArrStr(num, reader)

	Solution(tasks)
}

func Solution(tasks []string) {
	stack := NewStack()

	for _, task := range tasks {
		arr := strings.Split(task, " ")
		item := 0

		if len(arr) == 2 {
			item, _ = strconv.Atoi(arr[1])
		}

		switch arr[0] {
		case Push:
			stack.Push(item)
		case Pop:
			stack.Pop()
		case GetMax:
			stack.GetMax()
		default:
			continue
		}
	}
	stack.writer.Flush()
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)

	for i, val := range s.items {
		if i == 0 {
			s.max = val
			continue
		}
		if s.max < val {
			s.max = val
		}
	}
}

func (s *Stack) Pop() {
	if len(s.items) == 0 {
		s.writer.WriteString(Error)
		return
	}

	lastIndex := len(s.items) - 1
	s.items = s.items[:lastIndex]

	for i, val := range s.items {
		if i == 0 {
			s.max = val
			continue
		}
		if s.max < val {
			s.max = val
		}
	}

	return
}

func (s *Stack) GetMax() {
	if len(s.items) == 0 {
		s.writer.WriteString(None)
		return
	}
	s.writer.WriteString(fmt.Sprint(s.max, "\n"))
}

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func readNum(reader *bufio.Reader) int {
	num, _, _ := reader.ReadLine()
	intNum, _ := strconv.Atoi(string(num))

	return intNum
}

func readArrStr(num int, reader *bufio.Reader) []string {
	arrStr := make([]string, 0, 1)

	for i := num; i > 0; i-- {
		str, _ := reader.ReadString('\n')
		arrStr = append(arrStr, strings.TrimSpace(str))
	}

	return arrStr
}
