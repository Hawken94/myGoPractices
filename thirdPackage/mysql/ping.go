package main

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"sync"

	"gitlab.gumpcome.com/common/tools/logkit"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		ping(5)
	}()
	go func() {
		defer wg.Done()
		ping(8)
	}()
	go func() {
		defer wg.Done()
		ping(1)
	}()
	go func() {
		defer wg.Done()
		ping(2)
	}()
	wg.Wait()
	fmt.Println("-------done------")
}

func ping(age int) error {
	data := url.Values{
		"age": []string{strconv.Itoa(age)},
	}

	resp, err := http.PostForm("http://127.0.0.1:2333/add", data)
	if err != nil {
		logkit.Error(err)
		return err
	}
	logkit.Debug(resp)
	return nil
}
