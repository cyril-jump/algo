package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	writer := makeWriter()
	reader := makeReader()

	n := readNum(reader)
	m := readNum(reader)

	matrix := readMatrixStr(n, reader)
	res := solution(n, m, matrix)

	printResult(res, writer)
}

func solution(n, m int, matrix [][]string) [][]string {
	var result [][]string
	for i := 0; i < m; i++ {
		var line = make([]string, 0, n)
		for j := 0; j < n; j++ {
			line = append(line, matrix[j][i])
		}
		result = append(result, line)
	}

	return result
}

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func readNum(reader *bufio.Reader) int {
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))

	return num
}

func readArrStr(reader *bufio.Reader) []string {
	str, _ := reader.ReadString('\n')
	arrStr := strings.Split(strings.TrimSpace(str), " ")

	return arrStr
}

func readMatrixStr(n int, reader *bufio.Reader) [][]string {
	matrix := make([][]string, n)

	for i := 0; i < n; i++ {
		matrix[i] = readArrStr(reader)
	}

	return matrix

}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res [][]string, writer *bufio.Writer) {
	for _, col := range res {
		for i, val := range col {
			if len(col)-1 == i {
				writer.WriteString(val + "\n")
				continue
			}
			writer.WriteString(val + " ")

		}

	}
	writer.Flush()
}
