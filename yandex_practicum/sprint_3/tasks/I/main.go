package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Conf [][]int

func (c Conf) Len() int {
	return len(c)
}

func (c Conf) Less(k, m int) bool {
	if c[k][1] == c[m][1] {
		return c[k][0] < c[m][0]
	}

	return c[k][1] > c[m][1]
}

func (c Conf) Swap(k, m int) {
	c[k], c[m] = c[m], c[k]
}

func main() {
	reader := makeReader()
	writer := makeWriter()

	_ = readNum(reader)
	arr := readArrInt(reader)
	k := readNum(reader)

	res := solution(arr, k)

	printResult(res, writer)
}

func solution(m map[int]int, k int) []int {
	conf := Conf{}
	res := make([]int, 0, k)

	for i, val := range m {
		conf = append(conf, []int{i, val})
	}

	sort.Stable(conf)

	if k > conf.Len() {
		k = conf.Len()
	}

	for i := 0; i < k; i++ {
		res = append(res, conf[i][0])
	}

	return res
}

func makeReader() *bufio.Reader {
	reader := bufio.NewReader(os.Stdin)

	return reader
}

func readArrInt(reader *bufio.Reader) map[int]int {
	str, _ := reader.ReadString('\n')
	arrStr := strings.Split(strings.TrimSpace(str), " ")
	m := make(map[int]int, len(arrStr))

	for _, val := range arrStr {
		x, _ := strconv.Atoi(val)
		m[x]++
	}

	return m
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

func printResult(arr []int, writer *bufio.Writer) {
	for _, val := range arr {
		writer.WriteString(fmt.Sprintf("%d ", val))
	}

	writer.Flush()
}
