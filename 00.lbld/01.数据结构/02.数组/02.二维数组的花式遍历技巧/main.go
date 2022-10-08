package main

import "fmt"

func main() {

	// 01. 顺时针旋转图像
	//rotateTest()

	// 02. 矩阵的螺旋遍历
	//spiralOrderTest()

	// 03. 给定n，按照1到n*n的顺序螺旋生成矩阵
	generateMatrixTest()
}

func rotateTest() {
	var row, col = 3, 3
	var matrix [][]int
	for i := 0; i < row; i++ {
		var item = make([]int, col)
		matrix = append(matrix, item)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = i*2 + j
		}
	}
	traverseMatrix(matrix)
	rotate(matrix)
	traverseMatrix(matrix)
}

func rotate(matrix [][]int) {

	// 将整个矩阵按左上到右下的对角线翻转
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[0]); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	// 将每一行翻转
	for i := 0; i < len(matrix); i++ {
		var left, right = 0, len(matrix[0]) - 1
		for left < right {
			temp := matrix[i][left]
			matrix[i][left] = matrix[i][right]
			matrix[i][right] = temp
			left++
			right--
		}
	}
	return
}

func traverseMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], ",")
		}
		fmt.Println()
	}
	fmt.Println()
}

func spiralOrderTest() {
	var row, col = 3, 4
	var matrix [][]int
	for i := 0; i < row; i++ {
		var item = make([]int, col)
		matrix = append(matrix, item)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = i*2 + j
		}
	}

	traverseMatrix(matrix)
	fmt.Println(spiralOrder(matrix))

}

func spiralOrder(matrix [][]int) []int {

	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return nil
	}

	var row = len(matrix)
	var col = len(matrix[0])

	var left = 0
	var up = 0
	var down = len(matrix) - 1
	var right = len(matrix[0]) - 1

	var res []int
	for len(res) < row*col {
		fmt.Println(len(res), ",", row*col)
		if up <= down {
			for j := left; j <= right; j++ {
				res = append(res, matrix[up][j])
			}
			up++
		}

		if left <= right {
			for i := up; i <= down; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		if up <= down {
			for j := right; j >= left; j-- {
				res = append(res, matrix[down][j])
			}
			down--
		}

		if left <= right {
			for i := down; i >= up; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}

	}

	return res
}

func generateMatrixTest() {
	traverseMatrix(generateMatrix(5))
}

func generateMatrix(n int) [][]int {

	var matrix [][]int
	row, col := n, n

	for i := 0; i < row; i++ {
		var item = make([]int, col)
		matrix = append(matrix, item)
	}

	var left = 0
	var right = n - 1
	var up = 0
	var down = n - 1

	cur := 1
	for cur <= n*n {
		if up <= down {
			for j := left; j <= right; j++ {
				matrix[up][j] = cur
				cur++
			}
			up++
		}

		if left <= right {
			for i := up; i <= down; i++ {
				matrix[i][right] = cur
				cur++
			}
			right--
		}

		if up <= down {
			for j := right; j >= left; j-- {
				matrix[down][j] = cur
				cur++
			}
			down--
		}

		if left <= right {
			for i := down; i >= up; i-- {
				matrix[i][left] = cur
				cur++
			}
			left++
		}

	}

	return matrix
}
