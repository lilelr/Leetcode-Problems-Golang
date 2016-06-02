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
func main() {

	fmt.Println(isAnagram("abc","cba"))
}