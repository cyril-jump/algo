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

func solution(number int) []int {
	arr := make([]int, 0, 20)

	k := number

	for i := 2; i*i <= k; i++ {
		for {
			if k == 1 || k%i != 0 {
				break
			}
			if k%i == 0 {
				k = k / i
				arr = append(arr, i)
			}
		}
	}

	if k != 1 {
		arr = append(arr, k)
	}

	return arr
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
		writer.WriteString(fmt.Sprintf("%d ", val))
	}
	writer.WriteString("\n")
	writer.Flush()
}
