package sorting

/*
		Algo:
		1. Builds the sorted list one element at a time by taking each element and inserting it into its correct position in the already sorted part of the list.

		Best use case:
	    Small or nearly sorted datasets.
	    Online sorting (when data is coming in real-time).

		Time Complexity: O(n²) in the worst and average cases, O(n) in the best case (if the list is already sorted).

		Space Complexity: O(1) (in-place).
*/
func insertionSort(nums []int) []int {
	//6,2,5,4,3,1
	//6,2,5,4,3,1  k = 2 j = 0  // //2,6,5,4,3,1
	//2,6,5,4,3,1  k = 5, j = 1 // 2,6,6,4,3,1 // j = 0 // 2,5,6,4,3,1
	//2,5,6,4,3,1  k = 4 j = 2 // 2,5,6,6,3,1 // k = 4 j = 1 // 2,5,6,4,3,1
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}

		nums[j+1] = key

	}
	return nums

}
