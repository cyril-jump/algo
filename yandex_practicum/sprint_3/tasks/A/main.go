package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	reader := makeReader()
	writer := makeWriter()

	n := readNum(reader)

	solution(n, 0, 0, "", writer)

	writer.Flush()
}

func solution(n int, i, j int, line string, writer *bufio.Writer) {
	if i == n && j == n {
		writer.WriteString(line + "\n")
	} else {
		if i < n {
			solution(n, i+1, j, line+"(", writer)
		}
		if j < i {
			solution(n, i, j+1, line+")", writer)
		}
	}
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

func makeWriter() *bufio.Writer {
	writer := bufio.NewWriter(os.Stdout)

	return writer
}
