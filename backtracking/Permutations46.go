package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	start := 0
        var tempList []int
	backtracking(nums,start,res,tempList)
	return res

}

func backtracking(nums []int,start int,res [][]int,tempList []int)  {
	if len(tempList) == len(nums){
		res = append(res,tempList)
	}
	for _,v:= range nums{
		flag := false
		for _,tempListItem:= range tempList{
			if tempListItem == v{
				flag = true
			}
		}

		if !flag{
			tempList = append(tempList,v)
			backtracking(nums,0,res,tempList)
			tempList = tempList[0:len(tempList)-1]
		}
	}
}

func main() {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	res :=permute(nums)
	fmt.Print(res)
	fmt.Print("fe")

}
