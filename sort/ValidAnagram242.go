package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(w string) string{
	s := strings.Split(w,"")
	sort.Strings(s)
	return strings.Join(s,"")
}



// slow 60ms
func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen!=tLen {
		return false
	}
	s = sortString(s)
	t = sortString(t)
	for i:= range s{
		if s[i] != t[i] {
			return false
		}

	}
	return true

}

type strInfo struct  {
	s string
	l int
	m map[rune] int

}

// fast 26ms
// https://leetcode.com/discuss/103063/a-golang-solution
func isAnagram2(s string,t string) bool  {
	var array [2]strInfo
	for i,v := range []string{s,t} {
		array[i].s = v
		array[i].m = make(map[rune] int)
		for _,val := range array[i].s {
			array[i].l++
			array[i].m[val]++
		}
	}

	if array[0].l != array[1].l {
		return false
	}

	for k0,v0 := range array[0].m{
		if v1,ok := array[1].m[k0]; !ok || v0 != v1{
			return false
		}
	}
	return true

}

func main() {

	fmt.Println(isAnagram2("abc","cby"))
}