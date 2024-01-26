func twoSum(nums []int, target int) []int {
	for i:=0;i<len(nums);i++{
		for b:=0;b<len(nums);b++{
			 if nums[i]+nums[b]==target{
			 return int(nums[i],nums[b])
  
		 }
	}
	}
	}
  