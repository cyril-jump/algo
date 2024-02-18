package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := makeReader()
	writer := makeWriter()

	lenght := readNum(reader)
	arr := readArray(reader)
	x := readNum(reader)

	i := Solution(arr, x, lenght-lenght, lenght-1)
	j := Solution(arr, x*2, lenght-lenght, lenght-1)

	printResult(writer, i, j)
}

func Solution(arr []int, x, l, r int) int {
	if l > r {
		return -1
	}

	if arr[l] >= x {
		return l + 1
	}

	m := l + (r-l)/2

	if arr[m] >= x {
		if m == l+1 {
			return m + 1
		} else {
			return Solution(arr, x, l, m+1)
		}
	} else if arr[m] > x {
		return Solution(arr, x, l, m)
	} else {
		return Solution(arr, x, m+1, r)
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

func readArray(reader *bufio.Reader) []int {
	str, _ := reader.ReadString('\n')
	strArr := strings.Split(strings.TrimSpace(str), " ")
	intArr := make([]int, len(strArr))

	for i, val := range strArr {
		intArr[i], _ = strconv.Atoi(val)
	}

	return intArr
}

func makeWriter() *bufio.Writer {
	writer := bufio.NewWriter(os.Stdout)

	return writer
}

func printResult(writer *bufio.Writer, i, j int) {
	writer.WriteString(fmt.Sprintf("%d %d\n", i, j))
	writer.Flush()
}
