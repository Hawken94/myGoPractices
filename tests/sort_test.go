package tests

import (
	"sort"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestSort(t *testing.T) {
	var data []map[string]interface{}
	tmp := make(map[string]interface{})
	tmp["small"] = "abc"
	tmp["middle"] = "ssa"
	tmp["main"] = "yyy"
	data = append(data, tmp)

	tmp = make(map[string]interface{})
	tmp["small"] = "abc"
	tmp["middle"] = "ssa"
	tmp["main"] = "yya"
	data = append(data, tmp)
	tmp = make(map[string]interface{})
	tmp["small"] = "abc"
	tmp["middle"] = "ssb"
	tmp["main"] = "yyb"
	data = append(data, tmp)
	tmp = make(map[string]interface{})
	tmp["small"] = ""
	tmp["middle"] = "yya"
	data = append(data, tmp)

	sortList("", "small", data)
	// spew.Dump(data)
	sortList("small", "middle", data)
	// fmt.Println("----------")
	// spew.Dump(data)
	sortList("middle", "main", data)
	// fmt.Println("----------")
	spew.Dump(data)
	t.Error(data)
}

func sortList(baseName, goodsCateName string, list []map[string]interface{}) {

	sort.SliceStable(list, func(i int, j int) bool {
		if baseName == "" || list[i][baseName].(string) == list[j][baseName].(string) {
			return list[i][goodsCateName].(string) < list[j][goodsCateName].(string)
		}
		return i < j
	})
}
