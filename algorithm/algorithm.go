package algorithm

import (
	"math"
)

func min(a ...int) int {
	result := math.MaxInt32
	for _, v := range a {
		if result > v {
			result = v
		}
	}
	return result
}

func calculateLeftTopNearest(n, m int, data [][]bool) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if data[i][j] {
				result[i][j] = 0
				continue
			}
			result[i][j] = math.MaxInt32
			if i > 0 {
				result[i][j] = min(result[i][j], result[i-1][j]+1)
			}
			if j > 0 {
				result[i][j] = min(result[i][j], result[i][j-1]+1)
			}
		}
	}
	return result
}

func reverseVertical(n, m int, data [][]int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = data[n-i-1][j]
		}
	}
	return result
}

func reverseHorizontal(n, m int, data [][]int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = data[i][m-j-1]
		}
	}
	return result
}

func reverse(n, m int, data [][]int) [][]int {
	return reverseVertical(n, m, reverseHorizontal(n, m, data))
}

func reverseVerticalBool(n, m int, data [][]bool) [][]bool {
	result := make([][]bool, n)
	for i := 0; i < n; i++ {
		result[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			result[i][j] = data[n-i-1][j]
		}
	}
	return result
}

func reverseHorizontalBool(n, m int, data [][]bool) [][]bool {
	result := make([][]bool, n)
	for i := 0; i < n; i++ {
		result[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			result[i][j] = data[i][m-j-1]
		}
	}
	return result
}

func reverseBool(n, m int, data [][]bool) [][]bool {
	return reverseVerticalBool(n, m, reverseHorizontalBool(n, m, data))
}

// Solve calculates distance to nearest true for each element of array.
// Input data is expected to be a rectangular, i.e. len(data[i]) is equal for each i.
func Solve(n, m int, data [][]bool) [][]int {
	leftTopNearest := calculateLeftTopNearest(n, m, data)
	rightTopNearest := reverseHorizontal(n, m, calculateLeftTopNearest(n, m, reverseHorizontalBool(n, m, data)))
	leftBottomNearest := reverseVertical(n, m, calculateLeftTopNearest(n, m, reverseVerticalBool(n, m, data)))
	rightBottomNearest := reverse(n, m, calculateLeftTopNearest(n, m, reverseBool(n, m, data)))

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = min(
				leftTopNearest[i][j],
				rightTopNearest[i][j],
				leftBottomNearest[i][j],
				rightBottomNearest[i][j],
			)
		}
	}
	return result
}
