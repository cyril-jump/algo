package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	writer := makeWriter()
	reader := makeReader()

	_ = readNum(reader)
	bigNumer := readArrInt(reader)
	smallNumber := readNum(reader)

	res := solution(bigNumer, smallNumber)

	printResult(res, writer)
}

func solution(bigNumber []int, smallNumber int) []int {
	arrSmallNumber := make([]int, 0, 5)

	strArr := strings.Split(strconv.Itoa(smallNumber), "")

	for _, val := range strArr {
		y, _ := strconv.Atoi(val)
		arrSmallNumber = append(arrSmallNumber, y)
	}

	if len(bigNumber) > len(arrSmallNumber) {
		return twoSum(bigNumber, arrSmallNumber)
	}

	return twoSum(arrSmallNumber, bigNumber)
}

func twoSum(main, second []int) []int {
	var x, d int
	sum := make([]int, 0, 5)

	for i := 0; i < len(main); i++ {
		x1 := main[len(main)-1-i]

		if i < len(second) {
			x2 := second[len(second)-1-i]
			x = x1 + x2 + d
		} else {
			x = x1 + d
		}

		if x > 9 {
			sum = append(sum, x-10)
			d = 1
		} else {
			sum = append(sum, x)
			d = 0
		}
	}

	if d == 1 {
		sum = append(sum, d)
	}
	reversed := make([]int, len(sum))

	for i, val := range sum {
		reversed[len(sum)-1-i] = val
	}

	return reversed
}

func makeReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func readNum(reader *bufio.Reader) int {
	str, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))

	return num
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

func makeWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

func printResult(res []int, writer *bufio.Writer) {
	for _, val := range res {
		writer.WriteString(fmt.Sprintf("%d", val))
	}
	writer.WriteString("\n")
	writer.Flush()
}
