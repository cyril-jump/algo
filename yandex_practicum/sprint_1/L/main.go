package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func main() {
	writer := makeWriter()
	reader := makeReader()

	a := readStr(reader)
	b := readStr(reader)

	res := solution(a, b)

	printResult(res, writer)
}

func solution(s1 string, t1 string) string {
	s2 := strings.Split(s1, "")
	t2 := strings.Split(t1, "")
	sort.Strings(s2)
	sort.Strings(t2)

	if len(s2) > len(t2) {
		return extraLetter(s2, t2)
	}

	return extraLetter(t2, s2)
}

func extraLetter(big, small []string) string {
	for i, val := range small {
		if val != big[i] {
			return big[i]
		}
	}

	return big[len(big)-1]
}

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func readStr(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')

	return strings.TrimSpace(str)
}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res string, writer *bufio.Writer) {
	writer.WriteString(res + "\n")
	writer.Flush()
}
