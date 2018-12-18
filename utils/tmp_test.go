package utils

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		value []data
	}{
		/* {
			[]data{
				{89635, 20120},
				{47756, 20120},
				{93650, 20120},
				{96744, 20120},
				{66294, 20120},
				{38642, 20120},
				{88512, 20120},
				{49286, 96744},
				{47734, 49286},
				{10164, 47734},
				{45784, 47734},
				{29386, 88512},
				{20120, 0},
			},
		},
		{
			[]data{
				{20120, 0},
				{12345, 20120},
				{34567, 20120},
				{78892, 12345},
				{23847, 34567},
			},
		}, */
		{
			[]data{
				{89635, 20120},
				{47756, 20120},
				{93650, 20120},
				{96744, 20120},
				{66294, 20120},
				{38642, 20120},
				{88512, 20120},
				{49286, 96744},
				{47734, 49286},
				{10164, 47734},
				{45784, 47734},
				{29386, 88512},
				{20120, 0},
				{12345, 20120},
				{10001, 20120},
				{10002, 20120},
				{10003, 20120},
				{10004, 20120},
				{10005, 20120},
				{20001, 10001},
				{20002, 10001},
				{20003, 10001},
				// {30001, 20003},
				// {30002, 20003},
				{30003, 20001},
				// {40001, 30003},
				// {40002, 30001},
				// {40003, 30001},
			},
		},
	}

	/* list := []data{
		{89635, 20120},
		{47756, 20120},
		{93650, 20120},
		{96744, 20120},
		{66294, 20120},
		{38642, 20120},
		{88512, 20120},
		{49286, 96744},
		{47734, 49286},
		{10164, 47734},
		{45784, 47734},
		{29386, 88512},
		{20120, 0},
	} */
	// 预期结果
	/*
		result := []data{
			{20120, 0},
			{96744, 20120},
			{49286, 96744},
			{47734, 49286},
			{45784, 47734},
			{10164, 47734},

			{93650, 20120},
			{89635, 20120},
			{88512, 20120},
			{29386, 88512},
			{66294, 20120},
			{47756, 20120},
			{38642, 20120},
		}
	*/

	companyID := 20120
	for _, test := range tests {
		t.Error("------------", sortCompany(companyID, test.value))
	}
	// t.Error("------------", sortCompany(companyID, list))
}

var benchList = []data{
	{89635, 20120},
	{47756, 20120},
	{93650, 20120},
	{96744, 20120},
	{66294, 20120},
	{38642, 20120},
	{88512, 20120},
	{49286, 96744},
	{47734, 49286},
	{10164, 47734},
	{45784, 47734},
	{29386, 88512},
	{20120, 0},
	// {12345, 20120},
	// {10001, 20120},
	// {10002, 20120},
	// {10003, 20120},
	// {10004, 20120},
	// {10005, 20120},
	// {20001, 10001},
	// {20002, 10001},
	// {20003, 10001},
	// {30001, 20001},
	// {30002, 20001},
	// {30003, 20002},
	// {40001, 30003},
	// {40002, 30001},
	// {40003, 30001},
}

func BenchmarkSortMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortCompany(20120, benchList)
	}
}

// 需求:
// 给一个结构体切片排序，该结构体有本公司号和母公司号两个字段。一个切片中只有一个母公司(会给出该顶级公司,其pid 为0,如切片所示);
// 先排二级公司，然后三级等等；如果同级公司有多个，则按公司号降序排序.类似树的前序遍历(解释得不太清楚😁)

func sortCompany(companyID int, list []data) []data {
	newList := make([]data, 0)
	var nextCompanyID int
	var friends []int
	originPid := 20120 // 顶级公司号
	companyMap := make(map[int]int)

	for _, v := range list {
		companyMap[v.id] = v.companyPid
	}

	for len(companyMap) > 0 {
		// 遍历切片
		for id, pid := range companyMap {
			if len(companyMap) == 1 {
				newList = append(newList, data{id, pid})
				return newList
			}
			if companyID == id {
				// 将数据保存到列表中,以便返回
				newList = append(newList, data{id, pid})
				nextCompanyID = getNextCompanyID(companyID, companyMap)

				// 是否有子公司
				ok := hasSubCompany(nextCompanyID, companyMap)
				if !ok && nextCompanyID == 0 {
					// 判断是否有同伴
					friends = getFriends(companyID, companyMap)
					if len(friends) >= 1 {
						nextCompanyID = friends[0]
					} else {
						nextCompanyID = getNextCompanyID(originPid, companyMap)
					}
				}
				companyID = nextCompanyID
				// 剔除符合条件的公司
				delete(companyMap, id)
			}
		}
	}

	return newList
}

