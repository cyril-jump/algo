package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	Win  string = "WIN"
	Fail string = "FAIL"
)

func main() {
	writer := makeWriter()
	reader := makeReader()

	arr := readArrInt(reader)
	if ok := solution(arr); !ok {
		printResult(Fail, writer)
	} else {
		printResult(Win, writer)
	}
}

func solution(arr []int) bool {
	x := math.Abs(float64(arr[0]%2)) + math.Abs(float64(arr[1]%2)) + math.Abs(float64(arr[2]%2))
	if x == 3 || x == 0 {
		return true
	}
	return false
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

func printResult(res string, writer *bufio.Writer) {
	writer.WriteString(res + "\n")
	writer.Flush()
}
