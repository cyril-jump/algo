package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	True  string = "True\n"
	False string = "False\n"
)

type Stack struct {
	items []string
}

func main() {
	writer := makeWriter()
	reader := makeReader()

	arr := readArrStr(reader)

	if ok := Solution(arr); ok {
		printResult(True, writer)
	} else {
		printResult(False, writer)
	}
}

func NewStack() *Stack {
	return &Stack{
		items: []string{},
	}
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastItem
}

func makeReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)

	return reader
}

func readArrStr(reader *bufio.Reader) []string {
	str, _ := reader.ReadString('\n')
	strArr := strings.Split(strings.TrimSpace(str), "")

	return strArr
}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res string, writer *bufio.Writer) {
	writer.WriteString(res + "\n")
	writer.Flush()
}

func Solution(tasks []string) bool {
	stack := NewStack()

	for _, val := range tasks {
		if len(stack.items) == 0 {
			stack.Push(val)
			continue
		}
		item := stack.Pop()
		str := item + val

		if str == "()" || str == "{}" || str == "[]" {
			continue
		}
		stack.Push(item)
		stack.Push(val)
	}

	if len(stack.items) == 0 {
		return true
	}
	return false
}
