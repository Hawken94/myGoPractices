package algo

import ("math")

func array_right(nums []int, target int) []int {
	length := len(nums)
	result := make([]int, length) // 初始化没给足够的空间
	for i := 0; i < length; i++ {
		result[(i+target)%length] = nums[i]
	}
	return result
}

// 求质因数
func factorPrime(num int) []int {
	factors := factor(num)
	result := make([]int, 0)

	for _, i := range factors {
		if isPrime(i) {
			result = append(result, i)
		}
	}
	return result
}

// 因数
func factor(num int) []int {
	result := make([]int, 0)
	for i := 1; i < num/2; i++ {
		if num%i == 0 {
			result = append(result, i)
		}
	}
	result = append(result, num)
	return result
}

// 质数
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	flag := false
	sqrt := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			flag = true
		}

	}
	return !flag
}
