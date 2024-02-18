package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := makeReader()
	writer := makeWriter()

	n := readNum(reader)
	arr := readArrInt(reader)

	solution(n, arr, writer)

	writer.Flush()
}

func solution(n int, arr []int, writer *bufio.Writer) {
	p := 0
	if n == 0 {
		return
	}

	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			p++
		}
	}

	if p >= 1 {
		printResult(arr, writer)
	} else if p == 0 && n == len(arr) {
		printResult(arr, writer)
		return
	} else if p == 0 {
		return
	}

	solution(n-1, arr, writer)
}

func makeReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)

	return reader
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

func readNum(reader *bufio.Reader) int {
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))

	return num
}

func makeWriter() *bufio.Writer {
	writer := bufio.NewWriter(os.Stdout)

	return writer
}

func printResult(arr []int, writer *bufio.Writer) {
	for i, val := range arr {
		if i == len(arr)-1 {
			writer.WriteString(fmt.Sprintf("%d\n", val))
		} else {
			writer.WriteString(fmt.Sprintf("%d ", val))
		}
	}
}
