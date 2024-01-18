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

	k := readNum(reader)
	matrix := readMatrixInt(4, reader)

	res := solution(k, matrix)

	printResult(res, writer)
}

func solution(k int, matrix [][]int) int {
	sum := make([]int, 9)
	score := 0

	for _, row := range matrix {
		for _, val := range row {
			if val == 0 {
				continue
			}

			sum[val-1] += 1
		}
	}

	for _, val := range sum {
		if k*2 >= val && val != 0 {
			score++
		}
	}
	return score
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
	arrStr := strings.Split(strings.TrimSpace(str), "")
	arrInt := make([]int, len(arrStr))

	for i, val := range arrStr {
		if val == "." {
			arrInt[i] = 0
			continue
		}
		num, _ := strconv.Atoi(val)
		arrInt[i] = num
	}

	return arrInt
}

func readMatrixInt(n int, reader *bufio.Reader) [][]int {
	matrix := make([][]int, 4)

	for i := 0; i < n; i++ {
		matrix[i] = readArrInt(reader)
	}

	return matrix

}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res int, writer *bufio.Writer) {
	writer.WriteString(fmt.Sprintf("%d\n", res))
	writer.Flush()
}
