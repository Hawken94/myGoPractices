package main

import (
	"errors"
	"fmt"
)

func pressButton() {
	fmt.Println("I'm Mr.Meeseeks,look at me!!!")
	panic("It's gettin' weird!")
}
func doStuff() (err error) {
	// var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("the meeseeks went carzy!")
		}
	}()
	pressButton()
	return err
}
