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
	Top    string = "top"
	Error  string = "error\n"
	None   string = "None\n"
)

type Item struct {
	Num int
	Max int
}

type Stack struct {
	items  []Item
	writer *bufio.Writer
}

func main() {
	reader := makeReader()
	num := readNum(reader)
	tasks := readArrInt(num, reader)
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
		case Top:
			stack.Peek()

		default:
			continue
		}
	}

	stack.writer.Flush()
}

func NewStack() *Stack {
	writer := bufio.NewWriter(os.Stdout)
	return &Stack{
		items:  []Item{},
		writer: writer,
	}
}

func (s *Stack) Push(item int) {
	if len(s.items) == 0 {
		s.items = append(s.items, Item{
			Num: item,
			Max: item,
		})
		return
	}
	max := s.items[len(s.items)-1].Max
	if max < item {
		max = item
	}

	s.items = append(s.items, Item{
		Num: item,
		Max: max,
	})

}

func (s *Stack) Pop() {
	if len(s.items) == 0 {
		s.writer.WriteString(Error)
		return
	}

	lastIndex := len(s.items) - 1
	s.items = s.items[:lastIndex]

	return
}

func (s *Stack) Peek() {
	if len(s.items) == 0 {
		s.writer.WriteString(Error)
		return
	}

	s.writer.WriteString(fmt.Sprint(s.items[len(s.items)-1].Num, "\n"))

	return
}

func (s *Stack) GetMax() {
	if len(s.items) == 0 {
		s.writer.WriteString(None)
		return
	}
	s.writer.WriteString(fmt.Sprint(s.items[len(s.items)-1].Max, "\n"))
}

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func readNum(reader *bufio.Reader) int {
	num, _, _ := reader.ReadLine()
	intNum, _ := strconv.Atoi(string(num))

	return intNum
}

func readArrInt(num int, reader *bufio.Reader) []string {
	strArr := make([]string, 0, 1)

	for i := num; i > 0; i-- {
		str, _ := reader.ReadString('\n')
		strArr = append(strArr, strings.TrimSpace(str))
	}

	return strArr
}
