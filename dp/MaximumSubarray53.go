package main

func maxSubArray(nums []int)int  {
	res := nums[0]
	tempSub := 0
	for _,v := range nums{
		 tempSub += v
		res = maxofTwoItems(res,tempSub)
		res = maxofTwoItems(res,v)
		if tempSub<0{
			tempSub = 0
		}
	}
	return res
}

func maxofTwoItems(item1 int,item2 int)int  {
	if item1 > item2{
		return item1
	}
	return item2
}

func main() {
	nums := []int{4,-1,2,1}
	print(maxSubArray(nums))
}
