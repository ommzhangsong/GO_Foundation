/*
webprogram
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func page(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
	fmt.Println(w, "\n", r)
}
func main() {

	http.HandleFunc("/", page)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("listen", err)
	}
}
