package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	const input = "xhk is a good \n boy"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)
	fmt.Println(string(65), ":", strconv.Itoa(65))
	fmt.Println(8 & 7)
	GuessingGame()
}

func GuessingGame() {

	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
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
