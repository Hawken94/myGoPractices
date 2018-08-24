package tests

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestMap(t *testing.T) {
	strs := []string{"a", "a", "b", "c", "c"}
	result := make(map[string]interface{})
	for _, str := range strs {
		result[str] = str
	}
	spew.Dump(result)
	t.Error(len(result))
}
