package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	c := "package main;func main(){c:=%q;print(c,c)}"
	print(c, c)
}

func discoverGoRoot() string {
	goroot := os.Getenv("GOROOT")
	fmt.Println("GOROOT:", goroot)
	if goroot == "" {
		command := exec.Command("go", "env", "GOROOT")
		command.Env = append(os.Environ(), "GOPATH=/tmp/go")
		output, err := command.Output()
		log.Fatal(err, "could not find go binary")
		goroot = string(output)
	}
	return strings.TrimSpace(goroot)
}
