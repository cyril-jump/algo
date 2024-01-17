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
	str := readStr(reader)

	res := solution(str)

	printResult(res, writer)
}

func solution(line string) string {
	arr := strings.Split(line, " ")
	x, y := 0, 0

	for i, val := range arr {
		if x < len([]rune(val)) {
			y = i
			x = len([]rune(val))
		}

	}
	return arr[y]
}

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func readNum(reader *bufio.Reader) int {
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))

	return num
}

func readStr(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')

	return strings.TrimSpace(str)
}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res string, writer *bufio.Writer) {
	writer.WriteString(res + "\n")
	writer.WriteString(fmt.Sprintf("%d\n", len(res)))

	writer.Flush()
}
