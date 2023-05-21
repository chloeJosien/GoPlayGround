// problem found at https://leetcode.com/problems/two-sum/
package leetCodeTwoSum

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				var result []int
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return nil
}
