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
	writer := makeWriter()
	reader := makeReader()

	n := readNum(reader)
	_ = readNum(reader)

	matrix := readMatrixInt(n, reader)

	rowID := readNum(reader)
	colID := readNum(reader)

	res := solution(matrix, rowID, colID)

	printResult(res, writer)
}

func solution(matrix [][]int, rowID int, colID int) []int {
	var x []int
	var lenRow, lenCol int
	lenRow = len(matrix) - 1

	for _, valRow := range matrix {
		lenCol = len(valRow) - 1
		break
	}

	if rowID-1 >= 0 {
		x = append(x, matrix[rowID-1][colID])
	}

	if rowID+1 <= lenRow {
		x = append(x, matrix[rowID+1][colID])
	}

	if colID-1 >= 0 {
		x = append(x, matrix[rowID][colID-1])
	}

	if colID+1 <= lenCol {
		x = append(x, matrix[rowID][colID+1])
	}

	sort.Ints(x)

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

func readMatrixInt(n int, reader *bufio.Reader) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = readArrInt(reader)
	}

	return matrix

}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res []int, writer *bufio.Writer) {
	for _, val := range res {
		writer.WriteString(fmt.Sprintf("%d ", val))
	}
	writer.WriteString("\n")
	writer.Flush()
}
