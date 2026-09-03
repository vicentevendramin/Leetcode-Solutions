package main

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		indexComplement, ok := numsMap[complement]

		if ok && indexComplement != i {
			return []int{i, indexComplement}
		}
	}

	return nil
}
