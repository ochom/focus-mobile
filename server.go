package main

import (
	"fmt"
	"os"
)

const (
	usage = `Invalid input. see usage below:
    usage: ./app action "argument"
    actions: an action can either be 'search' or 'describe'
    example: ./app search "text to search"`
)

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println(usage)
		return
	}
	action := arguments[1]
	param := arguments[2]
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
	fmt.Printf("Searching ...%v\n", param)
}

func describe(param string) {
	fmt.Printf("Describing ...%v\n", param)
}
