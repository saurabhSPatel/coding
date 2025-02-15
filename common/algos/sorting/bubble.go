package sorting

/*
Algo :
1. Each pass find out max element by checking nums[i] > nums[i+1]
2 Run the loop till (len(nums) - pass - 1) and swap if nums[i] is greater

TC && SC(B/A/W) O(n) / O(n²) / O(n²)  && o(1) inplace algo

Stable if two elements have the same value, the algorithm ensures that the element that appeared first in the input will also appear first in the output.
Adaptive We can close ths early if array is already sorted

Best use for Small datasets or nearly sorted datasets.

*/

func bubbleSort(nums []int) []int {

	for pass := 0; pass < len(nums); pass++ {
		// var isSorted = true
		for i := 0; i < len(nums)-pass-1; i++ {
			if nums[i] > nums[i+1] {
				//if nums[i] < nums[i+1] { if want desc order
				nums[i+1], nums[i] = nums[i], nums[i+1]
				// isSorted = false
			}
		}
		// if isSorted {
		// 	break
		// }
	}
	return nums

}
