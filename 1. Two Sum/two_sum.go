package main

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := numsMap[target-n]; ok {
			return []int{j, i}
		}
		numsMap[n] = i
	}
	return nil
}
