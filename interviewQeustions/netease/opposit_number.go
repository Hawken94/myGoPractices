package netease

// OppositNum 求一个"相反数"; 规则:首先将原先的数颠倒,再加上原来的数,便得到相反数;前缀0去掉
// 1325: 1325+5231=6556
func OppositNum(num int) int {
	tmp := num
	res := 0
	for num > 0 {
		res = res*10 + num%10
		num = num / 10
	}

	return res + tmp
}
