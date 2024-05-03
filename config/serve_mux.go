package config

import (
	"fmt"
	"net/http"
)

func ServeMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "Hello World")
	})
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "Hi")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "Hello World item")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "Hello World item")
	})
	mux.HandleFunc("/images/thumbnail", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(writer, "Hello World item thumbnails/")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
