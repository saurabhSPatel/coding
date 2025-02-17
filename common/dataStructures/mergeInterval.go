package common

func mergeIntervals(intervals [][]int) [][]int {
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1] = []int{res[len(res)-1][0], max(intervals[i][1], res[len(res)-1][1])}
		} else {
			res = append(res, intervals[i])
		}
	}

	// Replace this placeholder return statement with your code
	return res
}
