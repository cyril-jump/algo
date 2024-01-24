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

	arr1 := readArrRune(reader)
	arr2 := readArrRune(reader)

	if ok := solution(arr1, arr2); ok {
		writer.WriteString("True\n")
	} else {
		writer.WriteString("False\n")
	}

	writer.Flush()
}

func solution(arr1, arr2 []rune) bool {
	i := 0
	j := len(arr1)

	for _, val := range arr2 {
		if val == arr1[i] {
			i++
		}

		if i == j {
			return true
		}

	}

	return false
}

func mergeSort(array []rune) []rune {
	if len(array) == 1 {
		return array
	}

	left := mergeSort(array[0 : len(array)/2])

	right := mergeSort(array[len(array)/2 : len(array)])

	result := make([]rune, len(array))

	l, r, k := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result[k] = left[l]
			l++
		} else {
			result[k] = right[r]
			r++
		}
		k++
	}

	for l < len(left) {
		result[k] = left[l]
		l++
		k++
	}
	for r < len(right) {
		result[k] = right[r]
		r++
		k++
	}

	return result
}

func makeReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)

	return reader
}

func readArrRune(reader *bufio.Reader) []rune {
	str, _ := reader.ReadString('\n')
	return []rune(strings.TrimSpace(str))
}

func readArrStr(reader *bufio.Reader) []string {
	str, _ := reader.ReadString('\n')
	arrStr := strings.Split(strings.TrimSpace(str), "")

	return arrStr
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
	writer := bufio.NewWriter(os.Stdout)

	return writer
}

func printArrInt(arr []int, writer *bufio.Writer) {
	for i, val := range arr {
		if i == len(arr)-1 {
			writer.WriteString(fmt.Sprintf("%d\n", val))
		} else {
			writer.WriteString(fmt.Sprintf("%d", val))
		}
	}
}

func printArrStr(arr []string, writer *bufio.Writer) {
	for i, val := range arr {
		if i == len(arr)-1 {
			writer.WriteString(fmt.Sprintf("%s\n", val))
		} else {
			writer.WriteString(fmt.Sprintf("%s", val))
		}
	}
}
