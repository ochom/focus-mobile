package main

import (
	"fmt"
	"log"
	"net/http"
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
	case "run":
		runServer()
	default:
		fmt.Println(domain.USAGE)
		return
	}
}

func runServer() {
	PORT := helpers.GetEnv("PORT")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello world")
	})
	http.HandleFunc("/register", helpers.Register())
	http.HandleFunc("/add-rx", helpers.AddPrescriptions())

	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
