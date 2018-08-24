package game

var (
	// Score 分数
	Score int
	step  int
)

// Status 游戏状态
type Status uint

// Win ...
const (
	Win Status = iota
	Lose
	And
	Max = 2048
)

// G2048 游戏中的16个格子使用4x4的二维数组
type G2048 [4][4]int

func convertPrintStr(x, y int, str string) {

}
