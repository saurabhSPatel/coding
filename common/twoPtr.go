package common

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums) - 1
	for i := 0; i <= n; i++ {

		if nums[i] > 0 || i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			}
		}
	}
	return res
}
