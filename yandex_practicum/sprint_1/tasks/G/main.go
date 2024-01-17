package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	writer := makeWriter()
	reader := makeReader()

	n := readNum(reader)
	res := solution(n)

	printResult(res, writer)
}

func solution(n int) []int {
	if n == 0 {
		return []int{0}
	}
	arr := make([]int, 0, 8)
	k := n
	m := 0

	for {
		if k == 0 {
			break
		}
		m = k % 2
		k = k / 2
		arr = append(arr, m)
	}

	reversed := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		reversed[i] = arr[len(arr)-1-i]
	}

	return reversed
}

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func readNum(reader *bufio.Reader) int {
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))

	return num
}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res []int, writer *bufio.Writer) {
	for _, val := range res {
		writer.WriteString(fmt.Sprintf("%d", val))
	}
	writer.WriteString("\n")
	writer.Flush()
}
