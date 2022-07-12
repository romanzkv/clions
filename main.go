package main

import (
	"fmt"
	"github.com/gbin/goncurses"
)

func main() {
	fmt.Println("let there be .... something")
	src, err := goncurses.Init()
	if err != nil {
		log.Fatal("init:", err)
	}
	defer goncurses.End()
}
