package common

import (
	"container/heap"
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]struct{})

	for i, num := range nums {
		if _, ok := mp[num]; ok {
			return true
		}
		mp[num] = struct{}{}
		if len(mp) > k {
			delete(mp, nums[i-k])
		}

	}
	return false
}

// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//
//	1 [3  -1  -3] 5  3  6  7       3
//	1  3 [-1  -3  5] 3  6  7       5
//	1  3  -1 [-3  5  3] 6  7       5
//	1  3  -1  -3 [5  3  6] 7       6
//	1  3  -1  -3  5 [3  6  7]      7
func maxSlidingWindow(nums []int, k int) []int {

	windMax := math.MinInt
	h := &MaxHeap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		windMax = max(windMax, nums[i])
		heap.Push(h, nums[i])
	}
	res := []int{windMax}

	for i := k; i < len(nums); i++ {
		heap.Push(h, nums[i])
		// if (*h)[0] >= res[len(res)-1] {
		// 	res = append(res, (*h)[0])
		// } else {
		// 	res = append(res, res[len(res)-1])
		// }
		res = append(res, (*h)[0])
	}
	return res

}

// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

/*
	approch := using window and 2pts

*/

func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	windSum := 0
	l := 0

	for r := range len(nums) {
		windSum += nums[r]
		for windSum >= target {
			res = min(res, r-l+1)
			windSum -= nums[l]
			l++
		}

	}
	if res == math.MaxInt {
		return 0
	}
	return res
}