func getFriends(companyID int, companyMap map[int]int) []int {
	var friends []int

	// find pid
	companyPid := 0
	for id, pid := range companyMap {
		if id == companyID {
			companyPid = pid
		}
	}

	for id, pid := range companyMap {
		if pid == companyPid && id != companyID {
			friends = append(friends, id)
		}
	}
	sort.SliceStable(friends, func(i, j int) bool {
		return friends[i] > friends[j]
	})
	return friends
}

func getNextCompanyID(companyID int, companyMap map[int]int) int {
	var ids []int
	for id, pid := range companyMap {
		if pid == companyID {
			ids = append(ids, id)
		}
	}

	sort.SliceStable(ids, func(i, j int) bool {
		return ids[i] > ids[j]
	})
	if len(ids) == 0 {
		return 0
	}
	return ids[0]
}

func hasSubCompany(companyID int, companyMap map[int]int) bool {
	for _, pid := range companyMap {
		if pid == companyID {
			return true
		}
	}
	return false
}

type data struct {
	id         int
	companyPid int
}

func TestUTF8(t *testing.T) {
	// t.Error(utf8.RuneCountInString("abc熊和宽"))
	// t.Error(len("0000和20"))
	var tests = []struct {
		value string
	}{
		{"123"},
		{"000000"},
		{""},
		{"abc123"},
		{"123456"},
		{"0012111"},
	}
	// 匹配六位连续的数字
	reg := regexp.MustCompile("^\\d{6}$")

	for _, test := range tests {
		if !reg.MatchString(test.value) {
			t.Error("无法匹配:", test.value)
		}
	}

}

func TestStrings(t *testing.T) {
	t.Error(strings.Join([]string{"a", "b", "c"}, ","))
}

func TestRecursion(t *testing.T) {
	var list = []data{
		{89635, 20120},
		{47756, 20120},
		{93650, 20120},
		{96744, 20120},
		{66294, 20120},
		{38642, 20120},
		{88512, 20120},
		{49286, 96744},
		{47734, 49286},
		{10164, 47734},
		{45784, 47734},
		{29386, 88512},
		// {20120, 0},
		{12345, 20120},
		{10001, 20120},
		{10002, 20120},
		{10003, 20120},
		{10004, 20120},
		{10005, 20120},
		{20001, 10001},
		{20002, 10001},
		{20003, 10001},
		{30001, 20003},
		{30002, 20003},
		{30003, 20001},
		{40001, 30003},
		{40002, 30001},
		{40003, 30001},
	}

	sort.Slice(list, func(i int, j int) bool {
		return list[i].id > list[j].id
	})
	fmt.Println(list)
	result := make([]data, 0)
	t.Error(recursion(list, &result, 20120))
}

func recursion(list []data, res *[]data, pid int) []data {

	for _, v := range list {
		if pid == v.companyPid {
			*res = append(*res, data{v.id, v.companyPid})
			recursion(list, res, v.id)
		}
	}
	return *res
}

var benchList1 = []data{
	{89635, 20120},
	{47756, 20120},
	{93650, 20120},
	{96744, 20120},
	{66294, 20120},
	{38642, 20120},
	{88512, 20120},
	{49286, 96744},
	{47734, 49286},
	{10164, 47734},
	{45784, 47734},
	{29386, 88512},
	// {20120, 0},
	{12345, 20120},
	{10001, 20120},
	{10002, 20120},
	{10003, 20120},
	{10004, 20120},
	{10005, 20120},
	{20001, 10001},
	{20002, 10001},
	{20003, 10001},
	{30001, 20001},
	{30002, 20001},
	{30003, 20002},
	{40001, 30003},
	{40002, 30001},
	{40003, 30001},
}

func BenchmarkSortByRecursion(b *testing.B) {
	res := make([]data, 0)
	for i := 0; i < b.N; i++ {
		recursion(benchList1, &res, 20120)
	}
}
