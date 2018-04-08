package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

const zero = 0.0

type user struct{}

func (*user) print() {
	fmt.Println("user ")
}

func main() {
	var u user
	u.print()

	ip := "test1"
	fmt.Printf("ip:%v isIn:%v\n", ip, isInAntiLists(ip))

	slice := []int{1, 2, 3}
	for k, v := range slice {
		fmt.Printf("index :%v value:%v\n", k, v)
	}
	fmt.Println("------------------")
	addSlice(&slice)
	for k, v := range slice {
		fmt.Printf("index :%v value:%v\n", k, v)
	}

	io.Copy()

}
func addSlice(slice *[]int) {
	*slice = append(*slice, 4)
	spew.Dump(slice)
}
func isInAntiLists(ip string) bool {
	// 打开文件
	lists := readFile()
	fmt.Printf("lists:%v lenghth:%v \n", lists, len(lists))
	for _, list := range lists {
		if list == ip {
			return true
		}
	}
	return false
}

func addAntiList(ip string) {
	// 写入文件
	// writeFile(ip)
}

// 读取文件
func readFile() []string {
	lists := make([]string, 0)
	inputFile, inputError := os.Open("test.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return nil // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		// fmt.Printf("The input was: %s", inputString)
		lists = append(lists, strings.TrimSpace(inputString))
		if readerError == io.EOF {
			return lists
		}
	}
}
