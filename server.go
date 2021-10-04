package main

import (
	"fmt"
	"os"
)

const (
	usage = `
    Invalid input. see usage below:
    usage: ./app action "argument"
    actions: an action can either be 'search' or 'describe'
    example: ./app search "text to search"`
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println(usage)
		return
	}
	action := arguments[0]
	param := arguments[1]
	switch action {
	case "search":
		search(param)
	case "describe":
		describe(param)
	default:
		fmt.Println(usage)
		return
	}
}

func search(param string) {
	fmt.Println("Searching ...")
}

func describe(param string) {
	fmt.Println("Describing ...")
}
