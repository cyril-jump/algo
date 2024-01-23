package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	writer := makeWriter()
	reader := makeReader()

	n := readNum(reader)
	res := Solution(n)

	printResult(res, writer)
}

func Solution(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	return Solution(n-2) + Solution(n-1)
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

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res int, writer *bufio.Writer) {
	writer.WriteString(fmt.Sprintf("%d\n", res))
	writer.Flush()
}
