package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type fn map[string]func(b, a int) int

func addition(b, a int) int {
	return b + a
}

func subtraction(b, a int) int {
	return b - a
}

func multiplication(b, a int) int {
	return b * a
}

func division(b, a int) int {
	return rounding(b, a)
}

const (
	Addition       string = "+"
	Subtraction    string = "-"
	Multiplication string = "*"
	Division       string = "/"
)

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

func main() {
	reader := makeReader()
	writer := makeWriter()

	data := readArrStr(reader)

	result := Solution(data)

	printResult(writer, result)
}

func Solution(data []string) int {
	stack := NewStack()
	calc := fn{
		Addition:       addition,
		Subtraction:    subtraction,
		Multiplication: multiplication,
		Division:       division,
	}

	for _, val := range data {
		item, err := strconv.Atoi(val)
		if err != nil {
			a, b := stack.PopItem(), stack.PopItem()
			item = calc[val](b, a)
		}
		stack.PushItem(item)
	}
	return stack.PopItem()
}

func (s *Stack) PushItem(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) PopItem() int {
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return lastItem
}

func rounding(b, a int) int {
	x := math.Floor(float64(b) / float64(a))

	return int(x)
}

func makeReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)

	return reader
}

func readArrStr(reader *bufio.Reader) []string {
	str, _ := reader.ReadString('\n')
	strArr := strings.Split(strings.TrimSpace(str), " ")

	return strArr
}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(writer *bufio.Writer, result int) {
	writer.WriteString(fmt.Sprintf("%d\n", result))
	writer.Flush()
}
