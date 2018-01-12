package interest

// 题目:
// 在一个二维数组中, 每一行都按照从左到右递增的顺序排序, 每一列都 good 照从上到下的顺序排序.
// 请完成一个函数, 输入这样的一个二维数组和一个函数, 判断数组中是否含有该整数

func findMatrix(matrix [][]int, target int) bool {
	found := false
	rows, cols := len(matrix), len(matrix[0])

	if rows > 0 && cols > 0 {
		row := 0
		col := cols - 1
		for row < rows && col >= 0 {
			if matrix[row][col] == target {
				found = true
				break
			} else if matrix[row][col] > target {
				col--
			} else {
				row++
			}
		}
	}
	return found
}
