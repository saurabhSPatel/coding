package sorting

/*
	Algo:
	1. Find min value
	2. swap it with the current pos
	3. Repeat 1,2


    TC && SC O(n^2),O(n^2),O(n^2)  && o(1) inplace algo

    Stable NO

    Adaptive No

*/

func selectionSort(nums []int) []int {

	for pass := 0; pass < len(nums); pass++ {
		min := pass
		for i := pass + 1; i < len(nums); i++ {
			if nums[i] < nums[min] {
				min = i
			}
		}
		nums[min], nums[pass] = nums[pass], nums[min]
	}
	return nums
}
