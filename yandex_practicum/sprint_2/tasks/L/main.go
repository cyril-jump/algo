package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

func (s *Stack) Put2Item(item1, item2 int) {
	s.items = append(s.items, item1, item2)
}

func (s *Stack) Get2Item() (int, int) {
	lastIndex := len(s.items) - 1
	lastItem1 := s.items[lastIndex]
	lastItem2 := s.items[lastIndex-1]
	s.items = s.items[:lastIndex-1]

	return lastItem1, lastItem2
}

func main() {
	writer := makeWriter()
	reader := makeReader()

	n, k := readNK(reader)

	res := Solution(n, k)

	printResult(res, writer)
}

func Solution(n, k int) int {
	if n <= 1 {
		return 1
	}

	stack := NewStack()
	stack.Put2Item(1, 1)
	k = int(math.Pow(10, float64(k)))

	for i := n - 2; i >= 1; i-- {
		a, b := stack.Get2Item()
		stack.Put2Item(a, (a+b)%k)
	}

	a, b := stack.Get2Item()

	return (a + b) % k
}

func makeReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)

	return reader
}

func readNK(reader *bufio.Reader) (int, int) {
	str, _, _ := reader.ReadLine()
	arr := strings.Split(strings.TrimSpace(string(str)), " ")
	n, _ := strconv.Atoi(arr[0])
	k, _ := strconv.Atoi(arr[1])

	return n, k
}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res int, writer *bufio.Writer) {
	writer.WriteString(fmt.Sprintf("%d\n", res))
	writer.Flush()
}
