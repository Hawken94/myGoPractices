package main

import (
	"myGoPractices/GolangPrograming/beegologs/logics"

	"gitlab.gumpcome.com/common/go_kit/logkit"
)

func main() {
	logkit.InitGraylog("")
	logics.Handle()
}
