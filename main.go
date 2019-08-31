package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type ByValue []float64

func (s ByValue) Len() int           { return len(s) }
func (s ByValue) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByValue) Less(i, j int) bool { return s[i] < s[j] }

func input() string {
	var s string
	_, _ = fmt.Scan(&s)
	return s
}

func inputLn() string {
	var s string
	_, _ = fmt.Scanln(&s)
	return s
}

func inputByScanner(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func convListStrToInt(strList []string) []int {
	intList := make([]int, len(strList))
	for k, v := range strList {
		intV, _ := strconv.Atoi(v)
		intList[k] = intV
	}
	return intList
}

func convListStrToFloat(strList []string) []float64 {
	intList := make([]float64, len(strList))
	for k, v := range strList {
		intV, _ := strconv.Atoi(v)
		intList[k] = float64(intV)
	}
	return intList
}

func main() {
	m, _ := strconv.Atoi(input())
	d, _ := strconv.Atoi(input())

	res := 0
	for i := 1; i <= m; i++ {
		for j := 10; j <= d; j++ {
			dList := convListStrToInt(strings.Split(strconv.Itoa(j), ""))
			prod := 1
			if dList[0] >= 2 && dList[1] >= 2 {
				for _, v := range dList {
					prod *= v
				}
				if i == prod {
					res++
				}
			}
		}
	}

	fmt.Println(res)
}
