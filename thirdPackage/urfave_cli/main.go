package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make the world better"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! | say!")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Now())
}
