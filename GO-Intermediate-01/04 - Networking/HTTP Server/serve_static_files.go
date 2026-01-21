package main

import "net/http"

func main() {

	// http://127.0.0.1:8080/static.go -> ./static/static.go

	// use /static as file loading
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)

}
