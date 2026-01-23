package main

import (
	"errors"
	"fmt"
	"net/http"
	"path"
	"strings"
)

// http.FileSystem is an interface, with only `Open(name string) (File, error)` method
type wrapFS struct {
	fh http.FileSystem
}

func (w wrapFS) Open(name string) (http.File, error) {
	clean := path.Clean(name)
	fmt.Println("WrapFS=", clean)
	base := path.Base(clean)
	if strings.HasPrefix(base, ".") == true {
		return nil, errors.New("No no, you can't see.")
	}
	return w.fh.Open(clean)
}

// logging
func printRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Incoming=", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // without this middleware doesn't work
	})
}

func main() {

	// use /static as file loading
	// without wrapFS http://127.0.0.1:8080/static/.env -> 200 OK
	// with wrapFS http://127.0.0.1:8080/static/.env -> 500 Internal Server Error
	root := http.Dir("./static")
	fs := http.FileServer(wrapFS{root})

	fsHandler := http.StripPrefix("/static/", fs)

	http.Handle("/", printRequest(fsHandler))

	http.ListenAndServe(":8080", nil)
}
