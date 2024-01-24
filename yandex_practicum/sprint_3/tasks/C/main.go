package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	True  string = "True\n"
	False string = "False\n"
)

func main() {
	reader := makeReader()
	writer := makeWriter()

	arr1 := readArrRune(reader)
	arr2 := readArrRune(reader)

	if ok := solution(arr1, arr2); ok {
		printResult(True, writer)
	} else {
		printResult(False, writer)
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

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func readArrRune(reader *bufio.Reader) []rune {
	str, _ := reader.ReadString('\n')
	return []rune(strings.TrimSpace(str))
}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res string, writer *bufio.Writer) {
	writer.WriteString(res)
	writer.Flush()
}
