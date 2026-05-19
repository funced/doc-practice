package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", project)
	http.ListenAndServe(":8080", nil)
}
func project(w http.ResponseWriter, r *http.Request) {
	cv := "welcome to my api world"
	fmt.Fprintln(w, cv)
}
