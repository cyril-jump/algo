package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	True  string = "True"
	False string = "False"
)

func main() {
	writer := makeWriter()
	reader := makeReader()

	str := readNum(reader)
	if ok := solution(str); !ok {
		printResult(False, writer)
	} else {
		printResult(True, writer)
	}
}

func solution(number int) bool {
	n := math.Log(float64(number)) / math.Log(4)
	_, frac := math.Modf(n)

	if frac != 0 {
		return false
	}

	return true
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

func printResult(res string, writer *bufio.Writer) {
	writer.WriteString(res + "\n")
	writer.Flush()
}
