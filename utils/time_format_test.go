package utils

import (
	"fmt"
	"testing"
	"time"

	"gitlab.gumpcome.com/common/go_kit/timekit"
)

func TestFormatTime(t *testing.T) {
	tests := []string{
		"20180131",
		"20180228",
		"20180331",
		"20180430",
		"20180531",
		"20180630",
		"20180731",
		"20180831",
		"20180930",
		"20181031",
		"20181130",
		"20181231",
	}

	for _, test := range tests {
		now, _ := timekit.StringToTime(test, timekit.DateFormat_YYYYMMDD)
		// lastMonth := now.Local().AddDate(0, -1, -now.Day()+1)
		lastMonth := now.Local().AddDate(0, 0, -now.Day())
		t.Errorf("xhk format_time :%v %v \n", formatTime(now), formatTime(lastMonth))
	}
	t.Error(time.Now().Year(), int(time.Now().Month()))
}

func TestMonth(t *testing.T) {
	now, _ := timekit.StringToTime("2018-08-01 00:06:00", timekit.DateFormat_YYYY_MM_DD_HH_MM_SS)
	// lastMonth := now.Local().AddDate(0, -1, -now.Day()+1)
	beforeFiveMin := now.Add(-5 * time.Minute)
	beforeTenMin := now.Add(-10 * time.Minute)

	month := int(beforeTenMin.Month())

	fmt.Println(beforeFiveMin, beforeTenMin, month)
	t.Errorf("xhk format_time :%v %v \n", beforeFiveMin.Month(), beforeTenMin.Month())
}

func TestTmp(t *testing.T) {
	str := "1_42_A5F6BC316A2511E8871F0017FA00DF17_10000_10003"
	t.Error(str[len(str)-11 : len(str)-6])
	isFirstOfMonth := time.Now().Day() == 1
	t.Error(isFirstOfMonth)

}
