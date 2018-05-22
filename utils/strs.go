package utils

// StrsNotBlank 判断[]string的元素是否全为空
func StrsNotBlank(strs []string) bool {
	for _, str := range strs {
		if str != "" {
			return true
		}
	}
	return false
}
