package main

import (
	"fmt"
	"os"

	"github.com/ochom/focus/domain"
	"github.com/ochom/focus/helpers"
)

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println(domain.USAGE)
		return
	}
	action := arguments[1]
	param := arguments[2]
	switch action {
	case "search":
		helpers.Search(param)
	case "describe":
		helpers.Describe(param)
	default:
		fmt.Println(domain.USAGE)
		return
	}
}
