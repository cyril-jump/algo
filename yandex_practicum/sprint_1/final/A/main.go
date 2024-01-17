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
	arr := readArrInt(reader)

	res := solution(n, arr)

	printResult(res, writer)
}

func solution(num int, street []int) []int {
	if num == 0 {
		return []int{0}
	}

	x := make([]int, num)
	zero := num

	for i := 0; i < num; i++ {
		if street[i] == 0 {
			zero = i
			continue
		}

		if zero != num {
			x[i] = i - zero
		}
	}

	zero = num

	for i := num - 1; i >= 0; i-- {

		if street[i] == 0 {
			zero = i
			continue
		}

		if (x[i] > zero-i || x[i] == 0) && zero != num {
			x[i] = zero - i
		}
	}

	return x
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

func readArrInt(reader *bufio.Reader) []int {
	str, _ := reader.ReadString('\n')
	arrStr := strings.Split(strings.TrimSpace(str), " ")
	arrInt := make([]int, len(arrStr))

	for i, val := range arrStr {
		arrInt[i], _ = strconv.Atoi(val)
	}

	return arrInt
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
