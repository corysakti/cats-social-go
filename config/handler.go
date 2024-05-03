package config

import (
	"fmt"
	"net/http"
)

func Handler() {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//logic web
		fmt.Println(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
