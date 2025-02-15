package common

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
