package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(arr []int) int {
	return arr[1]*(arr[0]*arr[1]+arr[2]) + arr[3]
}

func main() {
	writer := makeWriter()
	reader := makeReader()

	arr := readArrInt(reader)
	res := solution(arr)

	printResult(res, writer)
}

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
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

func printResult(res int, writer *bufio.Writer) {
	writer.WriteString(fmt.Sprintf("%d\n", res))
	writer.Flush()
}
