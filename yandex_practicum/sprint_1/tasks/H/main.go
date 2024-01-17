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

	a := readStr(reader)
	b := readStr(reader)

	res := solution(a, b)

	printResult(res, writer)
}

func solution(a, b string) string {
	r1 := make([]string, 0, 8)
	r2 := make([]string, 0, 8)
	sum := make([]string, 0, 8)
	x, d := 0, 0

	if len([]rune(a)) > len([]rune(b)) {
		r1 = strings.Split(a, "")
		r2 = strings.Split(b, "")
	} else {
		r2 = strings.Split(a, "")
		r1 = strings.Split(b, "")
	}

	for i := 0; i < len(r1); i++ {
		x1, _ := strconv.Atoi(r1[len(r1)-1-i])

		if i < len(r2) {
			x2, _ := strconv.Atoi(r2[len(r2)-1-i])
			x = x1 + x2 + d
		} else {
			x = x1 + d
		}

		switch x {
		case 0:
			sum = append(sum, "0")
			d = 0
		case 1:
			sum = append(sum, "1")
			d = 0
		case 2:
			sum = append(sum, "0")
			d = 1
		case 3:
			sum = append(sum, "1")
			d = 1
		}

	}

	if d != 0 {
		sum = append(sum, "1")
	}

	reversed := make([]string, len(sum))

	for i := 0; i < len(sum); i++ {
		reversed[i] = sum[len(sum)-1-i]
	}

	return strings.Join(reversed, "")
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
