package spewtest

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestDatastruct(t *testing.T) {
	a := 10
	spew.Dump(a)
	t.Error("this result is")
}
