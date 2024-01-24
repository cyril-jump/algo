package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var letters = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func main() {
	reader := makeReader()
	writer := makeWriter()

	digits := readArrInt(reader)

	solution(digits, "", writer)

	writer.Flush()
}

func solution(digits []int, prefix string, writer *bufio.Writer) {
	if len(digits) == 0 {
		writer.WriteString(prefix + " ")
		return
	}

	digit := digits[0]
	for _, letter := range letters[digit] {
		solution(digits[1:], prefix+string(letter), writer)
	}
}

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func readArrInt(reader *bufio.Reader) []int {
	str, _ := reader.ReadString('\n')
	arrStr := strings.Split(strings.TrimSpace(str), "")
	arrInt := make([]int, len(arrStr))

	for i, val := range arrStr {
		arrInt[i], _ = strconv.Atoi(val)
	}

	return arrInt
}

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}
