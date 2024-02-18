package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := makeReader()
	writer := makeWriter()

	_, k := read2Nums(reader)
	arr := readArrInt(reader)

	res := solution(k, arr)

	printResult(res, writer)
}

func solution(k int, arr []int) int {
	sort.Ints(arr)
	n := 0

	for _, val := range arr {
		k -= val
		if k >= 0 {
			n++
		} else {
			break
		}
	}

	return n
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

func read2Nums(reader *bufio.Reader) (int, int) {
	str, _ := reader.ReadString('\n')
	arrStr := strings.Split(strings.TrimSpace(str), " ")

	num1, _ := strconv.Atoi(arrStr[0])
	num2, _ := strconv.Atoi(arrStr[1])

	return num1, num2
}

func makeWriter() *bufio.Writer {
	writer := bufio.NewWriter(os.Stdout)

	return writer
}

func printResult(res int, writer *bufio.Writer) {
	writer.WriteString(fmt.Sprintf("%d\n", res))
	writer.Flush()
}
