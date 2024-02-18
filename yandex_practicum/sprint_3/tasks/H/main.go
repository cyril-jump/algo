package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Num []string

func (n Num) Len() int {
	return len(n)
}

func (n Num) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Num) Less(i, j int) bool {
	k, _ := strconv.Atoi(n[i] + n[j])
	m, _ := strconv.Atoi(n[j] + n[i])

	return k > m
}

func main() {
	reader := makeReader()
	writer := makeWriter()

	_ = readNum(reader)
	arr := readArrStr(reader)

	solution(arr)
	printResult(arr, writer)
}

func solution(arr Num) {
	sort.Stable(arr)

}

func makeReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)

	return reader
}

func readArrStr(reader *bufio.Reader) []string {
	str, _ := reader.ReadString('\n')
	arrStr := strings.Split(strings.TrimSpace(str), " ")

	return arrStr
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

func printResult(arr []string, writer *bufio.Writer) {
	for i, val := range arr {
		if i == len(arr)-1 {
			writer.WriteString(fmt.Sprintf("%s\n", val))
		} else {
			writer.WriteString(fmt.Sprintf("%s", val))
		}
	}

	writer.Flush()
}
