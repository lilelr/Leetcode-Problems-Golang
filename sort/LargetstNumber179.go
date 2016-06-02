package main

import (
	"strconv"
	"fmt"
)
//https://leetcode.com/discuss/104943/a-golang-solution
func largestNumber(nums []int) string {
	var result string
	var lens []int

	for _, v := range nums {
		var find bool
		s := strconv.Itoa(v)
		totalLen := 0

		for i, v := range lens {
			if (s + result[totalLen:]) > (result[totalLen:totalLen+v] + s + result[totalLen+v:] ){
				result = result[0:totalLen] + s + result[totalLen:]
				lens = append(lens[:i], append([]int{len(s)}, lens[i:]...)...)
				find = true
				break
			}
			totalLen += v
		}

		if !find {
			result += s
			lens = append(lens, len(s))
		}
	}

	for result[0] == '0' && len(result) > 1 {
		result = result[1:]
	}

	return result
}

func main() {
	nums := []int{34,30,9}
	res := largestNumber(nums)
	fmt.Println(res)
}
