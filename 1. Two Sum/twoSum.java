import java.util.HashMap;

class Solution {
	public int[] twoSum(int[] nums, int target) {
		HashMap<Integer, Integer> numsMap = new HashMap<>();
		int[] numbers = new int[2];

		for (int i = 0; i < nums.length; i++) {
			numsMap.put(nums[i], i);
		}

		for (int i = 0; i < nums.length; i++) {
			int complement = target - nums[i];

			if (numsMap.containsKey(complement) && numsMap.get(complement) != i) {
				numbers[0] = i;
				numbers[1] = numsMap.get(complement);
			}
		}

		return numbers;
	}
}