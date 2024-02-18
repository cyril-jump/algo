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

	if n == 0 {
		printEmptyString(writer)
	} else {
		arr := readArrInt(reader)
		res := solution(arr)
		printResult(res, writer)
	}
}

func solution(arr []int) []int {
	x := make([]int, 3)
	res := make([]int, 0, len(arr))

	for _, val := range arr {
		x[val]++
	}

	for i, val := range x {
		for j := 0; j < val; j++ {
			res = append(res, i)
		}
	}

	return res
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

func printEmptyString(writer *bufio.Writer) {
	writer.WriteString("")
	writer.Flush()
}

func printResult(arr []int, writer *bufio.Writer) {
	for i, val := range arr {
		if i == len(arr)-1 {
			writer.WriteString(fmt.Sprintf("%d\n", val))
		} else {
			writer.WriteString(fmt.Sprintf("%d ", val))
		}
	}

	writer.Flush()
}
