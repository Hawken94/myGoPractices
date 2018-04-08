package juejin

// 一个数是不是快乐是这么定义的：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，或是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。

func getNHappyNum(n int) []int {
	nums := make([]int, 0)
	for i := 1; i <= n; i++ {
		if isHappy(i) {
			nums = append(nums, i)
		}
	}
	return nums
}
func isHappy(num int) bool {
	x, y := num, num
	for x > 1 {
		x = getNextNum(x)
		if x == 1 {
			return true
		}
		y = getNextNum(getNextNum(y))
		if y == 1 {
			return true
		}
		if x == y {
			return false
		}
	}

	return true
}

func getNextNum(num int) int {
	sum := 0
	for num > 0 {
		sum += (num % 10) * (num % 10)
		num /= 10
	}
	return sum
}
