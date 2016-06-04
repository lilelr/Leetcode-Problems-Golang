package main

func hIndex(citations []int)int  {
	m,max := make(map [int]int),0
	for _,v := range citations{
		m[v]++
	}
	for i,h := 0,len(citations); i <= h;i++{
		max = i
		h = h -m[i]
	}
	return max
}

func main() {
}
