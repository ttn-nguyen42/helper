package helper

import (
	"bufio"
	"strconv"
	"strings"
)

// Reads integer from a number of lines
func ReadInteger(sc *bufio.Scanner, times int) ([]int, error) {
	err := isInputValid(sc, times)
	if err != nil {
		return nil, err
	}
	t := times
	var nums []int
	for t > 0 {
		n, err := readInt(sc)
		if err != nil {
			return nums, err
		}
		nums = append(nums, n)
		t -= 1
	}
	return nums, nil
}

// Reads strings from a number of lines.
func ReadString(sc *bufio.Scanner, times int) ([]string, error) {
	err := isInputValid(sc, times)
	if err != nil {
		return nil, err
	}
	t := times
	var str []string
	for t > 0 {
		s := readString(sc)
		str = append(str, s)
		t -= 1
	}
	return str, nil
}

// Reads int array for a number of times.
// Can be used to read a matrix
func ReadIntArray(sc *bufio.Scanner, times int) ([][]int, error) {
	err := isInputValid(sc, times)
	if err != nil {
		return nil, err
	}
	t := times
	var arrs [][]int
	for t > 0 {
		nums, err := readIntArray(sc, " ")
		if err != nil {
			return arrs, err
		}
		arrs = append(arrs, nums)
		t -= 1
	}
	return arrs, nil
}

// Reads string array for a number of times.
// Can be used to read a matrix
func ReadStringArray(sc *bufio.Scanner, times int) ([][]string, error) {
	err := isInputValid(sc, times)
	if err != nil {
		return nil, err
	}
	t := times
	var strs [][]string
	for t > 0 {
		str := readStringArray(sc, " ")
		strs = append(strs, str)
		t -= 1
	}
	return strs, nil
}

func readInt(sc *bufio.Scanner) (int, error) {
	str := readString(sc)
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, ErrIntegerInvalid
	}
	return num, err
}

func readString(sc *bufio.Scanner) string {
	sc.Scan()
	l := sc.Text()
	l = strings.Trim(l, " ")
	return l
}

func readStringArray(sc *bufio.Scanner, delim string) []string {
	sc.Scan()
	l := sc.Text()
	res := strings.Split(l, delim)
	return res
}

func readIntArray(sc *bufio.Scanner, delim string) ([]int, error) {
	s := readStringArray(sc, delim)
	var nums []int
	for _, n := range s {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nums, ErrIntegerInvalid
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func isInputValid(sc *bufio.Scanner, times int) error {
	if sc == nil {
		return ErrNilScanner
	}
	if times < 1 {
		return ErrInvalidTimes
	}
	return nil
}
