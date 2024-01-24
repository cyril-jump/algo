package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
