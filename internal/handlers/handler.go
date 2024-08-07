package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Spence156/go_windows_service/internal/tools"
)

func Web() {
	fmt.Println("Test Web.go")
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func StartServer() {

	config, err := tools.LoadConfig()

	if err != nil {
		panic(err)
	}

	var portNumber string = ":" + strconv.Itoa(config.Port)
	fmt.Printf("Port which we will be listening on: %v \n", portNumber)

	http.HandleFunc("/", HelloWorldHandler)
	http.ListenAndServe(portNumber, nil)
}
