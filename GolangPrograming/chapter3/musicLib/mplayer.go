package main

import (
	"awesomeProject/GolangPrograming/chapter3/musicLib/library"
	"awesomeProject/GolangPrograming/chapter3/musicLib/mp"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		{
			if len(tokens) == 6 {
				id++
				lib.Add(&library.Music{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
			} else {
				fmt.Println("USAGE: lib add <name><artist><source><type>")
			}
		}
	case "remove":
		if len(tokens) == 3 {
			/* if tempID, err := strconv.Atoi(tokens[2]); err == nil {
				fmt.Println(tempID)
				// todo 这里改成根据名字remove比较好
				lib.Remove(tempID)
			} */
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("USAGE： lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func hanlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music ", tokens[1], "does not exist.")
		return
	}
	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`
			Enter following commands to control the player:
			lib list -- View the existing music lib 
			lib add <name><artist><source><type> -- Add a music to the music lib
			lib remove <name> -- Remove the specified music from the lib
			play <name> -- Play the specifed music
	`)

	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command-> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			hanlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command: ", tokens[0])
		}
	}
}
