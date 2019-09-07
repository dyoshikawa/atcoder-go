package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

type IntListForSort []int

func (s IntListForSort) Len() int           { return len(s) }
func (s IntListForSort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntListForSort) Less(i, j int) bool { return s[i] < s[j] }

func sortIntList(intList []int) {
	sort.Sort(IntListForSort(intList))
}

type IntListForSortDesc []int

func (s IntListForSortDesc) Len() int           { return len(s) }
func (s IntListForSortDesc) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntListForSortDesc) Less(i, j int) bool { return s[i] > s[j] }

func sortIntListDesc(intList []int) {
	sort.Sort(IntListForSortDesc(intList))
}

// ASC前提
func binarySearchInt(intList []int, key int, iMin int, iMax int) int {
	if iMax < iMin {
		return -1
	} else {
		iMid := iMin + (iMax-iMin)/2
		if intList[iMid] > key {
			return binarySearchInt(intList, key, iMin, iMid-1)
		} else if intList[iMid] < key {
			return binarySearchInt(intList, key, iMid+1, iMax)
		} else {
			return iMid
		}
	}
}

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
	n, _ := strconv.Atoi(inputLn())
	aList := make([]int, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(inputLn())
		aList[i] = a
	}
	aListClone := make([]int, n)
	_ = copy(aListClone, aList)
	sortIntList(aListClone)
	for _, a := range aList {
		tgIndex := binarySearchInt(aListClone, a, 0, len(aListClone)-1)
		if tgIndex == n-1 {
			fmt.Println(aListClone[n-2])
		} else {
			fmt.Println(aListClone[n-1])
		}
	}
}
