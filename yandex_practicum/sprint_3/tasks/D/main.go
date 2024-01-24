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
	reader := makeReader()
	writer := makeWriter()

	n := readNum(reader)
	factors := readArrInt(reader)
	k := readNum(reader)
	sizes := readArrInt(reader)

	res := solution(n, k, factors, sizes)

	printResult(res, writer)
}

func solution(n, k int, factors, sizes []int) int {
	x, i, j := 0, 0, 0

	sort.Ints(factors)
	sort.Ints(sizes)

	for {
		if i == n || j == k {
			break
		}
		if factors[i] <= sizes[j] {
			x++
			i++
			j++
			continue
		}

		if factors[i] > sizes[j] {
			j++
		}

	}

	return x
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

func readNum(reader *bufio.Reader) int {
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))

	return num
}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res int, writer *bufio.Writer) {
	writer.WriteString(fmt.Sprintf("%d", res))
	writer.Flush()
}
