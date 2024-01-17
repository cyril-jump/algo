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

	_ = readNum(reader)
	temperatures := readArrInt(reader)
	res := solution(temperatures)

	printResult(res, writer)
}

func solution(temperatures []int) int {
	var x int

	if len(temperatures) == 1 {
		return 1
	}

	for i, val := range temperatures {
		switch i {
		case 0:
			if val > temperatures[i+1] {
				x++
			}
		case len(temperatures) - 1:
			if val > temperatures[i-1] {
				x++
			}
		default:
			if val > temperatures[i+1] && val > temperatures[i-1] {
				x++
			}
		}
	}
	return x
}

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func readNum(reader *bufio.Reader) int {
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))

	return num
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
