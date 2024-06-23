package main

import (
	"net/http"

	"myproject/api"
)

func main() {
	// route handler
	// w is  the response writer where we write the response to. has methods
	// r has info about the request. paths
	// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("the server is working"))
	// })
	// // initialize http server, port
	// http.ListenAndServe(":8080", nil)
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
