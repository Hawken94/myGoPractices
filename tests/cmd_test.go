package tests

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCMD(t *testing.T) {
	t.Error(discoverGoRoot())
}
func discoverGoRoot() string {
	goroot := os.Getenv("GOROOT")
	fmt.Println("GOROOT:", goroot)
	if goroot == "" {
		output, err := exec.Command("go", "env", "GOROOT").Output()
		log.Fatal(err, "could not find go binary")
		goroot = string(output)
	}
	return strings.TrimSpace(goroot)
}

func TestStrings(t *testing.T) {
	strs := strings.Split("20180830", "-")
	t.Error(strs)
	dateStr := strings.Join(strs, "")
	t.Error(dateStr)
}
