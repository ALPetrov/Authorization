package main

import (
	"fmt"
	"log"
	"net/http"
)

	func main() {

	http.HandleFunc("/", mainPage)
	
	port := ":8080"
	fmt.Println("Port open", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
func mainPage(w http.ResponseWriter, r *http.Request ) {
	w.Write([]byte("Hello"))
}




