package main

import (
	"bufio"
	"flag"
	"fmt"
	"myGoPractices/go_concurrent_programe/talk/chatbot"
	"os"
	"runtime/debug"
)

var chatbotName string

func init() {
	flag.StringVar(&chatbotName, "chatbot", "simple.cn", "The chatbot's name for dialogue.")
}

func main() {
	flag.Parse()
	chatbot.Register(chatbot.NewSimpleCN("simple.cn", nil))
	myChatbot := chatbot.Get(chatbotName)
	if myChatbot == nil {
		err := fmt.Errorf("Fatal error:Unsupported chatbot named %s", chatbotName)
		checkError(myChatbot, err, true)
	}
	inputReader := bufio.NewReader(os.Stdin)
	begin, err := myChatbot.Begin()
	checkError(myChatbot, err, true)
	fmt.Println(begin)
	input, err := inputReader.ReadString('\n')
	checkError(myChatbot, err, true)
	fmt.Println(myChatbot.Hello(input[:len(input)-1]))

	for {
		input, err = inputReader.ReadString('\n')
		if checkError(myChatbot, err, false) {
			continue
		}

		output, end, err := myChatbot.Talk(input)
		if checkError(myChatbot, err, false) {
			continue
		}
		if output != "" {
			fmt.Println(output)
		}
		if end {
			err = myChatbot.End()
			checkError(myChatbot, err, false)
			os.Exit(0)
		}

	}
}

func checkError(chatbot chatbot.Chatbot, err error, exit bool) bool {
	if err == nil {
		return false
	}
	if chatbot != nil {
		fmt.Println(chatbot.ReportError(err))
	} else {
		fmt.Println(err)
	}
	if exit {
		debug.PrintStack()
		os.Exit(1)
	}
	return true
}
