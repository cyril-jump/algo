package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Arr []int

func (a Arr) Len() int {
	return len(a)
}

func (a Arr) Less(k, m int) bool {
	return a[k] > a[m]
}

func (a Arr) Swap(k, m int) {
	a[k], a[m] = a[m], a[k]
}

func main() {
	reader := makeReader()
	writer := makeWriter()

	n := readNum(reader)
	arr := readArrInt(reader)

	res := solution(n, arr)

	printResult(res, writer)
}

func solution(n int, arr Arr) int {
	sort.Sort(arr)

	x := 0

	for i := 0; i < n-2; i++ {
		if arr[i] < arr[i+1]+arr[i+2] {
			x = arr[i] + arr[i+1] + arr[i+2]
			break
		}
	}

	return x
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

func printResult(res int, writer *bufio.Writer) {
	writer.WriteString(fmt.Sprintf("%d\n", res))
	writer.Flush()
}
