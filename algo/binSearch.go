package algo

// 返回元素所在的位置
// 1, 2, 3, 4, 5, 6, 7, 8, 9
func binarySearch(arrs []int, lo, hi, target int) int {

	for lo < hi {
		mi := (lo + hi) / 2
		if arrs[mi] < target {
			lo = mi + 1
		} else if target < arrs[mi] {
			hi = mi
		} else {
			return mi
		}
	}
	// lo--
	return lo
}
