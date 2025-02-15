package common

import "fmt"

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	top, botton := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	for top <= botton && left <= right {
		//left-->right
		for i := left; i <= right; i++ {
			fmt.Println(top, i, matrix[top][i])
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i <= botton; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if top <= botton {
			for i := right; i >= left; i-- {
				res = append(res, matrix[botton][i])
			}
			botton--
		}
		if left <= right {
			for i := botton; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}

	}
	return res

}
