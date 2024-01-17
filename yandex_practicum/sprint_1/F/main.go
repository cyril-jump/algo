package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

const (
	True  string = "True"
	False string = "False"
)

func main() {
	writer := makeWriter()
	reader := makeReader()

	str := readStr(reader)
	if ok := solution(str); !ok {
		printResult(False, writer)
	} else {
		printResult(True, writer)
	}
}

func solution(line string) bool {
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	str := re.ReplaceAllString(line, "")
	str = strings.ToLower(str)

	arr := []rune(str)

	reversed := make([]rune, len([]rune(str)))

	for i := 0; i < len(arr); i++ {
		reversed[i] = arr[len(arr)-1-i]
	}

	if string(arr) == string(reversed) {
		return true
	}

	return false
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
